package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

type (
	Template struct {
		templates *template.Template
	}
)

type Task struct {
	ID   int64
	Name string
	Age  int64
}

func init() {
	ctx := context.Background()
	projectID := "heokjin-home"
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	kind := "Poerson"
	name := "id=5629499534213120"
	taskKey := datastore.NameKey(kind, name, nil)
	task := Task{
		Name: "Buy milk",
	}
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}
	fmt.Printf("Saved %v: %v\n", taskKey, task.Name)

	t := &Template{
		templates: template.Must(template.ParseFiles("templates/welcome.html")),
	}
	e.Renderer = t
	e.GET("/welcome", welcome)
	e.GET("/gcs", welcomegcs)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func welcomegcs(c echo.Context) error {
	return c.JSON(http.StatusOK, "AA")
}

func welcome(c echo.Context) error {
	return c.Render(http.StatusOK, "welcome", "Joe")
}
