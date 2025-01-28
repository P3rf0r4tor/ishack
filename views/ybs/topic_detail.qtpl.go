// Code generated by qtc from "topic_detail.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/ybs/topic_detail.qtpl:1
package ybs

//line views/ybs/topic_detail.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/ybs/topic_detail.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/ybs/topic_detail.qtpl:1
func (p *TopicDetailPage) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/ybs/topic_detail.qtpl:1
	qw422016.N().S(`
<div class="entry">

    <article role="article">

        <header>
            <a href="/member/`)
//line views/ybs/topic_detail.qtpl:7
	qw422016.N().DUL(p.TopicFmt.UserId)
//line views/ybs/topic_detail.qtpl:7
	qw422016.N().S(`" rel="nofollow"><img alt="`)
//line views/ybs/topic_detail.qtpl:7
	qw422016.E().S(p.TopicFmt.Name)
//line views/ybs/topic_detail.qtpl:7
	qw422016.N().S(` avatar" src="/static/avatar/`)
//line views/ybs/topic_detail.qtpl:7
	qw422016.N().DUL(p.TopicFmt.UserId)
//line views/ybs/topic_detail.qtpl:7
	qw422016.N().S(`.jpg" class="avatar"></a>
            <h1 class="entry-title">
                <a href="/t/`)
//line views/ybs/topic_detail.qtpl:9
	qw422016.N().DUL(p.TopicFmt.ID)
//line views/ybs/topic_detail.qtpl:9
	qw422016.N().S(`" rel="bookmark">`)
//line views/ybs/topic_detail.qtpl:9
	qw422016.E().S(p.TopicFmt.Title)
//line views/ybs/topic_detail.qtpl:9
	qw422016.N().S(`</a>
            </h1>
            <p class="meta">
                <span class="categories">📁 <a class="category" href="/n/`)
//line views/ybs/topic_detail.qtpl:12
	qw422016.N().DUL(p.TopicFmt.NodeId)
//line views/ybs/topic_detail.qtpl:12
	qw422016.N().S(`" rel="category tag">`)
//line views/ybs/topic_detail.qtpl:12
	qw422016.E().S(p.DefaultNode.Name)
//line views/ybs/topic_detail.qtpl:12
	qw422016.N().S(`</a></span>
                `)
//line views/ybs/topic_detail.qtpl:13
	qw422016.N().S(p.TopicFmt.ClockEmoji)
//line views/ybs/topic_detail.qtpl:13
	qw422016.N().S(` <time datetime="`)
//line views/ybs/topic_detail.qtpl:13
	qw422016.E().S(p.TopicFmt.AddTimeFmt)
//line views/ybs/topic_detail.qtpl:13
	qw422016.N().S(`" pubdate data-updated="true">`)
//line views/ybs/topic_detail.qtpl:13
	qw422016.E().S(p.TopicFmt.AddTimeFmt)
//line views/ybs/topic_detail.qtpl:13
	qw422016.N().S(`</time>
                by <a href="/member/`)
//line views/ybs/topic_detail.qtpl:14
	qw422016.N().DUL(p.TopicFmt.UserId)
//line views/ybs/topic_detail.qtpl:14
	qw422016.N().S(`" rel="nofollow">`)
//line views/ybs/topic_detail.qtpl:14
	qw422016.E().S(p.TopicFmt.Name)
//line views/ybs/topic_detail.qtpl:14
	qw422016.N().S(`</a>
                `)
//line views/ybs/topic_detail.qtpl:15
	if p.CurrentUser.Flag >= 99 {
//line views/ybs/topic_detail.qtpl:15
		qw422016.N().S(`
                &bull; <a rel="bookmark" href="/admin/topic/edit?id=`)
//line views/ybs/topic_detail.qtpl:16
		qw422016.N().DUL(p.TopicFmt.ID)
//line views/ybs/topic_detail.qtpl:16
		qw422016.N().S(`&back=here">Edit</a>
                `)
//line views/ybs/topic_detail.qtpl:17
	}
