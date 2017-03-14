<!DOCTYPE html>
<html>
    <head>
        <title>每天一诗</title>
        <meta charset="utf-8">
        <link rel="icon" href="/static/img/shige.png">
        <link rel="stylesheet" type="text/css" href="/static/css/hello.css">
    </head>

<body>
  <div class="container">
    <div>
      <div class="photo">
        <div class="demo" id="demo" align="center">
          <div class="express">
            <h3>【{{ .ShigeDoc.Title }}】</h3>
            <p>{{ .ShigeDoc.Author }}</p>
            {{ .ShigeDoc.Doc | replace "\n" "<br>" | str2html }}
          </div>
        </div>
      </div>

    </div>
  </div>
  </body>
</html>
