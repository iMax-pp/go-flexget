angular.module('goFlexGet')
    .controller('LogsCtrl', ['$scope', '$http', '$sce', '$timeout',
        function($scope, $http, $sce, $timeout) {
            $scope.getLogs = function() {
                $scope.logsLoading = true;
                $http.get('/api/logs')
                    .success(function(data) {
                        $scope.logsLoading = false;
                        $scope.flexgetLogs = data;
                        // Wait for template before scrolling to bottom
                        $timeout(function() {
                            $('#logs').animate({
                                'scrollTop': $('#logs')[0].scrollHeight
                            }, 100);
                        });
                    })
                    .error(function(data, status) {
                        $scope.logsLoading = false;
                        var data = data || "Request failed";
                        $scope.logsError = $sce.trustAsHtml(
                            '<strong>Unable to retrieve FlexGet logs:</strong> ' + data +
                            ' (' +
                            status + ')');
                    });
            };
            $scope.getLogs();
        }
    ]);
