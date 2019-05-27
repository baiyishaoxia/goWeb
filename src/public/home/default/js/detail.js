var data = [];
//通过Ajax获取留言
$(function () {
    $.ajax({
        'url': "/message/ajax",
        'type': 'get',
        'dataType': 'json',
        'data': {'key':2,'article_id':$("#article").attr("value")},
        'success': function (d) {
            if(d.status == 200){
                console.log(d.data);
                data = d.data;
            }
        }
    });
});

layui.use(['jquery', 'form', 'layedit', 'flow'], function(){
    var form = layui.form;
    var $ = layui.jquery;
    var layedit = layui.layedit;
    var flow = layui.flow;

    //评论和留言的编辑器
    var editIndex = layedit.build('remarkEditor', {
        height: 150,
        tool: ['face', '|', 'left', 'center', 'right', '|', 'link'],
    });
    //评论和留言的编辑器的验证
    layui.form.verify({
        content: function(value) {
            value = $.trim(layedit.getText(editIndex));
            if(value == "") return "至少得有一个字吧";
            layedit.sync(editIndex);
        },
        userId: function(value) {
            if(value == "" || value == null) return "至少你得先登录吧！";
        },
        replyContent: function(value) {
            if($.trim(value) == "") {
                return "至少得有一个字吧!";
            }
        }
    });
    //监听留言提交
    form.on('submit(formLeaveMessage)', function(data) {
        var index = layer.load(1);
        //模拟留言回复
        var url = '/message/create';
        $.ajax({
            type: "POST",
            url: url,
            data: data.field,
            success: function(res) {
                var user = res["user"]; //用户数据
                var message = res["message"]; //留言信息
                if(res.status == 200) {
                    layer.close(index);
                    var content = data.field.content;
                    var html = '<li><div class="comment-parent"><img src="' + user["head_img"] + '" alt="' + user["name"] + '"/><div class="info"><span class="username">' + user["name"] + '</span>';
                    if(message["users_id"] == '1') {
                        html += " <span class=\"is_bloger\">博主</span>&nbsp;";
                    }
                    html += '</div><div class="content">' + data.field.content + '</div><p class="info info-footer"><span class="time">' + message["created_at"] + '</span>&nbsp;&nbsp;&nbsp;&nbsp;<span>中国</span>&nbsp;&nbsp;<a class="btn-reply"href="javascript:;" style="color: #009688;font-size:14px;" onclick="btnReplyClick(this)">回复</a></p></div><hr /><!--回复表单默认隐藏--><div class="replycontainer layui-hide">' +
                        '<form class="layui-form" action="">            ' +
                        '<input type="hidden" id="comment" name="parent_id" value="' + message["id"] + '" />       ' +
                        '<input type="hidden" id="user" lay-verify="userId" name="user_id" value="' +message["users_id"] + '" />                    ' +
                        '<input type="hidden" id="article" lay-verify="userId" name="article_id" value="' +$('#article').val() + '" />                    ' +
                        '<div class="layui-form-item"><textarea name="content" lay-verify="replyContent"placeholder="回复@' + user["name"] + '"class="layui-textarea"style="min-height:80px;"></textarea>' +
                        '</div><div class="layui-form-item"><button class="layui-btn layui-btn-mini"lay-submit="formReply"lay-filter="formReply">提交</button></div>' +
                        '</form></div></li>';
                    $('.blog-comment').prepend(html);
                    $('#remarkEditor').val('');
                    editIndex = layui.layedit.build('remarkEditor', {
                        height: 150,
                        tool: ['face', '|', 'left', 'center', 'right', '|', 'link'],
                    });
                    layer.msg("评论成功", {
                        icon: 1
                    });
                }else{
                    layer.msg(res.data);
                }
            },
            error: function(data) {
                layer.msg("网络错误！");
            }
        });
        return false;
    });
    //评论显示
    flow.load({
        elem: '#commentList' //流加载容器
        ,done: function(page, next){ //执行下一页的回调
            setTimeout(function(){
                var lis = [];
                for(var i = 0; i < data.length; i++) {
                    var str = "";
                    for(var r = 0; r < data[i]["child"].length; r++) {
                        str += "<div class=\"comment-child\">\n" +
                            "      <img src=\"" + data[i]["child"][r]["head_img"] + "\" alt=\"" + data[i]["child"][r]["name"] + "\" />\n" +
                            "      <div class=\"info\">\n" +
                            "          <span class=\"username\">	" + data[i]["child"][r]["name"] + " : </span>";
                        if(data[i]["child"][r]["users_id"] == '1') {
                            str += "<span class=\"is_bloger\">博主</span>&nbsp;";
                        }
                        str += "回复 <span class=\"username\">@" + data[i]["child"][r]["replay_name"] + "：</span>";
                        if(data[i]["users_id"] == '1') {
                            str += "<span class=\"is_bloger\">博主</span>&nbsp;";
                        }
                        str += "<span>" + data[i]["child"][r]["content"] + "</span>\n" +
                            "      </div>\n" +
                            "      <p class=\"info\"><span class=\"time\">" + data[i]["child"][r]["created_at"] + "</span>&nbsp;&nbsp;&nbsp;&nbsp;<span>中国</span></p>\n" +
                            "  </div>\n";
                    }
                    lis.push("<li>\n" +
                        "               <div class=\"comment-parent\">\n" +
                        "                   <img src=\"" + data[i]["head_img"] + "\" alt=\"" + data[i]["name"] + "\" />\n" +
                        "                   <div class=\"info\">\n" +
                        "                       <span class=\"username\">" + data[i]["name"] + "</span>\n");
                    if(data[i]["users_id"] == '1') {
                        lis.push("<span class=\"is_bloger\">博主</span>&nbsp;");
                    }
                    lis.push("       </div>\n" +
                        "                   <div class=\"content\">\n" +
                        "                       " + data[i]['content'] + "\n" +
                        "                   </div>\n" +
                        "                   <p class=\"info info-footer\"><span class=\"time\">" + data[i]['created_at'] + "</span>&nbsp;&nbsp;&nbsp;&nbsp;<span>中国</span>&nbsp;&nbsp;<a class=\"btn-reply\" style=\"color: #009688;font-size:14px;\" href=\"javascript:;\" onclick=\"btnReplyClick(this)\">回复</a></p>\n" +
                        "               </div>\n" +
                        "               <hr />\n" + str +
                        "               <!-- 回复表单默认隐藏 -->\n" +
                        "               <div class=\"replycontainer layui-hide\">\n" +
                        "                   <form class=\"layui-form\" action=\"/reply/list/\">\n" +
                        "                   <input type=\"hidden\" id=\"comment\" name=\"parent_id\" value=\"" + data[i]['id'] + "\" />\n" +
                        "                   <input type=\"hidden\" id=\"user\" lay-verify=\"userId\" name=\"user_id\" value=\"" + $('#user').val() + "\" />\n" +
                        "                   <input type=\"hidden\" id=\"article\" name=\"article_id\" value=\""+ $('#article').val() +"\">"+
                        "                       <div class=\"layui-form-item\">\n" +
                        "                           <textarea name=\"content\" lay-verify=\"replyContent\" placeholder=\"回复  @" + data[i]['name'] + ":\" class=\"layui-textarea\" style=\"min-height:80px;\"></textarea>\n" +
                        "                       </div>\n" +
                        "                       <div class=\"layui-form-item\">\n" +
                        "                           <button class=\"layui-btn layui-btn-mini\" lay-submit=\"formReply\" lay-filter=\"formReply\">提交</button>\n" +
                        "                       </div>\n" +
                        "                   </form>\n" +
                        "               </div>\n" +
                        "           </li>");
                }

                    //执行下一页渲染，第二参数为：满足“加载更多”的条件，即后面仍有分页
                    //pages为Ajax返回的总页数，只有当前页小于总页数的情况下，才会继续出现加载更多
                    next(lis.join(''), page < 1);
            }, 500);
        }
    });

    //监听留言回复提交
    form.on('submit(formReply)', function(data) {
        var index = layer.load(1);
        //模拟留言回复
        var url = '/message/create';
        $.ajax({
            type: "POST",
            url: url,
            data: data.field,
            success: function(res) {
                var user = res["user"]; //用户数据
                var message = res["message"]; //留言信息
                if(res.status == 200) {
                    layer.close(index);
                    var html = '<div class="comment-child"><img src="' + user["head_img"] + '" alt="' + user["name"] + '"/><div class="info"><span class="username">' + user["name"] + ' : </span>';
                    if(message["users_id"] == '1') {
                        html += " <span class=\"is_bloger\">博主</span>&nbsp;";
                    }
                    html += "回复 <span class=\"username\">@" + message["replay_name"] + " </span>";
                    if(message["users_id"] == '1') {
                        html += " <span class=\"is_bloger\">博主</span>&nbsp;";
                    }
                    html += '：<span>' + data.field.content + '</span></div><p class="info"><span class="time">' + message["created_at"] + '</span>&nbsp;&nbsp;&nbsp;&nbsp;<span>中国</span></p></div>';
                    $(data.form).find('textarea').val('');
                    $(data.form).parent('.replycontainer').before(html).siblings('.comment-parent').children('p').children('a').click();
                    layer.msg("回复成功", {
                        icon: 1
                    });
                } else {
                    layer.msg(res.data);
                }
            },
            error: function(data) {
                layer.msg("网络错误！");
            }
        });
        return false;
    });
});






