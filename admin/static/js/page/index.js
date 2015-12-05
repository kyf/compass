$(document).ready(function(){

	$.ajax({
		url : '/action/getsetting',
		type:'POST',
		dataType:'json',
		success:function(data){
			var state = true;
			if(data.data.ad_show > 0){
				state = false;
			}

			$("#ad_switcher").bootstrapSwitch({
				state:state,
				onSwitchChange:function(ev, state){
					console.log(state);	

					$.ajax({
						url:'/action/adsetting',
						data:{
							state:state
						},
						type:'POST',
						dataType:'json',
						success:function(data){
						}
					});
				}
			});


		}
	});


});
