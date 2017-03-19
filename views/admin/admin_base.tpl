<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/static/module/bootstrap/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/module/bootstrap/css/bootstrap-theme.min.css">
    {{.HtmlHead}}
</head>
<body>

<div class="container">
    {{.LayoutContent}}
</div>
<div>
    {{.SideBar}}
</div>
<script src="http://cdn.bootcss.com/jquery/3.2.0/jquery.min.js"></script>
<script type="text/javascript" src="/static/module/bootstrap/js/bootstrap.min.js"></script>
{{.Scripts}}
</body>
</html>