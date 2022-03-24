const LEFT = "left";
const RIGHT = "right";

const EVENT_MESSAGE = "Message"
const EVENT_CLOSE = "Close"
const EVENT_CONNECT = "Connect"

const userPhotos = [
    "assets/resource/wah_head.png"
]

var user_name = "";
var id = _uuid();

setName();

function setName(){
    user_name = prompt("輸入本次聊天名稱", "");
    alert(user_name)
}

function _uuid() {
    var d = Date.now();
    if (typeof performance !== 'undefined' && typeof performance.now === 'function'){
      d += performance.now(); //use high-precision timer if available
    }
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
      var r = (d + Math.random() * 16) % 16 | 0;
      d = Math.floor(d / 16);
        return (c === 'x' ? r : (r & 0x3 | 0x8)).toString(16);
    });
}

var url = "ws://" + window.location.host + "/ws?id=" + id + "&name="+ user_name;
var ws = new WebSocket(url);

var chatroom = document.getElementsByClassName("msger-chat")
var text = document.getElementById("msg");
var send = document.getElementById("send");

//送出訊息
function handleMessageEvent() {
    ws.send(JSON.stringify({
        "event": EVENT_MESSAGE,
        "name" : user_name,
        "id" : id,
        "content": text.value,
    }));
    text.value = "";
}
//按下發送鍵 --> 送出訊息
send.onclick = function (e) {
    handleMessageEvent()
}
//按下Enter --> 送出訊息
text.onkeydown = function (e) {
    if (e.keyCode === 13 && text.value !== "") {
        handleMessageEvent()
    }
};

//處理socket接收的訊息
function getEventMessage(msg) {
    var msg = `<div class="msg-left">${msg}</div>`
    return msg
}
function getMessage(name, img, side, text) {
    const d = new Date()
    //   Simple solution for small apps
    var msg = `
    <div class="msg ${side}-msg">
      <div class="msg-img" style="background-image: url(${img});"></div>

      <div class="msg-bubble">
        <div class="msg-info">
          <div class="msg-info-name">${name}</div>
          <div class="msg-info-time">${d.getFullYear()}/${d.getMonth()}/${d.getDay()} ${d.getHours()}:${d.getMinutes()}</div>
        </div>

        <div class="msg-text">${text}</div>
      </div>
    </div>
  `
    return msg;
}

ws.onmessage = function (e) {
    var m = JSON.parse(e.data)
    var msg = ""
    switch (m.event) {
        case EVENT_MESSAGE:
            if (m.id == id) {
                msg = getMessage(m.name, userPhotos[0], RIGHT, m.content);
            } else {
                msg = getMessage(m.name, userPhotos[0], LEFT, m.content);
            }
            break;
        case EVENT_CONNECT:
            if (m.id != id) {
                msg = getEventMessage(m.content)
            } else {
                msg = getEventMessage("您已加入聊天室")
            }
            break;
        case EVENT_CLOSE:
            if (m.id != id) {
                msg = getEventMessage(m.content)
            } else {
                msg = getEventMessage("您已離開聊天室")
            }
            break;
    }
    insertMsg(msg, chatroom[0]);
};

ws.onclose = function (e) {
    console.log(e)
}

function insertMsg(msg, domObj) {
    domObj.insertAdjacentHTML("beforeend", msg);
    domObj.scrollTop += 500;
}

