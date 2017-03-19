app.controller('sidebarCtrl', function($scope, $location){
    $scope.sidebars = [
        {titlename: '用户', navs: [{url: '', name: '所有用户'}, {url: '/admin', name: '管理员'}]},
        {titlename: '诗词', navs: [{url: '/shi', name: '诗'},{url: '', name: '歌'},{url: '', name: '词'}]}
    ]

    $scope.active = "所有用户";

    $scope.activeClass = function(nav){
        if (nav.name == $scope.active) {
            return "active"
        }
        return ""
    }

    $scope.url = function(nav){
        $scope.active = nav.name;
        $location.path(nav.url);
    }
})

app.controller('usersCtrl', function($scope) {
    $scope.title = "所有用户"
});

app.controller('adminCtrl', function($scope){

})

app.controller('shiCtrl', function($scope, $resource){
    $scope.data = {
        rows: []
    }

    var Doc = $resource('/admin/shige/');

    Doc.get(function(json){
        $scope.data.rows = json.data
    })

})