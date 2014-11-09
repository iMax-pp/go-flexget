angular.module('goFlexGet', ['ngRoute'])
.config(function($routeProvider) {
    $routeProvider
    .when('/', {
        controller:'MainCtrl',
        templateUrl:'/main.html'
    })
    .when('/config', {
        controller:'ConfigCtrl',
        templateUrl:'/config.html'
    })
    .when('/logs', {
        controller:'LogsCtrl',
        templateUrl:'/logs.html'
    })
    .otherwise({
        redirectTo:'/'
    });
})

.controller('HeaderCtrl', ['$scope', '$location', function($scope, $location) {
    $scope.isActive = function(viewLocation) {
        return viewLocation === $location.path();
    };
}])

.controller('MainCtrl', ['$scope', function($scope) {
}])

.controller('ConfigCtrl', ['$scope', '$http', '$sce', function($scope, $http, $sce) {
    $http.get('/api/config').
    success(function(data, status, headers, config) {
        $scope.flexgetConfig = data;
    }).
    error(function(data, status, headers, config) {
        $scope.retrieveError = $sce.trustAsHtml('<strong>Unable to retrieve FlexGet configuration:</strong> ' + data);
    });
}])

.controller('LogsCtrl', ['$scope', '$http', '$sce', function($scope, $http, $sce) {
    $http.get('/api/logs').
    success(function(data, status, headers, config) {
        $scope.flexgetLogs = data;
        $('#logs').animate({'scrollTop': $('#logs')[0].scrollHeight}, 100);
    }).
    error(function(data, status, headers, config) {
        $scope.retrieveError = $sce.trustAsHtml('<strong>Unable to retrieve FlexGet logs:</strong> ' + data);
    });
}]);
