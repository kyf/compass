$(document).ready(function(){
	var checkcode = $('#checkcode');


	checkcode.bind('click', function(){
		var src = '/checkcode?';
		$(this).attr('src', src + Math.random());
	});
});
