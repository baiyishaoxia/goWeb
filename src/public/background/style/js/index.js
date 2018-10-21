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