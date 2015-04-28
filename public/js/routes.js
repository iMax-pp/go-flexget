var app = angular.module('goFlexGet');

app.config(function($routeProvider) {
  'use strict';

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
});
