$(function() {
//region 列表页删除选中的ID
    $('.btndel').click(function () {
        if ($(".checkall input:checked").length < 1) {
            layer.msg('对不起，请选中您要操作的记录！');
            return false;
        }
        var msg = "删除记录后不可恢复，您确定吗？";
        if (arguments.length == 2) {
            msg = objmsg;
        }
        var obj = this;
        parent.layer.confirm(msg, {
            btn: ['確定', '取消'] //按钮
        }, function () {
            var url =$('.btndel').attr("href");
            $.post(url,$("#form1").serialize(),function (data) {
                if(data.status==201){
                    layer.msg(data.info);
                    return
                }
                if(data.status==200){
                    layer.msg(data.info);
                    setTimeout(function () {
                        window.location.href=data.url
                    },2000);
                    return
                }
                if (data.url!=undefined) {
                    setTimeout(function () {
                        window.location.href = data.url
                    }, 2000);
                }
            });
            parent.layer.closeAll();
        });
        return false;
    });
//endregion
//region 列表页保存排序
$('.btnsave').each(function () {
    var href = $(this).attr('href');
    $(this).attr('href','javascript:void(0)');
    $(this).attr('url',href);
});
$('.btnsave').click(function () {
    var action=$(this).attr('url');
    $(this).parents('form').eq(0).attr('action',action);
    $(this).parents('form').eq(0).attr('method','POST');
    $(this).parents('form').eq(0).submit();
});
//endregion

//region ajax2提示 [post提交  设置当前a标签属性href提交  返回json数据：status，info，url] ,用于保存等操作，无需关闭layout框
    $(".ajax2").click(function () {
        var url =$(this).attr("href");
        $.post(url,$("#form1").serialize(),function (data) {
            if(data.status==201){
                layer.msg(data.info);
                return
            }
            if(data.status==200){
                layer.msg(data.info);
                setTimeout(function () {
                    window.location.href=data.url
                },2000);
                return
            }
            if (data.url!=undefined) {
                setTimeout(function () {
                    window.location.href = data.url
                }, 2000);
            }
        });
        return false;
    });
    //endregion

    //region ajax3提示 [get提交  设置当前a标签属性href提交  返回json数据：status，info，url]  无需关闭layout框
    $(".ajax3").click(function () {
        var url =$(this).attr("href");
        $.get(url,[],function (data) {
            if(data.status==201){
                layer.msg(data.info);
                return
            }
            if(data.status==200){
                layer.msg(data.info,{icon:1,time:1000},function () {
                        window.location.href=data.url
                });
                return;
            }
        });
        return false;
    })
    //endregion
});