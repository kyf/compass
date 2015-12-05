$(document).ready(function(){
	var datarender = $('#datarender');

	String.prototype.replaceTpl = function(ds){
		var reg = /{([^}]+)}/ig;
		return this.replace(reg, function(main, group){
			return ds[group]? ds[group]:'';
		});
	};


	var tpl = '<table class="tablelist">'+
    	'<thead>'+
    	'<tr>'+
        '<th><input id="allselector" type="checkbox" /></th>'+
        '<th>ID<i class="sort"><img src="images/px.gif" /></i></th>'+
        '<th>应用名称</th>'+
        '<th>缩略图</th>'+
        '<th>下载地址</th>'+
        '<th>操作</th>'+
        '</tr>'+
        '</thead>'+
        '<tbody>'+
        '{datalist}</tbody>'+
    '</table>';

	var subtpl = '<tr>'+
        '<td><input name="itemselector" type="checkbox" value="{id}" /></td>'+
        '<td>{id}</td>'+
        '<td>{name}</td>'+
        '<td><img src="{icon}" width="35" height="35" /></td>'+
        '<td><a href="/{apk}" target="_blank">{apk}</a></td>'+
        '<td><a href="/form?id={id}" class="tablelink">编辑</a>     <a href="javascript:del({id})" class="tablelink"> 删除</a></td>'+
        '</tr>';

	var loadfn = function(){
		$.ajax({
			url : '/action/app/list',
			dataType:'json',
			type:'POST',
			success:function(data){
				if(data.status){
					var tmp = new Array();
					$.each(data.data, function(index, it){
						tmp.push(subtpl.replaceTpl(it));
					});
					datarender.html(tpl.replaceTpl({'datalist':tmp.join('')}));
					$('.tablelist tbody tr:odd').addClass('odd');
				}else{
					alert(data.msg);
				}
			}
		});
	}

	var delfn = function(id, cb){
		if(confirm("确认删除?")){
			$.ajax({
				url:'/action/app/delete',
				data:{
					id:id
				},
				type:'POST',
				dataType:'json',
				success:function(data){
					if(data.status){
						cb();
					}else{
						alert(data.msg);
					}
				}
			});	
		}
	};

	window.del = function(id){
		delfn(id, function(){
			loadfn();
		});
	};

	loadfn();

	$('#AddBt').on('click', function(){
		window.location.href = '/form';
	});

	$('#RemoveBt').on('click', function(){
		
	});
});