//line views/ybs/topic_detail.qtpl:17
	qw422016.N().S(`
            </p>
        </header>

        <div class="markdown-body entry-content">
            `)
//line views/ybs/topic_detail.qtpl:22
	qw422016.N().S(p.TopicFmt.ContentFmt)
//line views/ybs/topic_detail.qtpl:22
	qw422016.N().S(`
        </div>

        `)
//line views/ybs/topic_detail.qtpl:25
	if len(p.TopicFmt.Relative) > 0 {
//line views/ybs/topic_detail.qtpl:25
		qw422016.N().S(`
        <section>
            <h4 class="seealso-title">💘 相关文章</h4>
            <ul class="seealso">
                `)
//line views/ybs/topic_detail.qtpl:29
		for _, v := range p.TopicFmt.Relative {
//line views/ybs/topic_detail.qtpl:29
			qw422016.N().S(`
                <li><a href="/t/`)
//line views/ybs/topic_detail.qtpl:30
			qw422016.N().DUL(v.ID)
//line views/ybs/topic_detail.qtpl:30
			qw422016.N().S(`" rel="bookmark">`)
//line views/ybs/topic_detail.qtpl:30
			qw422016.E().S(v.Title)
//line views/ybs/topic_detail.qtpl:30
			qw422016.N().S(`</a></li>
                `)
//line views/ybs/topic_detail.qtpl:31
		}
//line views/ybs/topic_detail.qtpl:31
		qw422016.N().S(`
            </ul>
        </section>
        `)
//line views/ybs/topic_detail.qtpl:34
	}
//line views/ybs/topic_detail.qtpl:34
	qw422016.N().S(`

        <footer>

            <p class="meta gray">
                <span class="categories">
                    📁 Category: <a class="category" href="/n/`)
//line views/ybs/topic_detail.qtpl:40
	qw422016.N().DUL(p.TopicFmt.NodeId)
//line views/ybs/topic_detail.qtpl:40
	qw422016.N().S(`" rel="category tag">`)
//line views/ybs/topic_detail.qtpl:40
	qw422016.E().S(p.DefaultNode.Name)
//line views/ybs/topic_detail.qtpl:40
	qw422016.N().S(`</a>
                </span>
                `)
//line views/ybs/topic_detail.qtpl:42
	if len(p.TagLst) > 0 {
//line views/ybs/topic_detail.qtpl:42
		qw422016.N().S(`
                <span class="categories">🏷️ Tags:
                    `)
//line views/ybs/topic_detail.qtpl:44
		for _, tag := range p.TagLst {
//line views/ybs/topic_detail.qtpl:44
			qw422016.N().S(`
                    <a class="tag" href="/tag/`)
//line views/ybs/topic_detail.qtpl:45
			qw422016.N().U(tag.Name)
//line views/ybs/topic_detail.qtpl:45
			qw422016.N().S(`" rel="tag">`)
//line views/ybs/topic_detail.qtpl:45
			qw422016.E().S(tag.Name)
//line views/ybs/topic_detail.qtpl:45
			qw422016.N().S(`</a>
                    `)
//line views/ybs/topic_detail.qtpl:46
		}
//line views/ybs/topic_detail.qtpl:46
		qw422016.N().S(`
                </span>
                `)
//line views/ybs/topic_detail.qtpl:48
	}
//line views/ybs/topic_detail.qtpl:48
	qw422016.N().S(`
                <span class="categories">
                💬 <a href="#comments">Comments (`)
//line views/ybs/topic_detail.qtpl:50
	qw422016.N().DUL(p.TopicFmt.Comments)
//line views/ybs/topic_detail.qtpl:50
	qw422016.N().S(`)</a> 😊 PageView (`)
//line views/ybs/topic_detail.qtpl:50
	qw422016.N().DUL(p.TopicFmt.Views)
//line views/ybs/topic_detail.qtpl:50
	qw422016.N().S(`)
                </span>
            </p>

            <ul class="seealso">
                `)
