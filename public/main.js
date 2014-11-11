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
        $scope.statusLoading = true;
        $http.get('/api/status')
            .success(function(data) {
                $scope.statusLoading = false;
                $scope.fgStatus = data;
            })
            .error(function(data, status) {
                $scope.statusLoading = false;
                var data = data || "Request failed";
                $scope.configError = $sce.trustAsHtml(
                    '<strong>Unable to retrieve FlexGet status:</strong> ' + data +
                    ' (' + status + ')');
            });
    }
]);
