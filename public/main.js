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

.controller('MainCtrl', ['$scope',
    function($scope) {}
])

.controller('ConfigCtrl', ['$scope', '$http', '$sce', '$timeout',
    function($scope, $http, $sce, $timeout) {
        $scope.configLoading = true
        $http.get('/api/config')
            .success(function(data) {
                $scope.configLoading = false
                $scope.flexgetConfig = data;
                // Wait for template before applying prettify
                $timeout(function() {
                    prettyPrint();
                });
            })
            .error(function(data, status) {
                $scope.configLoading = false
                var data = data || "Request failed";
                $scope.configError = $sce.trustAsHtml(
                    '<strong>Unable to retrieve FlexGet configuration:</strong> ' + data +
                    ' (' + status + ')');
            });
    }
])

.controller('LogsCtrl', ['$scope', '$http', '$sce', '$timeout',
    function($scope, $http, $sce, $timeout) {
        $scope.logsLoading = true
        $http.get('/api/logs')
            .success(function(data) {
                $scope.logsLoading = false
                $scope.flexgetLogs = data;
                // Wait for template before scrolling to bottom
                $timeout(function() {
                    $('#logs').animate({
                        'scrollTop': $('#logs')[0].scrollHeight
                    }, 100);
                });
            })
            .error(function(data, status) {
                $scope.logsLoading = false
                var data = data || "Request failed";
                $scope.logsError = $sce.trustAsHtml(
                    '<strong>Unable to retrieve FlexGet logs:</strong> ' + data + ' (' +
                    status + ')');
            });
    }
]);