//line views/ybs/topic_detail.qtpl:55
	if p.NewTopic.ID > 0 {
//line views/ybs/topic_detail.qtpl:55
		qw422016.N().S(`
                <li>上一篇 › <a class="next" href="/t/`)
//line views/ybs/topic_detail.qtpl:56
		qw422016.N().DUL(p.NewTopic.ID)
//line views/ybs/topic_detail.qtpl:56
		qw422016.N().S(`" rel="next">`)
//line views/ybs/topic_detail.qtpl:56
		qw422016.E().S(p.NewTopic.Title)
//line views/ybs/topic_detail.qtpl:56
		qw422016.N().S(`</a></li>
                `)
//line views/ybs/topic_detail.qtpl:57
	}
//line views/ybs/topic_detail.qtpl:57
	qw422016.N().S(`
                `)
//line views/ybs/topic_detail.qtpl:58
	if p.OldTopic.ID > 0 {
//line views/ybs/topic_detail.qtpl:58
		qw422016.N().S(`
                <li>下一篇 › <a class="prev" href="/t/`)
//line views/ybs/topic_detail.qtpl:59
		qw422016.N().DUL(p.OldTopic.ID)
//line views/ybs/topic_detail.qtpl:59
		qw422016.N().S(`" rel="prev">`)
//line views/ybs/topic_detail.qtpl:59
		qw422016.E().S(p.OldTopic.Title)
//line views/ybs/topic_detail.qtpl:59
		qw422016.N().S(`</a></li>
                `)
//line views/ybs/topic_detail.qtpl:60
	}
//line views/ybs/topic_detail.qtpl:60
	qw422016.N().S(`
            </ul>

        </footer>

    </article>

    `)
