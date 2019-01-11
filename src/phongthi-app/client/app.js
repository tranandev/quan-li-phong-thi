// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_diemthi").hide();
	$("#success_create").hide();
	$("#error_diemthi").hide();
	$("#error_query").hide();
	
	$scope.timTatcaID = function(){

		appFactory.timTatcaID(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_phongthi = array;
		});
	}

	$scope.timID = function(){

		var id = $scope.phongthi_id;

		appFactory.timID(id, function(data){
			$scope.query_phongthi = data;

			if ($scope.query_phongthi == "Could not locate phongthi"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.taoID = function(){

		appFactory.taoID($scope.phongthi, function(data){
			$scope.create_phongthi = data;
			$("#success_create").show();
		});
	}

	$scope.suaDiemthi = function(){

		appFactory.suaDiemthi($scope.diemthi, function(data){
			$scope.change_diemthi = data;
			if ($scope.change_diemthi == "Error: ID not found"){
				$("#error_diemthi").show();
				$("#success_diemthi").hide();
			} else{
				$("#success_diemthi").show();
				$("#error_diemthi").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.timTatcaID = function(callback){

    	$http.get('/tim_tat_ca_ID/').success(function(output){
			callback(output)
		});
	}

	factory.timID = function(id, callback){
    	$http.get('/tim_ID/'+id).success(function(output){
			callback(output)
		});
	}

	factory.taoID = function(data, callback){

		//data.vipham = data.longitude + ", "+ data.latitude;

		var phongthi = data.id + "-" + data.vipham + "-" + data.hoten + "-" + data.diemthi + "-" + data.ghichu;

    	$http.get('/tao_ID/'+phongthi).success(function(output){
			callback(output)
		});
	}

	factory.suaDiemthi = function(data, callback){

		var diemthi = data.id + "-" + data.diem;

    	$http.get('/change_diemthi/'+diemthi).success(function(output){
			callback(output)
		});
	}

	return factory;
});


