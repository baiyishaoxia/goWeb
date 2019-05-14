/*
$(function() {
    $(".ajax").click(function () {
        var url = $(this).attr("href");
        $.get(url, [], function (data) {
            if (data.status == 201) {
                layer.msg(data.info);
                return
            }
            if (data.status == 200) {
                layer.msg(data.info);
                setTimeout(function () {
                    window.location.href = data.url
                }, 2000);
                return
            }
        });
        return false;
    })
})
*/
function AlertLnfo(alert_content,wide,height) {
    layer.open({
        type: 1,
        area: [wide,height],
        fix: false, //不固定
        maxmin: true,
        shade: 0.4,
        title: '查看信息',
        content:$(""+alert_content+""),
    });
}