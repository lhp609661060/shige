<!DOCTYPE html>
<html lang="zh-CN" ng-app="app">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>{{.Title}}</title>
    <link rel="icon" href="/static/img/logo1.jpg">

    <link href="http://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <link href="http://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap-theme.css" rel="stylesheet">
    <link href="/static/css/dashboard.css" rel="stylesheet">

</head>
<body>

    <nav class="navbar navbar-inverse navbar-fixed-top">
        <div class="container-fluid">
            <div class="navbar-header">
                <a class="navbar-brand" href="#">每天一首诗</a>
            </div>
            <div id="navbar" class="navbar-collapse collapse">
                <ul class="nav navbar-nav navbar-right">
                    <li><a href="#">设置</a></li>
                    <li><a href="#">我的地盘</a></li>
                    <li><a href="#">来问我</a></li>
                </ul>
                <form class="navbar-form navbar-right">
                    <input type="text" class="form-control" placeholder="上知天文">
                </form>
            </div>
        </div>
    </nav>

    <div class="container-fluid">
        <div class="row">
            <div class="col-sm-3 col-md-2 sidebar" ng-controller="sidebarCtrl">
                <div ng-repeat="sidebar in sidebars">
                    <h4 ng-bind="sidebar.titlename"></h4>
                    <ul class="nav nav-sidebar">
                        <li ng-repeat="nav in sidebar.navs" ng-class="activeClass(nav)">
                            <a ng-click="url(nav)" ng-bind="nav.name"></a>
                        </li>
                    </ul>
                </div>
            </div>
            <div ng-view class="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main"></div>
        </div>
    </div>

    <script src="http://cdn.bootcss.com/angular.js/1.6.3/angular.min.js"></script>
    <script src="http://cdn.bootcss.com/angular.js/1.6.3/angular-animate.min.js"></script>
    <script src="http://cdn.bootcss.com/angular.js/1.6.3/angular-route.js"></script>
    <script src="http://cdn.bootcss.com/angular.js/1.6.3/angular-resource.min.js"></script>
    <script src="http://cdn.bootcss.com/angular.js/1.6.3/angular-loader.min.js"></script>
    <script src="/static/app/app.js"></script>
    <script src="/static/app/route.js"></script>
    <script src="/static/app/controllers.js"></script>
</body>
</html>