!function(){
	var app = angular.module("login", []);
	app.controller("LoginController", function($scope, $http){
		$scope.username = $scope.password = '';
		$scope.submit = function(){
			this.username_tip = this.password_tip = '';
			if(this.username == ''){
				this.username_tip = '用户名不能为空';
				return;
			}

			if(this.password == ''){
				this.password_tip = '密码不能为空';
				return;
			}

			$http({
				url : '/login/check',
				method:'POST',
				data : {
					username: this.username,
					password: this.password
				}
			}).success(function(data, header, config, status){
				if(data.status){
					window.location.href = '/main';
				}else{
					alert(data.message);
				}	
			});
		}
	});
}()
