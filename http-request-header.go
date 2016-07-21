package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

var (
	httpAddr = flag.String("http", ":8081", "http listen address")

	index = template.Must(template.New("index").Parse(indexHtml))
)

type header struct {
	Key, Value string
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var headers = make([]header, 0, len(req.Header))
		for k, _ := range req.Header {
			headers = append(headers, header{k, req.Header.Get(k)})
		}
		index.Execute(w, headers)
	})

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

const indexHtml = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="../../favicon.ico">

    <title>显示http请求的头信息</title>

    <!-- Bootstrap core CSS -->
    <link href="https://dev.imaicloud.com/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet">

    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <link href="https://dev.imaicloud.com/bootstrap/assets/css/ie10-viewport-bug-workaround.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="https://dev.imaicloud.com/iam-web/signin.css" rel="stylesheet">

    <!-- Just for debugging purposes. Don't actually copy these 2 lines! -->
    <!--[if lt IE 9]><script src="https://dev.imaicloud.com/bootstrap/assets/js/ie8-responsive-file-warning.js"></script><![endif]-->
    <script src="https://dev.imaicloud.com/bootstrap/assets/js/ie-emulation-modes-warning.js"></script>

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>

  <body>

    <div class="container">
	<table>
	{{range .}}
		<tr><td><strong>{{.Key}}</strong</td><td>{{.Value}}</td></tr>
	{{end}}
	</table>
	
    </div> <!-- /container -->


    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="https://dev.imaicloud.com/bootstrap/assets/js/ie10-viewport-bug-workaround.js"></script>
</body>
</html>
`
