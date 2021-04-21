const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

function sendVote(val) {
    ws.send(JSON.stringify(val))
}

ws.onmessage = function(event) {
    const msg = JSON.parse(event.data);
}

