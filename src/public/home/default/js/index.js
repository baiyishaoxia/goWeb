layui.use(['jquery','carousel','flow','layer'], function () {
    var $ = layui.jquery;
    var flow = layui.flow;
    var layer = layui.layer;

    var width = width || window.innerWeight || document.documentElement.clientWidth || document.body.clientWidth;
    width = width>1200 ? 1170 : (width > 992 ? 962 : width);
    var carousel = layui.carousel;
    //建造实例
    carousel.render({
      elem: '#carousel'
      ,width: width+'px' //设置容器宽度
      ,height:'360px'
      ,indicator: 'inside'
      ,arrow: 'always' //始终显示箭头
      ,anim: 'default' //切换动画方式
      
    });

    $(".home-tips-container span").click(function(){
        layer.msg($(this).text(), {
            time: 20000, //20s后自动关闭
            btn: ['明白了']
        });
    });


    $(".recent-list .hotusers-list-item").mouseover(function () {
        var name = $(this).children(".remark-user-avator").attr("data-name");
        var str = "【"+name+"】的评论";
        layer.tips(str, this, {
            tips: [2,'#666666']
        });
    });


    $("#QQjl").mouseover(function(){
        layer.tips('QQ交流', this,{
            tips: 1
        });
    });
    $("#gwxx").mouseover(function(){
        layer.tips('给我写信', this,{
            tips: 1
        });
    });
    $("#xlwb").mouseover(function(){
        layer.tips('新浪微博', this,{
            tips: 1
        });
    });
    $("#htgl").mouseover(function(){
        layer.tips('后台管理', this,{
            tips: 1
        });
    });
    
    $(function () {
        $(".fa-home").parent().parent().addClass("layui-this");
        //播放公告
        playAnnouncement(5000);
    });
    
    function playAnnouncement(interval) {
        var index = 0;
        var $announcement = $('.home-tips-container>span');
        //自动轮换
        setInterval(function () {
            index++;    //下标更新
            if (index >= $announcement.length) {
                index = 0;
            }
            $announcement.eq(index).stop(true, true).fadeIn().siblings('span').fadeOut();  //下标对应的图片显示，同辈元素隐藏
        }, interval);
    }
});

function classifyList(id){
    article_list(id);
    layer.msg('加载完成！');
}


