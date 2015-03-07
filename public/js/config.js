angular.module('goFlexGet')
    .controller('ConfigCtrl', ['$scope', '$http', '$sce', '$timeout',
        function($scope, $http, $sce, $timeout) {
            $scope.configLoading = true;
            $http.get('/api/config')
                .success(function(data) {
                    $scope.configLoading = false;
                    $scope.flexgetConfig = data;
                    // Wait for template before applying prettify
                    $timeout(function() {
                        prettyPrint();
                    });
                })
                .error(function(data, status) {
                    $scope.configLoading = false;
                    data = data || 'Request failed';
                    $scope.configError = $sce.trustAsHtml(
                        '<strong>Unable to retrieve FlexGet configuration:</strong> ' +
                        data +
                        ' (' + status + ')');
                });
        }
    ]);
