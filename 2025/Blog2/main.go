package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Spatacoli/blog2/models"
	"github.com/Spatacoli/blog2/templates"
)

func main() {
	// Create a sample post
	post := models.Post{
		Title: "My First Blog Post",
		Content: "This is the content of my first blog post. It's exciting!",
		Author: "Todd Spatafore",
		Date: time.Now(),
	} 

	// Render the template
	file, err := os.Create("output.html")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	c := templates.Post(post)
	err = c.Render(context.Background(), file)
	if err != nil {
		log.Fatalf("Error rendering template: %v", err)
	}

	fmt.Println("Blog post generated successfully in output.html")
}