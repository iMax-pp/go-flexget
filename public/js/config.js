var app = angular.module('goFlexGet');

app.controller('ConfigCtrl', ['$scope', '$mdToast', '$http', '$sce', '$timeout',
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
        showErrorToast('Unable to retrieve FlexGet configuration: ' + data + ' (' + status + ')');
      });
  }
]);
