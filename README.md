# not hugo 

have you ever thought to yourself "I want to use Hugo but with a lot less functionality and also not as reusable"? 

this repo is a learning task for how I can create a static site generator in golang. I was originally doing it with my repo [maxwell-log](https://github.com/antonbriganti/maxwell-log) in mind.

it's not going to change the world or anything but I learned something *and* I had fun. isn't that why we build things? 

## how does it work? 
I am not kidding this is just like hugo but not as fully featured.

- it takes in any markdown files in a folder called `md`, absorbs basic frontmatter, and converts to html using [Goldmark](github.com/yuin/goldmark)
- it uses go html templates in a folder called `_templates` and then renders the static site including the Goldmark generated html 
- finally, it saves the file to a `dist` folder

the sites these are made for don't require any fancy functionality, literally just basic html. 