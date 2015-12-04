$(document).ready(function(){
	$("#ad_switcher").bootstrapSwitch({
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
					alert(data);
				}
			});
		}
	});


});
