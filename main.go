package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

var style string = `
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
`

var header string = `
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
var footer string = `
    <p>Updated at ` + time.Now().Format(time.UnixDate) + `</p>
  </body>
</html>
`

func md2html(md []byte) []byte {
	unsafe := blackfriday.Run(md)
	return bluemonday.UGCPolicy().SanitizeBytes(unsafe)
}

func wikipedize(md []byte) []byte {
	return []byte(header + string(md2html(md)) + footer)
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	wr := bufio.NewWriter(os.Stdout)
	wr.Write(wikipedize(bytes))
	wr.Flush()
}
