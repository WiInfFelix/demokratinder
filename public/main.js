const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);
const overlay_votingbtn = document.querySelector(".button-overlay");
const yesButton = document.querySelector("#yes");
const nopeButton = document.querySelector("#nope");
var tm;

overlay_votingbtn.onclick = function() {restoreVotingButtons();}

function sendVote(val, btn) {
    ws.send(JSON.stringify(val));
    if(btn) {voteSend(btn)};
    
}

ws.onopen = function () {
    setInterval(ping, 10000);
}

ws.onmessage = function(event) {
    const msg = event.data;
    if (msg == 'pong') {
        pong();
        return;
    }
    if(msg == 'reset_vote') {
        restoreVotingButtons();
        return;
    }
    console.debug("Uncatched message: "+ msg);

};

ws.onerror = function(event) {
    alert(event.data);
};

function voteSend(btn) {
    btn.classList.add("activated");
    btn.firstChild.classList.add("activated");
    overlay_votingbtn.style.display = "flex";
}

function restoreVotingButtons() {
    overlay_votingbtn.style.display = "none";
    yesButton.classList.remove("activated");
    yesButton.firstChild.classList.remove("activated");
    nopeButton.classList.remove("activated");
    nopeButton.firstChild.classList.remove("activated");
}

function ping() {
    ws.send('ping');
    tm = setTimeout(function () {

       const infotext = document.querySelector(".lead");
       infotext.innerHTML = "Lost Connection";
       infotext.style.color = "red";

    }, 5000);
}

function pong() {
    clearTimeout(tm);
}
