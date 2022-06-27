# Go Quotes API

This is an API built when I was learning Go. Here I am using 2 modules,

* (Gin)[github.com/gin-gonic/gin] - For quickly building the API
* (Colly)[github.com/gocolly/colly] - For web scrapping.

The API just hits the Goodreads quotes page and fires a search. It returns a list of Quotes with their authors. I have built it and deployed it on Heroku. 

Commands I ran basically from start to finish

### To initate the current folder/project as a module
```bash
go mod init vigneshm.me/go-quotes-api
```
This creates a file called go.mod. This has the name of the project and version of Go it uses.

### To import modules
```bash
go get github.com/gin-gonic/gin
go get github.com/gocolly/colly
```
This imports the modules, creates a go.sum file to add checksums for the modules it downloads.

Next, I filled in the main.go with my code.

Now, for the Heroku part.

### Install Heroku CLI
```bash
npm install -g heroku
```
### Login to Heroku
```bash
heroku login -i
```

### Initate a Git repo and commit the code
```bash
git init
git add .
git commit -m "inital commit"
```

### Create a Procfile for Heroku
```bash
touch Procfile
echo "web: bin/go-quotes-api" > Procfile
```
This basically tells Heroku to run the command on startup.

### Create Heroku App
```bash
heroku create go-quote-api
git push heroku master
```
This create the Heroku app and pushes the code.

Heroku takes this code runs the build and exposes the app at https://<app-name>.herokuapp.com/

My app is available at https://go-quote-app.herokuapp.com/quotes/dumbledore

The last word can be replaced with any search criteria to search for quotes from your favorite character or book.