//line views/ybs/topic_detail.qtpl:67
	if len(p.CommentLst) > 0 {
//line views/ybs/topic_detail.qtpl:67
		qw422016.N().S(`
    <section>
        <h1 class="br-nav">评论</h1>
        <div id="comments">
            <h2 class="comments bot-line">共`)
//line views/ybs/topic_detail.qtpl:71
		qw422016.N().DUL(p.TopicFmt.Comments)
//line views/ybs/topic_detail.qtpl:71
		qw422016.N().S(`条关于"`)
//line views/ybs/topic_detail.qtpl:71
		qw422016.E().S(p.TopicFmt.Title)
//line views/ybs/topic_detail.qtpl:71
		qw422016.N().S(`"的评论</h2>
            `)
//line views/ybs/topic_detail.qtpl:72
		for _, item := range p.CommentLst {
//line views/ybs/topic_detail.qtpl:72
			qw422016.N().S(`
            <article id="r`)
//line views/ybs/topic_detail.qtpl:73
			qw422016.N().DUL(item.ID)
//line views/ybs/topic_detail.qtpl:73
			qw422016.N().S(`">
                <header>
                    <a href="/member/`)
//line views/ybs/topic_detail.qtpl:75
			qw422016.N().DUL(item.UserId)
//line views/ybs/topic_detail.qtpl:75
			qw422016.N().S(`" rel="nofollow"><img alt="`)
//line views/ybs/topic_detail.qtpl:75
			qw422016.E().S(item.Name)
//line views/ybs/topic_detail.qtpl:75
			qw422016.N().S(` avatar" src="/static/avatar/`)
//line views/ybs/topic_detail.qtpl:75
			qw422016.N().DUL(item.UserId)
//line views/ybs/topic_detail.qtpl:75
			qw422016.N().S(`.jpg" class="avatar"></a>
                    <div class="meta gray5">
                        <a href="`)
//line views/ybs/topic_detail.qtpl:77
			qw422016.E().S(item.Link)
//line views/ybs/topic_detail.qtpl:77
			qw422016.N().S(`" class="comment-count comment-id">#`)
//line views/ybs/topic_detail.qtpl:77
			qw422016.N().DUL(item.ID)
//line views/ybs/topic_detail.qtpl:77
			qw422016.N().S(`</a> <a href="/member/`)
//line views/ybs/topic_detail.qtpl:77
			qw422016.N().DUL(item.UserId)
//line views/ybs/topic_detail.qtpl:77
			qw422016.N().S(`" rel="nofollow">`)
//line views/ybs/topic_detail.qtpl:77
			qw422016.E().S(item.Name)
//line views/ybs/topic_detail.qtpl:77
			qw422016.N().S(`</a>
                        `)
//line views/ybs/topic_detail.qtpl:78
			if p.CurrentUser.Flag >= 99 {
//line views/ybs/topic_detail.qtpl:78
				qw422016.N().S(`
                        <a href="/admin/comment/edit?tid=`)
//line views/ybs/topic_detail.qtpl:79
				qw422016.N().DUL(item.TopicId)
//line views/ybs/topic_detail.qtpl:79
				qw422016.N().S(`&cid=`)
//line views/ybs/topic_detail.qtpl:79
				qw422016.N().DUL(item.ID)
//line views/ybs/topic_detail.qtpl:79
				qw422016.N().S(`&back=here">Edit</a>
                        `)
//line views/ybs/topic_detail.qtpl:80
			}
//line views/ybs/topic_detail.qtpl:80
			qw422016.N().S(`
                        <a rel="nofollow" class="right" href="#respond" onclick="replyTo('`)
//line views/ybs/topic_detail.qtpl:81
			qw422016.E().S(item.Name)
//line views/ybs/topic_detail.qtpl:81
			qw422016.N().S(`',`)
//line views/ybs/topic_detail.qtpl:81
			qw422016.N().DUL(item.ID)
//line views/ybs/topic_detail.qtpl:81
			qw422016.N().S(`)">回复</a>
                        <br>
                        <time datetime="`)
//line views/ybs/topic_detail.qtpl:83
			qw422016.N().DL(item.AddTime)
//line views/ybs/topic_detail.qtpl:83
			qw422016.N().S(`" pubdate data-updated="true"><em>`)
//line views/ybs/topic_detail.qtpl:83
			qw422016.E().S(item.AddTimeFmt)
//line views/ybs/topic_detail.qtpl:83
			qw422016.N().S(`</em></time>
                    </div>
                </header>
                <div class="markdown-body entry-content">
                    `)
//line views/ybs/topic_detail.qtpl:87
			qw422016.N().S(item.ContentFmt)
//line views/ybs/topic_detail.qtpl:87
			qw422016.N().S(`
                </div>
            </article>
            `)
//line views/ybs/topic_detail.qtpl:90
		}
//line views/ybs/topic_detail.qtpl:90
		qw422016.N().S(`
        </div>
    </section>
    `)
//line views/ybs/topic_detail.qtpl:93
	}
//line views/ybs/topic_detail.qtpl:93
	qw422016.N().S(`

    <section id="respond">
        <h2 class="br-nav">写一条评论</h2>
        <div class="write-comment pure-form">
            <form action="" method="post" id="commentform">
                <textarea name="comment" id="id-comment" class="pure-u-1" placeholder="* 字符限制 `)
//line views/ybs/topic_detail.qtpl:99
	qw422016.N().D(p.SiteCf.CommentConMaxLen)
//line views/ybs/topic_detail.qtpl:99
	qw422016.N().S(`"></textarea>
                `)