//加载右侧文章推荐数据
function article_right(){
    $.ajax({
        'url': "/article/right",
        'type': 'get',
        'dataType': 'json',
        'data': {'limit':10},
        'success': function (data) {
            if (data.status == 200) {
                var _red = data.data.red;
                var _click = data.data.click;
                var strVar1 = "";
                var strVar2 = "";
                for(var i=0;i<_click.length;i++) {
                    if(i==0){
                        strVar1+="<li>\n" +
                            "         <span><i class=\"layui-badge-rim layui-bg-red\" id=\"item"+i+"\">"+(i+1)+"</i></span> &nbsp;&nbsp;\n" +
                            "           <a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
                            "   </li>"
                    }else if(i==1){
                        strVar1+="<li>\n" +
                            "         <span><i class=\"layui-badge-rim layui-bg-orange\" id=\"item"+i+"\">"+(i+1)+"</i></span> &nbsp;&nbsp;\n" +
                            "           <a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
                            "   </li>"
                    }else if(i==2){
                        strVar1+="<li>\n" +
                            "         <span><i class=\"layui-badge-rim layui-bg-green\" id=\"item"+i+"\">"+(i+1)+"</i></span> &nbsp;&nbsp;\n" +
                            "           <a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
                            "   </li>";
                    }else{
                        strVar1+="<li>\n" +
                            "         <span><i class=\"layui-badge-rim\" id=\"item"+i+"\">"+(i+1)+"</i></span> &nbsp;&nbsp;\n" +
                            "           <a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
                            "   </li>";
                    }
                }
                strVar1+=" <div class=\"clear\"></div>";
                for(var i=0;i<_red.length;i++) {
                    strVar2+="<li>\n" +
                        "          <span class=\"article_is_yc\">原创</span> &nbsp;&nbsp;\n" +
                        "           <a href=\"/article/detail/"+_click[i].id+"\">"+_red[i].title+"</a>\n" +
                        "    </li>"
                }
                $('#clickList').empty().append(strVar1);
                $('#redList').empty().append(strVar2);
            }
        }
    });
}
//加载作者信息数据
function site_author(){
    $.ajax({
        'url': "/about/ajax",
        'type': 'post',
        'dataType': 'json',
        'data': {'index': "about_author"},
        'success': function (data) {
            if (data.status == 200) {
                var msg = data.data.info;
                $('.blogerinfo-nickname').html(msg[0].title);
                $('.blogerinfo-introduce').html(msg[0].intro);
                $('.site-author').html(msg[0].abstract);
            }
        }
    });

}
//加载文章数据
function article_list(id=0,page=1,limit=10,load=""){
    $.ajax({
        'url':"/article/ajax",
        'type':'post',
        'dataType':'json',
        'data':{'id':id,'page':page,'limit':limit},
        'success':function(data){
            if(data.status == 200){
                if(data.data.num / limit <= page){
                    $("#page_load").remove();
                    $("#articleList").append($('<div class="layui-flow-more">\n' +
                        '                        <a href="javascript:;"><cite>没有更多了</cite></a>\n' +
                        '                    </div>'));
                }
                var _list = data.data.news;
                var strVar = "";
                for(var i=0;i<_list.length;i++) {
                    strVar += "<div class=\"article shadow animated zoomIn\">\n" +
                        "                        <div class=\"article-left \">\n" +
                        "                            <img src=\""+_list[i].image+"\" alt=\""+_list[i].title+"\">\n" +
                        "                        </div>\n" +
                        "                        <div class=\"article-right\">\n" +
                        "                            <div class=\"article-title\">\n";
                        if(_list[i].is_top == true){
                          strVar+="<span class=\"article_is_top\">置顶</span>&nbsp;";
                        }
                        if(_list[i].source=="原创"){
                          strVar+=" <span class=\"article_is_yc\">原创</span>&nbsp;\n";
                        }
                        strVar+="    <a href=\"/article/detail/"+_list[i].id+"\">"+_list[i].title+"</a>\n" +
                        "                            </div>\n" +
                        "                            <div class=\"article-abstract\">\n" +
                        "                                "+_list[i].intro+"</div>\n" +
                        "                        </div>\n" +
                        "                        <div class=\"clear\"></div>\n" +
                        "                        <div class=\"article-footer\">\n" +
                        "                            <span><i class=\"fa fa-clock-o\"></i>&nbsp;&nbsp;"+_list[i].created_at+"</span>\n" +
                        "                            <span class=\"article-author\"><i class=\"fa fa-user\"></i>&nbsp;&nbsp;"+_list[i].author+"</span>\n" +
                        "                            <span><i class=\"fa fa-tag\"></i>&nbsp;&nbsp;<a href=\"javascript:void(0);\"> "+_list[i].keywords+"</a></span>\n" +
                        "                            <span><i class=\"fa fa-fa\"></i>&nbsp;&nbsp;<a href=\"javascript:classifyList("+_list[i].cate_id+");\"> "+_list[i].cate_title+"</a></span>\n" +
                        "                            <span class=\"article-viewinfo\"><i class=\"fa fa-eye\"></i>&nbsp;"+_list[i].click_num+"</span>\n" +
                        "                            <span class=\"article-viewinfo\"><i class=\"fa fa-commenting\"></i>&nbsp;"+_list[i].count_num+"</span>\n" +
                        "                        </div>\n" +
                        "                    </div>";
                }
                if(load!=""){
                    $("#articleList").append(strVar);
                }else{
                    strVar+="                <div class=\"layui-flow-more\">\n" +
                        "                        <a href=\"javascript:void(0);\" id=\"page_load\" page=\""+data.data.page+"\" total=\""+data.data.num +"\"><cite>加载更多</cite></a>\n" +
                        "                    </div>";
                    $("#articleList").empty().append(strVar);
                }
            }else{
                layer.msg('加载失败');
            }
        }
    })
}
//加载更多ajax请求
$(document).on('click','#page_load',function(){
    var page = $(this).attr('page'); //分页的页码
    page = parseInt(page)+1;
    $(this).attr('page',page); //下一页+1
    var sort = '';
    var total = $(this).attr('total'); //数据总数
    var limit = 10; //每页分页的数量
    var status  = '';
    var keywords  = '';
    article_list(0,page,limit,"page_load");
});
//加载最近评论
function user_message(){
    $.ajax({
        'url': "/message/ajax",
        'type': 'get',
        'dataType': 'json',
        'data': {'limit': 10,'type':"index"},
        'success': function (data) {
            if (data.status == 200) {
                var strVar = "";
                var _list = data.data;
                for(var i=0;i<_list.length;i++) {
                    strVar +="<li class=\"hotusers-list-item\">\n" +
                             "<div data-name=\""+ _list[i]["name"]+"\" class=\"remark-user-avator\">\n" +
                             "  <img src=\""+ _list[i]["head_img"]+"\" width=\"45\" height=\"45\">\n" +
                             "</div>\n" +
                             "<div class=\"remark-content\">"+_list[i]["content"]+"</div>\n" +
                             "<span class=\"hotusers-icons\"></span>\n" +
                             "</li>"
                }
                $(".recent-list").empty().append(strVar);
            }
        }
    });
}
$(function () {
    article_right();
    site_author();
    article_list();
    user_message();
});


