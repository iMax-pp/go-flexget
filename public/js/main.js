angular.module('goFlexGet', ['ngRoute'])
    .config(function($routeProvider) {
        $routeProvider
            .when('/', {
                controller: 'MainCtrl',
                templateUrl: '/public/pages/main.html'
            })
            .when('/config', {
                controller: 'ConfigCtrl',
                templateUrl: '/public/pages/config.html'
            })
            .when('/logs', {
                controller: 'LogsCtrl',
                templateUrl: '/public/pages/logs.html'
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
                    data = data || 'Request failed';
                    $scope.statusError = $sce.trustAsHtml(
                        '<strong>Unable to retrieve FlexGet status:</strong> ' + data +
                        ' (' + status + ')');
                });
        };
        $scope.getStatus();

        $scope.isStarting = false;
        $scope.startFlexGet = function() {
            $scope.isStarting = true;
            $http.get('/api/flexget/start')
                .success(function(data) {
                    $scope.getStatus();
                    $scope.isStarting = false;
                })
                .error(function(data, status) {
                    data = data || 'Request failed';
                    console.error(data);
                    $scope.getStatus();
                    $scope.isStarting = false;
                });
        };

        $scope.isStopping = false;
        $scope.stopFlexGet = function() {
            $scope.isStopping = true;
            $http.get('/api/flexget/stop')
                .success(function(data) {
                    $scope.getStatus();
                    $scope.isStopping = false;
                })
                .error(function(data, status) {
                    data = data || 'Request failed';
                    console.error(data);
                    $scope.getStatus();
                    $scope.isStopping = false;
                });
        };

        $scope.isReloading = false;
        $scope.reloadFlexGet = function() {
            $scope.isReloading = true;
            $http.get('/api/flexget/reload')
                .success(function(data) {
                    $scope.getStatus();
                    $scope.isReloading = false;
                })
                .error(function(data, status) {
                    data = data || 'Request failed';
                    console.error(data);
                    $scope.getStatus();
                    $scope.isReloading = false;
                });
        };
    }
]);
