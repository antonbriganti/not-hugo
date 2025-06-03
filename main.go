package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
)

type PageData struct {
	Title    string
	Layout string
	Filename string
	Filepath string
	Content  template.HTML
}

type Metadata struct {
	Title string `yaml:"title"`
	Layout string `yaml:"layout"`
}

func markdownToHTML(mdContent string) (Metadata, string) {
	var buf bytes.Buffer

	// create new goldmark parser using the frontmatter extension and unsafe HTML rendering (risky but I trust me)
	md := goldmark.New(goldmark.WithExtensions(&frontmatter.Extender{}, extension.Strikethrough), goldmark.WithRendererOptions(html.WithUnsafe()))
	ctx := parser.NewContext()

	// convert markdown string into html
	if err := md.Convert([]byte(mdContent), &buf, parser.WithContext(ctx)); err != nil {
		log.Fatal((err))
	}

	var meta Metadata

	// decode frontmatter into metadata struct
	d := frontmatter.Get(ctx)
	if err := d.Decode(&meta); err != nil {
		log.Fatal(err)
	}

	return meta, buf.String()
}

func renderHtml(pd PageData) {
	templates := []string{
		"_templates/head.html",
		"_templates/footer.html",
	}

	layoutTemplate := fmt.Sprintf("_templates/%s.html", pd.Layout)
	templates = append(templates, layoutTemplate)

	outputDir := "dist/" + pd.Filepath
	err := os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	outputFile := fmt.Sprintf("%s/%s.html", outputDir, pd.Filename)

	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := tmpl.ExecuteTemplate(f, pd.Layout+".html", pd); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generated %s\n", outputFile)
}

func main() {
	re := regexp.MustCompile(`^.*?-`)
	err := filepath.Walk("md", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// Read the markdown file
			mdContent, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Convert markdown into html
			meta, mdConverted := markdownToHTML(string(mdContent))

			// create filename + path
			filename := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			var filepath string

			if meta.Layout != "index" {
				filename = re.ReplaceAllString(filename, "")
				filepath += meta.Layout
			}

			// convert into PageData struct
			data := PageData{
				Title:   meta.Title,
				Layout: meta.Layout,
				Filename: filename,
				Filepath: filepath,
				Content: template.HTML(mdConverted),
			}

			// render the HTML
			renderHtml(data)

		}
		return nil

	})

	if err != nil {
		log.Fatal(err)
	}

}
