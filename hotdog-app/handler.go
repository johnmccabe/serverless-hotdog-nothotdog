package function

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	var resp bytes.Buffer

	funcMap := template.FuncMap{
		"safeURL":   safeURL,
		"safeJSStr": safeJSStr,
	}
	t := template.Must(template.New("webpage.tmpl").Funcs(funcMap).ParseFiles("static/webpage.tmpl"))
	err := t.Execute(&resp, classifierConfig())
	if err != nil {
		panic(err)
	}
	return resp.String()
}

func safeURL(s interface{}) template.URL {
	return template.URL(fmt.Sprint(s))
}
func safeJSStr(s interface{}) template.JSStr {
	return template.JSStr(fmt.Sprint(s))
}

type ClassifierConfig struct {
	URL    string
	Entity string
}

func classifierConfig() ClassifierConfig {
	return ClassifierConfig{
		URL:    setDefault("classifier_url", "/function/hotdog-classifier"),
		Entity: setDefault("classifier_entity", "hotdog"),
	}
}

func setDefault(envvar, defaultValue string) string {
	if v := os.Getenv(envvar); len(v) > 0 {
		return v
	}
	return defaultValue
}
