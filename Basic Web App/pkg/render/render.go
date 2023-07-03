package render

import (
	"Basic_Web_App/pkg/config"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	const PagePath = "./templates/*.page.tmpl"
	const LayoutPath = "./templates/*.layout.tmpl"

	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(PagePath)
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(LayoutPath)
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(LayoutPath)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