$(document).ready(function() {
    $(".fa-file-text").parent().parent().addClass("layui-this");
});


function classifyList(id) {
	layer.msg('功能要自己写哦！');
}

function btnReplyClick(elem) {
    var $ = layui.jquery;
    $(elem).parent('p').parent('.comment-parent').siblings('.replycontainer').toggleClass('layui-hide');
    if($(elem).text() == '回复') {
        $(elem).text('收起')
    } else {
        $(elem).text('回复')
    }
}
//请求数据
$(function () {
    category_list();//分类数据
    article_right();//文章数据
});
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
							"       <span><i class=\"layui-badge-rim layui-bg-red \">"+(i+1)+"</i></span>\n" +
							"       &nbsp;&nbsp;<a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
							"    </li>";
                    }else if(i==1){
                        strVar1+="<li>\n" +
                            "       <span><i class=\"layui-badge-rim layui-bg-orange \">"+(i+1)+"</i></span>\n" +
                            "       &nbsp;&nbsp;<a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
                            "    </li>";
                    }else if(i==2){
                        strVar1+="<li>\n" +
                            "       <span><i class=\"layui-badge-rim layui-bg-green \">"+(i+1)+"</i></span>\n" +
                            "       &nbsp;&nbsp;<a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
                            "    </li>";
                    }else{
                        strVar1+="<li>\n" +
                            "       <span><i class=\"layui-badge-rim\">"+(i+1)+"</i></span>\n" +
                            "       &nbsp;&nbsp;<a href=\"/article/detail/"+_click[i].id+"\">"+_click[i].title+"</a>\n" +
                            "    </li>";
                    }
                }
                for(var i=0;i<_red.length;i++) {
                    if(i==0){
                        strVar2+="<li>\n" +
                            "       <span><i class=\"layui-badge-rim layui-bg-red \">"+(i+1)+"</i></span>\n" +
                            "       &nbsp;&nbsp;<a href=\"/article/detail/"+_red[i].id+"\">"+_red[i].title+"</a>\n" +
                            "    </li>";
                    }else if(i==1){
                        strVar2+="<li>\n" +
                            "       <span><i class=\"layui-badge-rim layui-bg-orange \">"+(i+1)+"</i></span>\n" +
                            "       &nbsp;&nbsp;<a href=\"/article/detail/"+_red[i].id+"\">"+_red[i].title+"</a>\n" +
                            "    </li>";
                    }else if(i==2){
                        strVar2+="<li>\n" +
                            "       <span><i class=\"layui-badge-rim layui-bg-green \">"+(i+1)+"</i></span>\n" +
                            "       &nbsp;&nbsp;<a href=\"/article/detail/"+_red[i].id+"\">"+_red[i].title+"</a>\n" +
                            "    </li>";
                    }else{
                        strVar2+="<li>\n" +
                            "       <span><i class=\"layui-badge-rim\">"+(i+1)+"</i></span>\n" +
                            "       &nbsp;&nbsp;<a href=\"/article/detail/"+_red[i].id+"\">"+_red[i].title+"</a>\n" +
                            "    </li>";
                    }
                }
                $('#clickList').empty().append(strVar1);
                $('#redList').empty().append(strVar2);
            }
        }
    });
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
                    strVar+="<a href=\"/article?keywords="+_list[i].id+"\">"+_list[i].title+"</a>"
                }
                strVar+=" <div class=\"clear\"></div>";
                $('#categoryList').empty().append(strVar);
            }
        }
    });
}