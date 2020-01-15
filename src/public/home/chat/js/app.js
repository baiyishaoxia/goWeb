new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false,  // True if email and username have been filled in
        userContent: '' // 用户列表
    },

    created: function() {
        var self = this;
        // 实例化socket
        this.ws = new WebSocket('wss://' + window.location.host + '/v2/ws');
        // 监听socket连接
        this.ws.onopen = this.wsonopen(this.ws);
        // 监听socket消息
        this.ws.addEventListener('message', function(e) {
            //console.log("ws",e);
            //console.log(e);
            if(isJsonString(e.data)){
                var msg = JSON.parse(e.data);
            }else {
                var msg = e.data;
            }
            if(isJsonString(e.data) && !msg.hasOwnProperty("current_user")) { //当前活跃用户
                if(msg.username == self.username){
                    self.chatContent+='<div style="text-align: right">'
                        + emojione.toImage(msg.message)
                        + '<div class="chip">'
                            + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                            + msg.username
                        + '</div>'
                        +'</div>';
                }else{
                    self.chatContent += '<div>'
                        + '<div class="chip">'
                            + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                            + msg.username
                        + '</div>'
                        + emojione.toImage(msg.message) + '<br/>'
                        +'</div>';
                }
                // Parse emojis
                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
            }else if(isJsonString(e.data) && msg.hasOwnProperty("current_user")){  //在线用户列表
                $("#user-mess").empty();
                self.userContent ='<div class="col s12 m12"> '+msg.current_user
                   +'    <div class="card blue-grey darken-1">\n' +
                    '      <span class="card-content white-text">'+msg.user_list_str+'</span>\n' +
                    '    </div>'
                   +'</div>'
                +' <div class="row">\n' +
                    '        <div class="col s12 m12">\n' +
                    '          <div class="card blue-grey darken-1">\n' +
                    '            <div class="card-content white-text">\n' +
                    '              <span class="card-title">聊天室说明</span>\n' +
                    '              <p>本系统采用GO+WebSocket打造一个简洁、方便、实用、高性能、高并发即时通讯平台。我很方便，因为我只需要一个游客的身份就可以有效地使用。它在服务器和浏览器之间建立了全双工通信！</p>\n' +
                    '            </div>\n' +
                    '            <div class="card-action">\n' +
                    '              <a href="#">这是一个说明</a>\n' +
                    '              <a href="#">这是一个平台</a>\n' +
                    '            </div>\n' +
                    '          </div>\n' +
                    '        </div>\n' +
                    '      </div>'
            }else{
                //系统广播
                self.chatContent += '<div class="chip">'
                    + '<img src="/public/home/chat/images/admin.jpg">'
                    + msg+'(来自系统消息)'
                    + '</div>'
                    + '<br/>';
                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight;
            }
        });
        // 监听close消息
        this.ws.addEventListener('close', function(e) {
            alert('请重新登录,Lost connection!');
        });
        // 监听error消息
        this.ws.addEventListener('error', function(e) {
            alert('请重新登录,Connection ERROR!');
        });
    },
    destroyed: function() {
        //页面销毁时关闭长连接
        this.ws.close();
    },
    methods: {
        wsonopen(ws) {
            console.log("WebSocket连接成功");
        },
        send: function () {
            if (this.newMsg != '') {
                if(this.ws.readyState == 3){
                    alert('请重新登录,Lost connection!');
                    return
                }
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
        },
        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    }
});


var vm2 = new Vue({
    el: '#h5',
    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        joined: false,  // True if email and username have been filled in
        userContent: '' // 用户列表
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('wss://' + window.location.host + '/v2/ws');
        this.ws.addEventListener('message', function(e) {
            //console.log("ws",e);
            //console.log(e);
            if(isJsonString(e.data)){
                var msg = JSON.parse(e.data);
            }else {
                var msg = e.data;
            }
            if(isJsonString(e.data) && !msg.hasOwnProperty("current_user")) {
                self.chatContent += '<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
                    + msg.username
                    + '</div>'
                    + emojione.toImage('<font color="white">'+ msg.message + '</font>') + '<br/>'; // Parse emojis
                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
            }else if(isJsonString(e.data) && msg.hasOwnProperty("current_user")){  //在线用户列表
                var ss = msg.user_list_str.split("：");
                self.userContent ='<div class="card blue-grey darken-1">\n' +
                    '          <span class="card-title red-text"> 【活跃】'+msg.current_user+'</span>'+
                    '          <div class="text-center white-text">'+ss[1]+'</div>' +
                    '    </div>'
            }else{
                //系统广播
                self.chatContent += '<div class="chip">'
                    + '<img src="/public/home/chat/images/admin.jpg">'
                    + msg+'(来自系统消息)'
                    + '</div>'
                    + '<br/>';
                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight;
            }
        });
        this.ws.addEventListener('close', function(e) {
            alert('请重新登录,Lost connection!');
        });
        this.ws.addEventListener('error', function(e) {
            alert('请重新登录,Connection ERROR!');
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                if(this.ws.readyState == 3){
                    alert('请重新登录,Lost connection!');
                    return
                }
                this.ws.send(
                    JSON.stringify({
                            email: this.email,
                            username: this.username,
                            message: $('<p>').html(this.newMsg).text() // Strip out html
                        }
                    ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
        },
        gravatarURL: function(email) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(email);
        }
    }
});
//是否是json数据
function isJsonString(str) {
    try {
        if (typeof JSON.parse(str) == "object") {
            return true;
        }
    } catch(e) {
    }
    return false;
}