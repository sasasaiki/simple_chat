$(function () {
  init()
})

var socket:WebSocket = null

function init() {
  let $msgBox = $("#chatbox textarea")
  let $messages = $("#messages")
  let $chatBox = $("#chatbox")


  $chatBox.submit(function() :boolean{doChat($msgBox,$messages);return false;})
  setWebSocket($messages)
}

function doChat($msgBox :JQuery,$messages:JQuery){
  if(!$msgBox.val())
  if(!socket) {
    alert("エラー:WebSocket接続が行われていません")

  }

  socket.send($msgBox.val())
  $msgBox.val("")
  console.log("sendしました")
}

function setWebSocket($messages:JQuery) {
  if (!window["WebSocket"]) {
    alert("エラー:対応していないブラウザです")
  } else {
    socket = new WebSocket("ws://localhost:8080/room")
    socket.onclose = function() {
      console.log("接続が終了しました")
    }
    socket.onmessage = function (e) {
      $messages.append($("<li>").text(e.data))
      console.log("onmessages");
    }
  }
}
