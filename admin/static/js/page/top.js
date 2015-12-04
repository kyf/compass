$(document).ready(function(){
	$('#logoutbt').bind('click', function(){
		$.ajax({
			url:'/action/logout',
			type:'POST',
			success:function(){
				window.parent.location.href = '/login';
			}
		});
	});
});
