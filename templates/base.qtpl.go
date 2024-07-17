// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/base.qtpl:1
package templates

//line templates/base.qtpl:1
import "github.com/painhardcore/Gobrou/models"

//line templates/base.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/base.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/base.qtpl:4
type Page interface {
//line templates/base.qtpl:4
	Title() string
//line templates/base.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line templates/base.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line templates/base.qtpl:4
	Style() string
//line templates/base.qtpl:4
	StreamStyle(qw422016 *qt422016.Writer)
//line templates/base.qtpl:4
	WriteStyle(qq422016 qtio422016.Writer)
//line templates/base.qtpl:4
	Head() string
//line templates/base.qtpl:4
	StreamHead(qw422016 *qt422016.Writer)
//line templates/base.qtpl:4
	WriteHead(qq422016 qtio422016.Writer)
//line templates/base.qtpl:4
	Body() string
//line templates/base.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line templates/base.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line templates/base.qtpl:4
}

//line templates/base.qtpl:12
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page, cfg *models.Config) {
//line templates/base.qtpl:12
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>`)
//line templates/base.qtpl:18
	p.StreamTitle(qw422016)
//line templates/base.qtpl:18
	qw422016.N().S(`</title>
    <style>
        body {
            max-width: 650px;
            margin: 40px auto;
            padding: 0 10px;
            font: 18px/1.5 -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
            color: #444
       }
       h1,h2,h3 {
           line-height: 1.2
       }

       @media (prefers-color-scheme: dark) {
           body {
               color:#c9d1d9;
               background: #0d1117
           }
           a:link {
               color: #58a6ff
           }
           a:visited {
	       color: #8e96f0
           }
      }

`)
//line templates/base.qtpl:44
	p.StreamStyle(qw422016)
//line templates/base.qtpl:44
	qw422016.N().S(`
    </style>
`)
//line templates/base.qtpl:46
	p.StreamHead(qw422016)
//line templates/base.qtpl:46
	qw422016.N().S(`
</head>
<body>
`)
//line templates/base.qtpl:49
	p.StreamBody(qw422016)
//line templates/base.qtpl:49
	qw422016.N().S(`
    <footer>
        <hr>
        <p>&copy; `)
//line templates/base.qtpl:52
	qw422016.N().S(cfg.Author)
//line templates/base.qtpl:52
	qw422016.N().S(`</p>
    </footer>
    <hr>
</body>
</html>
`)
//line templates/base.qtpl:57
}

//line templates/base.qtpl:57
func WritePageTemplate(qq422016 qtio422016.Writer, p Page, cfg *models.Config) {
//line templates/base.qtpl:57
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/base.qtpl:57
	StreamPageTemplate(qw422016, p, cfg)
//line templates/base.qtpl:57
	qt422016.ReleaseWriter(qw422016)
//line templates/base.qtpl:57
}

//line templates/base.qtpl:57
func PageTemplate(p Page, cfg *models.Config) string {
//line templates/base.qtpl:57
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/base.qtpl:57
	WritePageTemplate(qb422016, p, cfg)
//line templates/base.qtpl:57
	qs422016 := string(qb422016.B)
//line templates/base.qtpl:57
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/base.qtpl:57
	return qs422016
//line templates/base.qtpl:57
}