//line views/ybs/topic_detail.qtpl:100
	if p.CurrentUser.ID > 0 {
//line views/ybs/topic_detail.qtpl:100
		qw422016.N().S(`
                <div class="pure-button-group">
                `)
//line views/ybs/topic_detail.qtpl:102
		if p.SiteCf.CloseReply && p.CurrentUser.Flag < 99 {
//line views/ybs/topic_detail.qtpl:102
			qw422016.N().S(`
                    <input id="btn-preview" type="button" value="评论已关闭" name="submit" onclick="return false;" class="pure-button" disabled="" />
                `)
//line views/ybs/topic_detail.qtpl:104
		} else {
//line views/ybs/topic_detail.qtpl:104
			qw422016.N().S(`
                    <input id="btn-preview" type="button" value="预览" name="submit" onclick="previewComment(); return false;" class="pure-button button-success" />
                    <input id="btn-submit" type="button" value="发表" name="submit" onclick="submitComment(); return false;" class="pure-button pure-button-primary" />
                    `)
//line views/ybs/topic_detail.qtpl:107
			if !p.SiteCf.UploadLimit || (p.SiteCf.UploadLimit && p.CurrentUser.Flag >= 99) {
//line views/ybs/topic_detail.qtpl:107
				qw422016.N().S(`
                    <input id="fileUpload" type="file" accept="image/*,video/*,audio/*" onChange="uploadFile()" class="pure-button" name="fileUpload" style="font-size: .8334em;width: 95px;" />
                    `)
//line views/ybs/topic_detail.qtpl:109
			}
//line views/ybs/topic_detail.qtpl:109
			qw422016.N().S(`
                    <button id="insert-break" type="button" class="pure-button">插入分割线</button>
                `)
//line views/ybs/topic_detail.qtpl:111
		}
//line views/ybs/topic_detail.qtpl:111
		qw422016.N().S(`
                </div>
                <span id="id-msg"></span>
                `)
//line views/ybs/topic_detail.qtpl:114
	} else {
//line views/ybs/topic_detail.qtpl:114
		qw422016.N().S(`
                <a href="/login" rel="nofollow" class="pure-button">登录发表评论</a>
                `)
//line views/ybs/topic_detail.qtpl:116
	}
//line views/ybs/topic_detail.qtpl:116
	qw422016.N().S(`
            </form>
        </div>
        <div id="id-preview" class="markdown-body entry-content"></div>

        <script>

            var toReplyId = 0;
            var conEle = document.getElementById("id-comment");
            var msgEle = document.getElementById("id-msg");
            var reviewEle = document.getElementById("id-preview");

            `)
