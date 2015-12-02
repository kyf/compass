$(document).ready(function(){
	var checkcode = $('#checkcode');


	checkcode.bind('click', function(){
		var src = '/checkcode?';
		$(this).attr('src', src + Math.random());
	});


	var loginfn = function(){
	
		$.ajax({
			url:'/action/login',
			type:'POST',
			data:{
				'username':$('#username').val(),
			'password':$('#password').val(),
			'checkcode':$('#checkcode_text').val()
			},
			dataType:'json',
			success:function(data, status, response){

			}
		});
	};

	$('#submitbt').bind('click', loginfn);
});
