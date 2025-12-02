package main

import (
	"log"
	"os"

	templatetagger "github.com/undeadtokenart/Homepage/templateTagger"
)

func main() {

	// run the server with the "run" tag present
	runServer()

	// --tagger filestart make a template from html file (also I know this is a teribble name)
	if len(os.Args) > 1 && os.Args[1] == "--tagger" {
		file := ""
		filePath := "templates/html/"
		if len(os.Args) > 2 {
			file = filePath + os.Args[2]
		}
		if len(os.Args) < 2 {
			log.Panicln("File name was not provided")
		}
		err := templatetagger.TagTemplateText(file)
		if err != nil {
			log.Println(err)
		}
	}
}
