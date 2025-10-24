# Gobrou

A minimalist static blog generator inspired by the philosophy of [The Best Motherfucking Website](https://thebestmotherfucking.website/) and [Perfect Motherfucking Website](https://perfectmotherfuckingwebsite.com/). Gobrou generates clean, fast, and beautiful static websites with minimal dependencies and maximum performance.

> Educational + practical: This is a learning-driven project to explore interesting ideas while building something useful in a pragmatic way, avoiding overbloat.
> Status: Evolving project without a defined 1.0 yet; APIs and internals may change.

## Philosophy

- **No security risks, no third party** - Self-contained static generation
- **Small** - Minimal footprint and dependencies
- **Easy preview** - Built-in development server
- **Minimal third party deps** - Currently uses some dependencies, but planning to rewrite core libraries

> **Note**: The project currently uses external dependencies for TOML parsing, Markdown processing, and slug generation. However, there are plans to rewrite these core libraries to achieve true minimalism and eliminate third-party dependencies.

## Features

- 📝 **Markdown posts** with TOML frontmatter
- 🎨 **Clean, minimal design** with dark mode support
- ⚡ **Fast generation** using Go and QuickTemplate
- 🔧 **Simple configuration** via TOML
- 📱 **Responsive design** that works everywhere
- 🚀 **Static output** ready for any hosting service
- 🛠️ **Development mode** with live preview

## Installation

```bash
git clone https://github.com/painhardcore/Gobrou.git
cd Gobrou
go build
```

## Usage

### 1. Create your configuration

Create a `config.toml` file in your project root:

```toml
website = "https://yoursite.com"
author = "Your Name"
description = "Your blog description"
title = "Your Blog Title"
postsOnIndex = 3
```

### 2. Create your posts

Create a `posts/` directory and add your markdown files. Each post should have TOML frontmatter:

```markdown
+++
title = "Your Post Title"
description = "A brief description of your post"
date = 2023-06-16T14:07:32-05:00
draft = false
css = "/* Custom CSS for this post */"
+++

## Your Markdown Content

Write your post content here using standard Markdown syntax.
```

### 3. Add media files (optional)

Place any static assets (images, CSS, JS) in a `media/` directory. They will be copied to the output.

### 4. Generate your site

```bash
# Generate static files
./Gobrou

# Or run in development mode with live preview
./Gobrou -dev
```

The generated site will be in the `dist/` directory, ready to deploy to any static hosting service.

## Project Structure

```
your-blog/
├── config.toml          # Site configuration
├── posts/               # Your markdown posts
│   ├── hello-world.md
│   └── another-post.md
├── media/               # Static assets (optional)
│   ├── images/
│   └── styles/
└── dist/                # Generated site (created automatically)
    ├── index.html
    ├── posts/
    └── media/
```

## Post Format

Posts are written in Markdown with TOML frontmatter:

```markdown
+++
title = "Post Title"
description = "Post description"
date = 2023-06-16T14:07:32-05:00
draft = false
css = "/* Optional custom CSS */"
+++

Your markdown content here...
```

### Frontmatter Fields

- `title` (required): The post title
- `description` (optional): Post description for meta tags
- `date` (required): Publication date in RFC3339 format
- `draft` (optional): Set to `true` to exclude from generation
- `css` (optional): Custom CSS for this specific post

## Development

### Prerequisites

- Go 1.22.1 or later
- QuickTemplate (for template generation)

### Building

```bash
# Generate templates
go generate

# Build the binary
go build

# Run tests
go test ./...
```

### Template Development

Templates are written using QuickTemplate and located in the `templates/` directory:

- `base.qtpl` - Base template with common HTML structure
- `index.qtpl` - Homepage template
- `post.qtpl` - Individual post template

After modifying templates, regenerate them with:

```bash
go generate
```

## Roadmap

- [x] Separate templates
- [x] CSS per post
- [x] Write a Readme for Gobrou
- [x] Deploy of the painhardcore.ru
- [ ] About Page
- [ ] Tags cloud
- [ ] 404 page
- [ ] **Rewrite TOML lib** - Replace BurntSushi/toml with custom implementation
- [ ] **Rewrite MD lib** - Replace yuin/goldmark with custom implementation
- [ ] **Rewrite slug lib** - Replace gosimple/slug with custom implementation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source. Check the repository for license details.

## Inspiration

- [The Best Motherfucking Website](https://thebestmotherfucking.website/)
- [Perfect Motherfucking Website](https://perfectmotherfuckingwebsite.com/)
- [Awesome Motherfucking Website](https://github.com/lyoshenka/awesome-motherfucking-website)
