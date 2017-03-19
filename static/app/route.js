app.config(['$routeProvider', function($routeProvider){
    $routeProvider
        .when('/', {
            controller : "usersCtrl",
            templateUrl : '/static/app/tpl/users.html'
        })
        .when('/admin', {
            controller : 'adminCtrl',
            templateUrl : '/static/app/tpl/admins.html'
        })
        .when('/shi', {
            controller : 'shiCtrl',
            templateUrl : '/static/app/tpl/shi.html'
        })
        .otherwise({
            redirectTo : '/'
        });
}]);
