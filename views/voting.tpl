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
  <div class="container">
    <div class="jumbotron">
      <h1>Hello, {{name}}!</h1>
      <p class="lead">What do you want to do?</p>
      <hr class="my-4">
      <form method="post">
        <button type="submit" name="vote" value="yes" class="btn">Yes</button>
        <button type="submit" name="vote" value="nope" class="btn">Nope</button>
        <button type="submit" name="vote" value="next" class="btn">Next Picture</button>
        <button type="submit" name="vote" value="info" class="btn">Show Info</button>
        <button type="submit" name="vote" value="no_info" class="btn">Hide Info</button>
      </form>
    </div>
  </div>
  <div class="container qrcode">
      <img src="./views/qrcode.png">
      <p>This is a QR-Code to invite your friends to join and vote!</p>
  </div>

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
  </body>
</html>