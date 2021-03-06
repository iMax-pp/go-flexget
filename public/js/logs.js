var app = angular.module('goFlexGet');

app.controller('LogsCtrl', ['$scope', '$mdToast', '$http', '$sce', '$timeout',
  function($scope, $mdToast, $http, $sce, $timeout) {
    'use strict';

    showErrorToast = function(error) {
      $mdToast.show(
        $mdToast.simple()
        .content(error)
        .action('X')
        .hideDelay(0)
      );
    };

    $scope.getLogs = function() {
      $scope.logsLoading = true;
      $http.get('/api/logs')
        .success(function(data) {
          $scope.logsLoading = false;
          $scope.flexgetLogs = data;
          // Wait for template before scrolling to bottom
          $timeout(function() {
            var logs = angular.element(document.getElementById('logs'));
            var logsBottom = angular.element(document.getElementById('logs-bottom'));
            logs.scrollToElementAnimated(logsBottom);
          });
        })
        .error(function(data, status) {
          $scope.logsLoading = false;
          data = data || 'Request failed';
          showErrorToast('Unable to retrieve FlexGet logs: ' + data + ' (' + status + ')');
        });
    };
    $scope.getLogs();
  }
]);