//line views/ybs/topic_detail.qtpl:128
	if p.CurrentUser.ID > 0 {
//line views/ybs/topic_detail.qtpl:128
		qw422016.N().S(`
                document.getElementById("insert-break").addEventListener('click', function (event) {
                    let break_line = "\n`)
//line views/ybs/topic_detail.qtpl:130
		qw422016.N().S(p.ReadMoreBreak)
//line views/ybs/topic_detail.qtpl:130
		qw422016.N().S(`\n";
                    let pos = conEle.selectionStart;
                    let con = conEle.value;
                    conEle.value = con.slice(0, pos) + break_line + con.slice(pos);
                }, false);
                function previewComment() {
                    var con = conEle.value.trim();
                    if (con === "") {
                        conEle.focus();
                        return
                    }
                    postAjax("/content/preview", JSON.stringify({Act: "commentPreview", Content: con}), function(data){
                        var obj = JSON.parse(data)
                        //console.log(obj);
                        if(obj.Code === 200) {
                            msgEle.style.display = "none";
                            reviewEle.innerHTML = obj.Html;
                            reviewEle.style.display = "block";
                        }else{
                            reviewEle.innerHTML = "";
                            reviewEle.style.display = "none";
                            msgEle.innerText = obj.Msg;
                        }
                    });
                }
                function submitComment() {
                    var con = conEle.value.trim();
                    if (con === "") {
                        conEle.focus();
                        return
                    }
                    postAjax("/t/`)
//line views/ybs/topic_detail.qtpl:161
		qw422016.N().DUL(p.TopicFmt.ID)
//line views/ybs/topic_detail.qtpl:161
		qw422016.N().S(`", JSON.stringify({Content: con, ReplyId: toReplyId}), function(data){
                        var obj = JSON.parse(data)
                        msgEle.innerText = obj.Msg;
                        conEle.focus();
                        conEle.value = "";
                        toReplyId = 0;
                        if(obj.Code === 200) {
                            window.location.href = "/t/`)
//line views/ybs/topic_detail.qtpl:168
		qw422016.N().DUL(p.TopicFmt.ID)
//line views/ybs/topic_detail.qtpl:168
		qw422016.N().S(`#r"+obj.Tid;
                            window.location.reload(true);
                            return;
                        } else if (obj.Code === 201) {
                            window.location.href = "/member/`)
//line views/ybs/topic_detail.qtpl:172
		qw422016.N().DUL(p.CurrentUser.ID)
//line views/ybs/topic_detail.qtpl:172
		qw422016.N().S(`?type=comment";
                            return;
                        }
                        reviewEle.style.display = "none";
                        msgEle.style.display = "block";
                    });
                }
                `)
//line views/ybs/topic_detail.qtpl:179
		if !p.SiteCf.UploadLimit || (p.SiteCf.UploadLimit && p.CurrentUser.Flag >= 99) {
//line views/ybs/topic_detail.qtpl:179
			qw422016.N().S(`
                document.addEventListener('paste', function (evt) {
                    var url = "/file/upload";
                    var items = evt.clipboardData && evt.clipboardData.items;
                    var file = null;
                    if(items && items.length) {
                        for(var i=0; i!==items.length; i++) {
                            var iType = items[i].type;
                            if(iType.indexOf('image') !== -1 || iType.indexOf('video') !== -1 || iType.indexOf('audio') !== -1) {
                                file = items[i].getAsFile();
                                if(!!!file) {
                                    continue;
                                }

                                // upload file object.
                                var form = new FormData();
                                form.append('file', file);

                                postAjax("/file/upload", form, function(data){
                                    let obj = JSON.parse(data)
                                    //console.log(obj);
                                    if(obj.Code === 200) {
                                        let img_url = "\n" + s2tag(obj.Url, `)
//line views/ybs/topic_detail.qtpl:201
			qw422016.E().V(p.SiteCf.AutoDecodeMp4)
//line views/ybs/topic_detail.qtpl:201
			qw422016.N().S(`) + "\n";
                                        let pos = conEle.selectionStart;
                                        let con = conEle.value;
                                        conEle.value = con.slice(0, pos) + img_url + con.slice(pos);
                                    }else{
                                        console.warn(obj.Msg);
                                    }
                                });
                            }
                        }
                    }
                });
                function uploadFile() {
                    let form = new FormData();
                    form.append("file", fileUpload.files[0]);
                    postAjax("/file/upload", form, function(data){
                        let obj = JSON.parse(data)
                        if(obj.Code === 200) {
                            let img_url = "\n" + s2tag(obj.Url, `)
//line views/ybs/topic_detail.qtpl:219
			qw422016.E().V(p.SiteCf.AutoDecodeMp4)
//line views/ybs/topic_detail.qtpl:219
			qw422016.N().S(`) + "\n";
                            let pos = conEle.selectionStart;
                            let con = conEle.value;
                            conEle.value = con.slice(0, pos) + img_url + con.slice(pos);
                        }else{
                            console.warn(obj.Msg);
                        }
                    });
                }
                `)
//line views/ybs/topic_detail.qtpl:228
		}
//line views/ybs/topic_detail.qtpl:228
		qw422016.N().S(`
            `)
//line views/ybs/topic_detail.qtpl:229
	}
