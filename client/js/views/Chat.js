import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Chat")
    }

    async init() {
        // let response = await fetch('http://localhost:8383/api/chat')
        window.onload = function () {

                var conn;
                var msg = document.getElementById("msg");
                var log = document.getElementById("log");

                function appendLog(item) {
                    var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
                    log.appendChild(item);
                    if (doScroll) {
                        log.scrollTop = log.scrollHeight - log.clientHeight;
                    }
                }

                document.getElementById("form").onsubmit = function () {
                    if (!conn) {
                        return false;
                    }
                    if (!msg.value) {
                        return false;
                    }
                    conn.send(msg.value);
                    msg.value = "";
                    return false;
                };


                conn = new WebSocket("ws://localhost:8383/api/chat");

                conn.onclose = function (evt) {
                    var item = document.createElement("div");
                    item.innerHTML = "<b>Connection closed.</b>";
                    appendLog(item);
                };
                conn.onmessage = function (evt) {
                    var messages = evt.data.split('\n')
                    console.log(messages)
                    for (var i = 0; i < messages.length; i++) {
                        var item = document.createElement("div");
                        item.innerText = messages[i];
                        appendLog(item);
                    }
                };

            }
        
    }

    async getHtml() {
        return `
        
            <div>Chat</div>
            <div id="log"></div>
<form id="form">
    <input type="submit" value="Send" />
    <input type="text" id="msg" size="64" autofocus />
</form>
        `
    }
}