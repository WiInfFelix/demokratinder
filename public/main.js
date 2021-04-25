const { console } = require("window-or-global");

const url = "ws://" + window.location.host + "/ws";
const ws = new WebSocket(url);
const overlay_votingbtn = document.querySelector(".button-overlay");
const yesButton = document.querySelector("#yes");
const nopeButton = document.querySelector("#nope");
const infotext = document.querySelector(".lead");
var intervalTimer;
var tm;

overlay_votingbtn.onclick = function() {restoreVotingButtons();}

function sendVote(val, btn) {
    ws.send(JSON.stringify(val));
    if(btn) {voteSend(btn)};
    
}

ws.onclose = function() {
    clearInterval(intervalTimer);
    infotext.innerHTML = "Connection closed!"
}

ws.onopen = function () {
    intervalTimer = setInterval(ping, 10000);
}

ws.onmessage = function(event) {
    const msg = JSON.parse(event.data);
    console.log(JSON.stringify(msg))
    if (msg.msg_type == 'pong') {
        pong();
        return;
    }
    if(msg.msg_type == 'reset_vote') {
        restoreVotingButtons();
        infotext.innerHTML = msg.decision_taken + " given! Please vote!";
        return;
    }
    if(msg.msg_type == 'votemade') {
        const text = msg.votes_given + " of " + msg.participants + " votes are in";
        infotext.innerHTML = text;
        return;
    }
    console.debug("Uncatched message: "+ JSON.stringify(msg));

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

       infotext.innerHTML = "Lost Connection";
       infotext.style.color = "red";

    }, 5000);
}

function pong() {
    clearTimeout(tm);
}
