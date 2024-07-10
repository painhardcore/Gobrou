// Code generated by qtc from "index.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/index.qtpl:1
package templates

//line templates/index.qtpl:1
import "github.com/painhardcore/Gobrou/models"

//line templates/index.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/index.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/index.qtpl:5
type IndexPage struct {
	Cfg   *models.Config
	Posts []models.Post
}

//line templates/index.qtpl:12
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/index.qtpl:12
	qw422016.N().S(`
`)
//line templates/index.qtpl:13
	qw422016.E().S(p.Cfg.Title)
//line templates/index.qtpl:13
	qw422016.N().S(`
`)
//line templates/index.qtpl:14
}

//line templates/index.qtpl:14
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/index.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:14
	p.StreamTitle(qw422016)
//line templates/index.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:14
}

//line templates/index.qtpl:14
func (p *IndexPage) Title() string {
//line templates/index.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:14
	p.WriteTitle(qb422016)
//line templates/index.qtpl:14
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:14
	return qs422016
//line templates/index.qtpl:14
}

//line templates/index.qtpl:16
func (p *IndexPage) StreamStyle(qw422016 *qt422016.Writer) {
//line templates/index.qtpl:16
	qw422016.N().S(`
`)
//line templates/index.qtpl:17
}

//line templates/index.qtpl:17
func (p *IndexPage) WriteStyle(qq422016 qtio422016.Writer) {
//line templates/index.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:17
	p.StreamStyle(qw422016)
//line templates/index.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:17
}

//line templates/index.qtpl:17
func (p *IndexPage) Style() string {
//line templates/index.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:17
	p.WriteStyle(qb422016)
//line templates/index.qtpl:17
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:17
	return qs422016
//line templates/index.qtpl:17
}

//line templates/index.qtpl:19
func (p *IndexPage) StreamHead(qw422016 *qt422016.Writer) {
//line templates/index.qtpl:19
	qw422016.N().S(`
`)
//line templates/index.qtpl:20
}

//line templates/index.qtpl:20
func (p *IndexPage) WriteHead(qq422016 qtio422016.Writer) {
//line templates/index.qtpl:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:20
	p.StreamHead(qw422016)
//line templates/index.qtpl:20
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:20
}

//line templates/index.qtpl:20
func (p *IndexPage) Head() string {
//line templates/index.qtpl:20
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:20
	p.WriteHead(qb422016)
//line templates/index.qtpl:20
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:20
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:20
	return qs422016
//line templates/index.qtpl:20
}

//line templates/index.qtpl:22
func (p *IndexPage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/index.qtpl:22
	qw422016.N().S(`
    <h1>`)
//line templates/index.qtpl:23
	qw422016.E().S(p.Cfg.Title)
//line templates/index.qtpl:23
	qw422016.N().S(`</h1>
    <p>`)
//line templates/index.qtpl:24
	qw422016.N().S(p.Cfg.Description)
//line templates/index.qtpl:24
	qw422016.N().S(`</p>
    <hr>
    <h2>Recent Posts</h2>
    <ul>
        `)
//line templates/index.qtpl:28
	for _, post := range p.Posts {
//line templates/index.qtpl:28
		qw422016.N().S(`
        <li><a href="/posts/`)
//line templates/index.qtpl:29
		qw422016.E().S(post.Slug)
//line templates/index.qtpl:29
		qw422016.N().S(`/">`)
//line templates/index.qtpl:29
		qw422016.E().S(post.Title)
//line templates/index.qtpl:29
		qw422016.N().S(`</a></li>
        `)
//line templates/index.qtpl:30
	}
//line templates/index.qtpl:30
	qw422016.N().S(`
    </ul>
`)
//line templates/index.qtpl:32
}

//line templates/index.qtpl:32
func (p *IndexPage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/index.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/index.qtpl:32
	p.StreamBody(qw422016)
//line templates/index.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:32
}

//line templates/index.qtpl:32
func (p *IndexPage) Body() string {
//line templates/index.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/index.qtpl:32
	p.WriteBody(qb422016)
//line templates/index.qtpl:32
	qs422016 := string(qb422016.B)
//line templates/index.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/index.qtpl:32
	return qs422016
//line templates/index.qtpl:32
}