angular.module('goFlexGet', ['ngRoute'])
    .config(function($routeProvider) {
        $routeProvider
            .when('/', {
                controller: 'MainCtrl',
                templateUrl: '/main.html'
            })
            .when('/config', {
                controller: 'ConfigCtrl',
                templateUrl: '/config.html'
            })
            .when('/logs', {
                controller: 'LogsCtrl',
                templateUrl: '/logs.html'
            })
            .otherwise({
                redirectTo: '/'
            });
    })

.controller('HeaderCtrl', ['$scope', '$location',
    function($scope, $location) {
        $scope.isActive = function(viewLocation) {
            return viewLocation === $location.path();
        };
    }
])

.controller('MainCtrl', ['$scope', '$http',
    function($scope, $http) {
        $scope.getStatus = function() {
            $scope.statusLoading = true;
            $http.get('/api/status')
                .success(function(data) {
                    $scope.statusLoading = false;
                    $scope.fgStatus = data;
                })
                .error(function(data, status) {
                    $scope.statusLoading = false;
                    var data = data || "Request failed";
                    $scope.statusError = $sce.trustAsHtml(
                        '<strong>Unable to retrieve FlexGet status:</strong> ' + data +
                        ' (' + status + ')');
                });
        };
        $scope.getStatus();

        $scope.isStarting = false;
        $scope.startFlexGet = function() {
            $scope.isStarting = true;
            $http.get('/api/start')
                .success(function(data) {
                    $scope.getStatus();
                    $scope.isStarting = false;
                })
                .error(function(data, status) {
                    var data = data || "Request failed";
                    console.error(data)
                    $scope.getStatus();
                    $scope.isStarting = false;
                });
        };

        $scope.isStopping = false;
        $scope.stopFlexGet = function() {
            $scope.isStopping = true;
            $http.get('/api/stop')
                .success(function(data) {
                    $scope.getStatus();
                    $scope.isStopping = false;
                })
                .error(function(data, status) {
                    var data = data || "Request failed";
                    console.error(data)
                    $scope.getStatus();
                    $scope.isStopping = false;
                });
        };
    }
]);
