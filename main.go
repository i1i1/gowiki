package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	bf "github.com/russross/blackfriday/v2"
)

type Renderer struct {
	html  *bf.HTMLRenderer
	theme string
}

var (
	style string = `
header a {
	background:      #ececec;
	border-color:    #ececec;
	border-width:    5px;
	border-style:    solid;
	border-radius:   8px;
	color:           black;
	margin-left:     10px;
	margin-right:    10px;
	text-decoration: none;
}
header a:hover {
	color:        black;
	background:   #dcdcdc;
	border-color: #dcdcdc;
}
#footer {
	bottom:    0px;
	color:     #404040;
	font-size: smaller;
	position:  absolute;
}
`
	header string = `
<html>
  <head>
    <meta charset="utf-8"/>
	<link rel="stylesheet" href="https://fonts.xz.style/serve/cascadia.css"> 
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@exampledev/new.css@1.1.2/new.min.css">
    <style>` + style + `</style>
  </head>
  <body>
    <header>
      <h1>
        <a href="/">Wiki</a>
        <a href="/about">About</a>
        <a href="/index">Index</a>
      </h1>
    </header>
`
	footer string = `
    <div id="footer">
      <p>Updated at ` + time.Now().Format(time.UnixDate) + `</p>
    </div>
  </body>
</html>
`
)

// Code from here:
//   https://eddieflores.com/tech/blackfriday-chroma/
func (r *Renderer) RenderNode(w io.Writer, node *bf.Node, entering bool) bf.WalkStatus {
	if node.Type != bf.CodeBlock {
		return r.html.RenderNode(w, node, entering)
	}

	var lexer chroma.Lexer

	lang := string(node.CodeBlockData.Info)
	if lang != "" {
		lexer = lexers.Get(lang)
	} else {
		lexer = lexers.Analyse(string(node.Literal))
	}

	if lexer == nil {
		lexer = lexers.Fallback
	}

	// Set a syntax highlighting theme
	style := styles.Get(r.theme)
	if style == nil {
		style = styles.Fallback
	}

	// Apply highlighting with Chroma.
	iterator, err := lexer.Tokenise(nil, string(node.Literal))
	if err != nil {
		panic(err)
	}

	// Write out the highlighted code to the io.Writer.
	err = html.New().Format(w, style, iterator)
	if err != nil {
		panic(err)
	}

	return bf.GoToNext
}

// Leaving these blank to satisfy the Renderer interface, not useful to us.
func (r *Renderer) RenderHeader(w io.Writer, ast *bf.Node) {
	io.WriteString(w, header)
}
func (r *Renderer) RenderFooter(w io.Writer, ast *bf.Node) {
	io.WriteString(w, footer)
}

func NewRenderer(theme string) *Renderer {
	return &Renderer{
		html:  bf.NewHTMLRenderer(bf.HTMLRendererParameters{}),
		theme: theme,
	}
}

func main() {
	markdown, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	wr := bufio.NewWriter(os.Stdout)
	opt := bf.WithRenderer(NewRenderer("solarized-light"))
	wr.Write(bf.Run(markdown, opt))
	wr.Flush()
}
