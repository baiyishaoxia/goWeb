var pics = {
  "data": [
      {"src": "/public/home/default/images/pic/1.jpg",   "alt": "紫霞仙子和至尊宝",},
      {"src": "/public/home/default/images/pic/2.jpg",   "alt": "紫霞仙子"}
   ]};

var currpage = 1;
var total = 10;
function picture_list(page=1,limit=10,load="") {
    $.ajax({
        url: "/mixed/pic/ajax",
        type: 'post',
        dataType: 'json',
        data: {'limit': limit, 'page': page},
        async: false,
        success: function (data) {
            //console.log(data.data.item);
            currpage = data.data.page;
            total = data.data.num;
            if(load==""){
                pics = data.data.item;
            }else{
                //加载更多
                var _list = data.data.item["data"];
                var strVar = "";
                for(var i=0;i<_list.length;i++){
                    strVar+="<div class=\"mixed shadow animated zoomIn\">\n" +
                        "<div class=\"mixed-pic\">\n" +
                        '<a href="javascript:view('+_list[i].id+')" ><img src=\"'+_list[i].src +'\" alt=\"'+ _list[i].alt+'\"></a>\n' +
                        "</div>\n" +
                        "<div class=\"mixed-info\">"+_list[i].alt+"</div>\n" +
                        "<div class=\"mixed-footer\">\n" +_list[i].intro +
                        "</div>\n" +
                        "</div>"
                }
                $(".mixed-main").append(strVar);
                if(total / limit <= page){
                    $(".layui-flow-more").remove();
                    $(".mixed-main").append($('<div class="layui-flow-more">\n' +
                        '                        <a href="javascript:;"><cite>没有更多了</cite></a>\n' +
                        '                    </div>'));
                }else{
                    $(".layui-flow-more").remove();
                    $("#page_load").remove();
                    var end = '<a href="javascript:void(0);" id="page_load" page="'+currpage +'" total="'+total+'"><cite>加载更多</cite></a>\n';
                    $(".mixed-main").append(end);
                }
			}
        }
    });
}
//自动加载
picture_list(1,6);
layui.use(['jquery','flow'], function () {
	var $ = layui.jquery;
	// 流加载 图片
    var flow = layui.flow;
    var count = pics.data.length-1;
    flow.load({
    	elem: '.mixed-main', //流加载容器
    	isAuto: true,
    	end: '<div class="layui-flow-more">\n' +
        '                        <a href="javascript:void(0);" id="page_load" page="'+currpage +'" total="'+total+'"><cite>加载更多</cite></a>\n' +
        '                    </div>',
    	done: function(page,next) {
    		var lis = [];
    		for (var i=0; i<8; i++) {
    			if (count < -1) break;
    			if (count==-1) {
    				lis.push('<div class="mixed shadow animated zoomIn">'+
    	                 '<div class="mixed-pic">'+
    	                    '<a href="javascript:"><img src="/public/home/default/images/pic/0.jpg" alt="图片还在拍摄中" /></a>'+
	                    '</div>'+
	                    '<div class="mixed-info">相册还在整理中</div>'+
	                    '<div class="mixed-footer">'+
	                       '没有相册啦'+
	                    '</div>',
	                '</div>');
    			} else {
	    			lis.push('<div class="mixed shadow animated zoomIn">'+
	                    '<div class="mixed-pic">'+
	                        '<a href="javascript:view('+pics.data[count].id+')"><img src="'+pics.data[count].src+'" alt="'+pics.data[count].alt+'" /></a>'+
	                    '</div>'+
	                    '<div class="mixed-info">'+pics.data[count].alt+'</div>'+
	                    '<div class="mixed-footer">'+pics.data[count].intro +
	                    '</div>',
	                '</div>');
    			}
    			count--;
    		}
    		next(lis.join(''), page < pics.data.length/8);
    	}
    });
});
function view(start) {
	window.location.href = "/mixed/pic/detail/"+start;
}

$(function () {
	$(".fa-paper-plane-o").parent().parent().addClass("layui-this");
});

//加载更多ajax请求
$(document).on('click','#page_load',function(){
    var page = $(this).attr('page'); //分页的页码
    page = parseInt(page)+1;
    $(this).attr('page',page); //下一页+1
    var total = $(this).attr('total'); //数据总数
    var limit = 6; //每页分页的数量
    picture_list(page,limit,"page_load");
});