//line views/ybs/topic_detail.qtpl:229
	qw422016.N().S(`

            function replyTo(name, cid) {
                toReplyId = parseInt(cid, 10);
                var con = conEle.value;
                document.getElementsByTagName('textarea')[0].focus();
                conEle.value = " @"+name+" #" + cid + " " + con;
                return false
            }

            function linkClick() {
                //console.log("ele clicked:"+this.href);
                var curEle = this;
                postAjax("/get/link/count", JSON.stringify({Act: "set", Items: [this.href]}), function(data){
                    var obj = JSON.parse(data)
                    if(obj.Code===200) {
                        var num = obj.Num;
                        var nextEle = curEle.nextElementSibling;
                        if(nextEle && nextEle.nodeType === 1) {
                            nextEle.innerHTML = num;
                            nextEle.title = num + "次点击"
                        }else{
                            var newNode = document.createElement('span');
                            newNode.innerHTML = num;
                            newNode.classList.toggle("clicks");
                            newNode.title = num + "次点击"
                            curEle.after(newNode);
                        }
                    }
                });
            }

            function getContentLinkCount() {
                var urlEleDict = {};
                var linkDict = {};
                var conLst = document.querySelectorAll(".entry-content a");
                for (var i = 0, max = conLst.length; i < max; i++) {
                    var aEle = conLst[i];
                    if(!aEle.classList.contains('anchor') && !aEle.getAttribute("href").startsWith('/name/')) {
                        var fullLink = aEle.href;
                        linkDict[aEle.href] = null;
                        if (urlEleDict.hasOwnProperty(fullLink)) {
                            urlEleDict[fullLink].push(aEle)
                        }else{
                            urlEleDict[fullLink] = [aEle];
                        }
                        if(aEle.getAttribute("target")==="_blank"){
                            aEle.classList.toggle("external_link");
                        }
                        aEle.addEventListener("click", linkClick);
                    }
                }
                var linkArray = Object.keys(linkDict);
                if(linkArray.length > 0) {
                    postAjax("/get/link/count", JSON.stringify({Items: linkArray}), function(data){
                        var obj = JSON.parse(data)
                        if(obj.Code===200) {
                            Object.keys(obj.Info).forEach(function(key) {
                                var num = obj.Info[key];
                                if (urlEleDict.hasOwnProperty(key)) {
                                    for (var i = 0, max = urlEleDict[key].length; i < max; i++) {
                                        var newNode = document.createElement('span');
                                        newNode.innerHTML = num;
                                        newNode.classList.toggle("clicks");
                                        newNode.title = num + "次点击"
                                        var tmpEle = urlEleDict[key][i];
                                        tmpEle.after(newNode);
                                    }
                                }
                            });
                        }
                    });
                }
            }

let audioLst = Array.from(document.querySelectorAll(".markdown-body audio"), audio =>audio.src);
if(audioLst.length>1){
    let curIndex = 0;
    const audio = document.createElement('audio');
    audio.controls = true;
    audio.src = audioLst[curIndex];

    let mainEle = document.querySelector(".markdown-body");
    mainEle.insertAdjacentElement('afterbegin', audio);

    audio.addEventListener("ended", (event) => {
        if(curIndex<audioLst.length-1){
            curIndex++;
        }else{
            curIndex = 0;
        }
        audio.src = audioLst[curIndex];
        console.log(audio.src);
        audio.play();
    });
}

            docReady(function() {
                getContentLinkCount();
            });

        </script>

    </section>

</div>

`)
//line views/ybs/topic_detail.qtpl:336
}

//line views/ybs/topic_detail.qtpl:336
func (p *TopicDetailPage) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/ybs/topic_detail.qtpl:336
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/ybs/topic_detail.qtpl:336
	p.StreamMainBody(qw422016)
//line views/ybs/topic_detail.qtpl:336
	qt422016.ReleaseWriter(qw422016)
//line views/ybs/topic_detail.qtpl:336
}

//line views/ybs/topic_detail.qtpl:336
func (p *TopicDetailPage) MainBody() string {
//line views/ybs/topic_detail.qtpl:336
	qb422016 := qt422016.AcquireByteBuffer()
//line views/ybs/topic_detail.qtpl:336
	p.WriteMainBody(qb422016)
//line views/ybs/topic_detail.qtpl:336
	qs422016 := string(qb422016.B)
//line views/ybs/topic_detail.qtpl:336
	qt422016.ReleaseByteBuffer(qb422016)
//line views/ybs/topic_detail.qtpl:336
	return qs422016
//line views/ybs/topic_detail.qtpl:336
}
