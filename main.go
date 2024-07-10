//go:generate qtc -dir=templates
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/gosimple/slug"
	"github.com/painhardcore/Gobrou/models"
	"github.com/painhardcore/Gobrou/pkg/files"
	"github.com/painhardcore/Gobrou/templates"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func main() {
	// Parse the config.toml file
	var config models.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// Parse each post file
	posts := make([]models.Post, 0)
	err := filepath.Walk("posts", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", path, err)
		}

		var post models.Post
		parts := bytes.SplitN(data, []byte("+++"), 3)
		if len(parts) != 3 {
			return fmt.Errorf("invalid post format %s", path)
		}
		if _, err := toml.Decode(string(parts[1]), &post); err != nil {
			return fmt.Errorf("failed to parse post %s: %v", path, err)
		}
		post.Content = string(parts[2])

		// Generate the slug from the filename
		post.Slug = slug.Make(strings.TrimSuffix(info.Name(), ".md"))

		markdown := goldmark.New(goldmark.WithExtensions(extension.GFM))
		var buf strings.Builder
		if err := markdown.Convert([]byte(post.Content), &buf); err != nil {
			return fmt.Errorf("failed to convert markdown for post %s: %v", path, err)
		}
		post.Content = buf.String()

		posts = append(posts, post)
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to parse posts: %v", err)
	}

	// log.Printf("%+v\n", posts)
	// log.Printf("%+v\n", config)

	// clean
	os.RemoveAll("dist")

	// Create the dist directory if it doesn't exist
	if err := os.MkdirAll("dist", os.ModePerm); err != nil {
		log.Fatalf("Failed to create dist directory: %v", err)
	}

	for _, post := range posts {
		// Generate the post HTML file
		var postBuf bytes.Buffer
		templates.WritePageTemplate(&postBuf, &templates.PostPage{
			Post: &post,
		},
			&config,
		)

		// Write the post HTML file to the dist directory
		postPath := filepath.Join("dist", "posts", post.Slug, "index.html")
		if err := os.MkdirAll(filepath.Dir(postPath), os.ModePerm); err != nil {
			log.Fatalf("Failed to create directory for post %s: %v", post.Slug, err)
		}
		if err := os.WriteFile(postPath, postBuf.Bytes(), 0o644); err != nil {
			log.Fatalf("Failed to write post HTML file %s: %v", postPath, err)
		}
	}

	// Generate the index HTML file
	var indexBuf bytes.Buffer
	templates.WritePageTemplate(
		&indexBuf,
		&templates.IndexPage{
			Posts: posts,
			Cfg:   &config,
		},
		&config,
	)

	// Write the index HTML file to the dist directory
	indexPath := filepath.Join("dist", "index.html")
	if err := os.WriteFile(indexPath, indexBuf.Bytes(), 0o644); err != nil {
		log.Fatalf("Failed to write index HTML file %s: %v", indexPath, err)
	}

	err = files.CopyDir("media", "dist/media")
	if err != nil {
		log.Fatal(err)
	}
	devMode := flag.Bool("dev", false, "Run in development mode")
	flag.Parse()
	if !*devMode {
		fmt.Println("Static blog generated successfully!")
		return
	}
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
