# !hugo 

have you ever thought to yourself "I want to use Hugo but with less functionality **and** not as reusable"? 

it's made with my repo [maxwell-log](https://github.com/antonbriganti/maxwell-log) in mind but I guess it's probably built up enough that it can be used for other things. I would't recommend it for anyone who isn't me though.

this started as a way for me to learn go, so the code is a bit 😬 but the fun has been 😍. and isn't that why we build things? 

## how does it work? 
I am not kidding this is just like hugo but not as fully featured.

- it takes in any markdown files in a folder called `md`, absorbs basic frontmatter, and converts to html using [Goldmark](github.com/yuin/goldmark)
- it uses go html templates in a folder called `_templates` and then renders the static site including the Goldmark generated html 
- finally, it saves the file to a `dist` folder

the sites these are made for don't require any fancy functionality, literally just basic html. 