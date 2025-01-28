// Code generated by qtc from "search.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/ybs/search.qtpl:1
package ybs

//line views/ybs/search.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/ybs/search.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/ybs/search.qtpl:1
func (p *SearchPage) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/ybs/search.qtpl:1
	qw422016.N().S(`

<div class="index">

    <article>
        <header>
            <h1 class="entry-title">Search: `)
//line views/ybs/search.qtpl:7
	qw422016.E().S(p.Q)
//line views/ybs/search.qtpl:7
	qw422016.N().S(`</h1>
        </header>
    </article>

    `)
//line views/ybs/search.qtpl:11
	for _, item := range p.TopicPageInfo.Items {
//line views/ybs/search.qtpl:11
		qw422016.N().S(`
    <article>

        <header>
            `)
//line views/ybs/search.qtpl:15
		if item.Comments > 0 {
//line views/ybs/search.qtpl:15
			qw422016.N().S(`
            <a href="/t/`)
//line views/ybs/search.qtpl:16
			qw422016.N().DUL(item.ID)
//line views/ybs/search.qtpl:16
			qw422016.N().S(`#r`)
//line views/ybs/search.qtpl:16
			qw422016.N().DUL(item.Comments)
//line views/ybs/search.qtpl:16
			qw422016.N().S(`"><img alt="`)
//line views/ybs/search.qtpl:16
			qw422016.E().S(item.Title)
//line views/ybs/search.qtpl:16
			qw422016.N().S(` icon" src="/icon/t/`)
//line views/ybs/search.qtpl:16
			qw422016.N().DUL(item.ID)
//line views/ybs/search.qtpl:16
			qw422016.N().S(`.jpg?r=`)
//line views/ybs/search.qtpl:16
			qw422016.N().DUL(item.Comments)
//line views/ybs/search.qtpl:16
			qw422016.N().S(`" class="avatar"></a>
            `)
//line views/ybs/search.qtpl:17
		} else {
//line views/ybs/search.qtpl:17
			qw422016.N().S(`
            <a href="/t/`)
//line views/ybs/search.qtpl:18
			qw422016.N().DUL(item.ID)
//line views/ybs/search.qtpl:18
			qw422016.N().S(`"><img alt="`)
//line views/ybs/search.qtpl:18
			qw422016.E().S(item.Title)
//line views/ybs/search.qtpl:18
			qw422016.N().S(` icon" src="/static/avatar/`)
//line views/ybs/search.qtpl:18
			qw422016.N().DUL(item.UserId)
//line views/ybs/search.qtpl:18
			qw422016.N().S(`.jpg" class="avatar"></a>
            `)
//line views/ybs/search.qtpl:19
		}
//line views/ybs/search.qtpl:19
		qw422016.N().S(`
            <h1 class="entry-title"><a href="/t/`)
//line views/ybs/search.qtpl:20
		qw422016.N().DUL(item.ID)
//line views/ybs/search.qtpl:20
		qw422016.N().S(`" rel="bookmark" title="Permanent Link to `)
//line views/ybs/search.qtpl:20
		qw422016.E().S(item.Title)
//line views/ybs/search.qtpl:20
		qw422016.N().S(`">`)
//line views/ybs/search.qtpl:20
		qw422016.E().S(item.Title)
//line views/ybs/search.qtpl:20
		qw422016.N().S(`</a></h1>
            <p class="meta">
                <time datetime="`)
//line views/ybs/search.qtpl:22
		qw422016.E().S(item.AddTimeFmt)
//line views/ybs/search.qtpl:22
		qw422016.N().S(`" pubdate data-updated="true">`)
//line views/ybs/search.qtpl:22
		qw422016.E().S(item.EditTimeFmt)
//line views/ybs/search.qtpl:22
		qw422016.N().S(`</time>
                in <a href="/n/`)
//line views/ybs/search.qtpl:23
		qw422016.N().DUL(item.NodeId)
//line views/ybs/search.qtpl:23
		qw422016.N().S(`" rel="bookmark">`)
//line views/ybs/search.qtpl:23
		qw422016.E().S(item.NodeName)
//line views/ybs/search.qtpl:23
		qw422016.N().S(`</a>
                by <a href="/member/`)
//line views/ybs/search.qtpl:24
		qw422016.N().DUL(item.UserId)
//line views/ybs/search.qtpl:24
		qw422016.N().S(`" rel="nofollow">`)
//line views/ybs/search.qtpl:24
		qw422016.E().S(item.AuthorName)
//line views/ybs/search.qtpl:24
		qw422016.N().S(`</a>
                `)
//line views/ybs/search.qtpl:25
		if item.Comments > 0 {
//line views/ybs/search.qtpl:25
			qw422016.N().S(`
                <a class="count" href="/t/`)
//line views/ybs/search.qtpl:26
			qw422016.N().DUL(item.ID)
//line views/ybs/search.qtpl:26
			qw422016.N().S(`#r`)
//line views/ybs/search.qtpl:26
			qw422016.N().DUL(item.Comments)
//line views/ybs/search.qtpl:26
			qw422016.N().S(`" title="Comment on `)
//line views/ybs/search.qtpl:26
			qw422016.E().S(item.Title)
//line views/ybs/search.qtpl:26
			qw422016.N().S(`">`)
//line views/ybs/search.qtpl:26
			qw422016.N().DUL(item.Comments)
//line views/ybs/search.qtpl:26
			qw422016.N().S(`</a>
                `)
//line views/ybs/search.qtpl:27
		}
//line views/ybs/search.qtpl:27
		qw422016.N().S(`
            </p>
        </header>

        <div class="entry-content">
            `)
//line views/ybs/search.qtpl:32
		qw422016.E().S(item.FirstCon)
//line views/ybs/search.qtpl:32
		qw422016.N().S(`
        </div>

    </article>

    `)
//line views/ybs/search.qtpl:37
	}
//line views/ybs/search.qtpl:37
	qw422016.N().S(`

</div>

`)
//line views/ybs/search.qtpl:41
}

//line views/ybs/search.qtpl:41
func (p *SearchPage) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/ybs/search.qtpl:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/ybs/search.qtpl:41
	p.StreamMainBody(qw422016)
//line views/ybs/search.qtpl:41
	qt422016.ReleaseWriter(qw422016)
//line views/ybs/search.qtpl:41
}

//line views/ybs/search.qtpl:41
func (p *SearchPage) MainBody() string {
//line views/ybs/search.qtpl:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/ybs/search.qtpl:41
	p.WriteMainBody(qb422016)
//line views/ybs/search.qtpl:41
	qs422016 := string(qb422016.B)
//line views/ybs/search.qtpl:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/ybs/search.qtpl:41
	return qs422016
//line views/ybs/search.qtpl:41
}
