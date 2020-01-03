layui.use('form', function(){
    var form = layui.form;
    form.on('submit(searchForm)', function(data){
        var keywords=$("#keywords").val();
        // if(keywords==null || keywords==""){
        //     layer.msg('请输入要搜索的关键字');
        //     return false;
        // }
        search();
        return false; //阻止表单跳转。如果需要表单跳转，去掉这段即可。
    });
});



$(function(){
    $(".fa-file-text").parent().parent().addClass("layui-this");
    var keywords=$("#keywords").val();
    $("#keywords").keydown(function (event) {
        if (event.keyCode == 13) {
            var keyword=$("#keywords").val();
            if(keyword==null || keyword==""){
                layer.msg('请输入要搜索的关键字');
                return false;
            }
            search();
        }
    });
    $("#keywords").bind('input porpertychange',function(){
        search();
    });
});

function search() {
    var keywords=$("#keywords").val();
    article_list(0,1,10,"",keywords);   //获取数据
}

function classifyList(id) {
    article_list(id);
	layer.msg('加载完成！');
}
//公共请求加载文章数据 function
function article_list(id=0,page=1,limit=10,load="",keywords=""){
    $.ajax({
        'url':"/article/ajax",
        'type':'post',
        'dataType':'json',
        'data':{'id':id,'page':page,'limit':limit,'keywords':keywords},
        'success':function(data){
            if(data.status == 200){
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
                    //一键分享
                    var share1 ="               <div class=\"projectCon-share\">\n" +
                        "                            <span>\n" +
                        "                                <div class=\"bdsharebuttonbox bdshare-button-style0-32\">\n" +
                        "                                    <a target=\"_blank\" href=\"http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=http://go.afurun.xyz/article/detail/"+ _list[i].id +"&title="+ _list[i].title+"&summary="+ _list[i].intro+"...&pics=http://go.afurun.xyz/public/home/default/img/admin.jpg\"\n" +
                        "                                       class=\"bds_qzone\" title=\"分享到QQ空间\"></a>\n" +
                        "                                    <a target=\"_blank\" href=\"http://service.weibo.com/share/share.php?url=http://go.afurun.xyz/article/detail/"+ _list[i].id +"&sharesource=weibo&title="+  _list[i].title +"&pic=http://go.afurun.xyz/public/home/default/img/admin.jpg&searchPic=false\" class=\"bds_tsina\"  title=\"分享到新浪微博\"></a>\n" +
                        "                                    <a href=\"#\" class=\"bds_weixin\" data-id=\""+_list[i].id +"\" data-cmd=\"weixin\" title=\"分享到微信\"></a>\n" +
                        "                                    <a target=\"_blank\" href=\"http://connect.qq.com/widget/shareqq/index.html?url=http://go.afurun.xyz/article/detail/"+ _list[i].id +"&sharesource=qzone&title="+ _list[i].title+"&summary="+ _list[i].intro +"...&pics=http://go.afurun.xyz/public/home/default/img/admin.jpg\"\n" +
                        "                                       class=\"bds_sqq\" title=\"分享到QQ好友\"></a>\n" +
                        "                                </div>\n" +
                        "                            </span>\n" +
                        "                        </div>";
                    var share2 = "\t\t<span class=\"shareHover\">\n" +
                        "\t\t\t<div class=\"bdsharebuttonbox shareBox\" >\n" +
                        "\t\t\t\t<a target=\"_blank\" href=\"http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=url=http://go.afurun.xyz/article/detail/"+  _list[i].id +"&title="+ _list[i].title+"&summary="+ _list[i].intro +"&pics=http://go.afurun.xyz/public/home/default/img/admin.jpg\"\n" +
                        "\t\t\t\tclass=\"bds_qzone\" title=\"分享到QQ空间\">QQ空间</a>\n" +
                        "\t\t\t\t<a target=\"_blank\" href=\"http://connect.qq.com/widget/shareqq/index.html?url=http://go.afurun.xyz/article/detail/"+_list[i].id +"&sharesource=qzone&title="+_list[i].title +"&summary="+ _list[i].intro+"&pics=http://go.afurun.xyz/public/home/default/img/admin.jpg\"\n" +
                        "\t\t\t\tclass=\"bds_sqq\" title=\"分享到QQ好友\">QQ好友</a>\n" +
                        "\t\t\t\t<a target=\"_blank\" href=\"http://v.t.sina.com.cn/share/share.php?url=http://go.afurun.xyz/article/detail/"+_list[i].id  +"&amp;title="+ _list[i].title +"&amp;pic=http://go.afurun.xyz/public/home/default/img/admin.jpg\" class=\"bds_tsina\" title=\"分享到新浪微博\">新浪微博</a>\n" +
                        "\t\t\t\t<a href=\"#\" class=\"bds_weixin\" data-id=\""+_list[i].id+"\" data-cmd=\"weixin\" title=\"分享到微信\">微信</a>\n" +
                        "\t\t\t</div>\n" +
                        "\t   </span>";

                    strVar+="             <a href=\"/article/detail/"+_list[i].id+"\">"+_list[i].title+"</a>\n" +
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
                        ""+share1+""+
                        "                        </div>\n" +
                        "                    </div>";
                }
                if(load!=""){
                    $(".layui-flow-more").remove();
                    strVar +="                <div class=\"layui-flow-more\">\n" +
                        "                        <a href=\"javascript:void(0);\" id=\"page_load\" page=\""+data.data.page+"\" total=\""+data.data.num +"\"><cite>加载更多</cite></a>\n" +
                        "                    </div>";
                    $("#articleList").append(strVar);
                    window._bd_share_main.init(); //ajax加载完成后再执行一下分享才能起效
                }else{
                    strVar+="                <div class=\"layui-flow-more\">\n" +
                        "                        <a href=\"javascript:void(0);\" id=\"page_load\" page=\""+data.data.page+"\" total=\""+data.data.num +"\"><cite>加载更多</cite></a>\n" +
                        "                    </div>";
                    $("#articleList").empty().append(strVar);
                }
                if(data.data.num / limit <= page){
                    $("#page_load").remove();
                    $("#articleList").append($('<div class="layui-flow-more">\n' +
                        '                        <a href="javascript:;"><cite>没有更多了</cite></a>\n' +
                        '                    </div>'));
                }
            }else{
                layer.msg('加载失败');
            }
        }
    })
}
//加载分类数据
function category_list() {
    $.ajax({
        'url': "/category/ajax",
        'type': 'post',
        'dataType': 'json',
        'data': {},
        'success': function (data) {
            if (data.status == 200) {
                var _list = data.data.category;
                var strVar = "";
                strVar+="<div class=\"article-category-title\">分类导航</div>";
                for(var i=0;i<_list.length;i++) {
                    strVar+="<a href=\"javascript:classifyList("+_list[i].id+");\">"+_list[i].title+"</a>"
                }
                strVar+=" <div class=\"clear\"></div>";
                $('#categoryList').empty().append(strVar);
            }
        }
    });
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
                    //点击排行 top(1,2,3)的样式
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
//自动请求文章数据
$(function () {
    if(GetQueryString("keywords")==""){
        article_list();
    }
    category_list();
    article_right();
});
//加载更多ajax请求
$(document).on('click','#page_load',function(){
    var page = $(this).attr('page'); //分页的页码
    page = parseInt(page)+1;
    $(this).attr('page',page); //下一页+1
    var total = $(this).attr('total'); //数据总数
    var limit = 10; //每页分页的数量
    var keyword=$("#keywords").val();
    article_list(0,page,limit,"page_load",keyword);
});
//自定义搜索
function GetQueryString(name) {
      var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
      var r = window.location.search.substr(1).match(reg);
      if(r!=null) {return  unescape(r[2]);}
      return null;
 }
 //搜索请求
$(function () {
    var category_id = GetQueryString("keywords");//接收关键词
    $("#articleList").empty(); //清空数据
    article_list(category_id);   //获取数据
});
//---------------------------页面分享---------------------------
setTimeout(function(){
    var ShareId = "";
    var url = $("#url").attr('url');
    //绑定所有分享按钮所在A标签的鼠标移入事件，从而获取动态ID
    $(".bdsharebuttonbox a").on("mouseover",function () {
        ShareId = $(this).attr("data-id");
    });
    function SetShareUrl(cmd, config) {
        if (ShareId) {
            config.bdUrl = url + '/' + ShareId;
        }
        return config;
    }

    //百度分享
    window._bd_share_config = {
        "common": {
            onBeforeClick: SetShareUrl,
            "bdSnsKey": {},
            "bdText": "",
            "bdMini": "2",
            "bdMiniList": false,
            "bdPic": "",
            "bdStyle": "0",
            "bdSize": "32"
        },
        "share": {}};
    with(document)0[(getElementsByTagName('head')[0]||body).appendChild(createElement('script')).src='/static/api/js/share.js?v=89860593.js?cdnversion='+~(-new Date()/36e5)];
    //with (document) 0[(getElementsByTagName('head')[0] || body).appendChild(createElement('script')).src = 'http://bdimg.share.baidu.com/static/api/js/share.js?v=89860593.js?cdnversion=' + ~(-new Date() / 36e5)];
    window._bd_share_main.init();
},500);
