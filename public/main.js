const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);

function sendVote(val) {
    ws.send(JSON.stringify(val))
}

