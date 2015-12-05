$(document).ready(function(){

	var getUrlParam = function(name) {
		var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
		var r = window.location.search.substr(1).match(reg);  //匹配目标参数
		if (r != null) return unescape(r[2]); return null; //返回参数值
	};


	var app_name = $('#app_name'),
	 	app_icon = $('#app_icon'),
	 	app_desc = $('#app_desc'),
	 	app_apk = $('#app_apk'),
		app_submitbt = $('#app_submitbt'),
		app_id = $('#app_id');

	var id = getUrlParam('id');
	if(id != null){
		$.ajax({
			url:'/action/app/id',
			type:'POST',
			dataType:'json',
			data:{
				id:id
			},
			success:function(data){
				if(data.status){
					app_name.val(data.data.name);
					app_desc.val(data.data.desc);
					app_id.val(id);
				}else{
					alert(data.msg);
				}
			}
		});
	}

	window.callback = function(data){
		if(data.status){
			alert('保存成功');
			window.location.href = '/right';
		}else{
			alert(data.msg);
			app_submit.removeAttr('disabled');
			app_submit.val('确认保存');
		}
	};


	app_submitbt.on('click', function(){
		var data = {
			id:app_id.val(),
			name:$.trim(app_name.val()),
			icon:app_icon.val(),
			desc:app_desc.val(),
			apk:app_apk.val()
		};

		if(data.name == ''){
			alert('请填写应用标题');
			app_name.focus();
			return;
		}

		if(data.icon == '' && data.id == ''){
			alert('请上传应用缩略图');
			return;
		}

		if(data.apk == '' && data.id == ''){
			alert('请上传应用APK文件');
			return;
		}

		$(this).attr('disabled', '');
		$('#myform').get(0).submit();

		$(this).val('正在保存中...');
	});

});
