<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    <link rel="stylesheet" href="./views/style.css">

    <title>Demokratinder</title>
  </head>
  <body>

  <script>
    function flushcookie() {
        var http = new XMLHttpRequest();
        http.open("POST", window.location.href, true);
        http.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        http.send();
        http.onload = function () {alert('Cookies flushed!')}
      }
  </script>

  <div class="container">
    <div class="jumbotron">
      <h1>Hello, {{name}}!</h1>
      <p class="lead">This is a democratic client for your tinder account. You can use it to search for the love of your live
      with your friends!</p>
      <h2>How To:</h2>
      <p>You obviously succesfully started the server already. The next things you have to do is:</p>
      <ol>
        <li>Let everyone who wants to vote scan the QR code</li>
        <li>Open the Tinder tab that opened in your browser</li>
        <li>Login if you aren't already </li>
        <li>Make sure that the tinder tab is always the active tab</li>
        <li>Tinder with your friends!</li>
      </ol>
    </div>
  </div>
  <div class="container qrcode">
      <img src="./views/qrcode.png">
      <p>This is a QR-Code to invite your friends to join and vote!</p>
  </div>
  <div class="d-flex justify-content-center">
    <button onclick="flushcookie()" name="flush" value="flush" class="btn" id="no_info">Flush Cookies</button>
    <a href="./stop" class="btn">Stop Server</a>
  </div>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
  </body>
</html>