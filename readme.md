# Demokratinder

A democratic client to your Tinder experience.


## What you will need
* A computer with internet connection
* Smartphones to use as controllers
* Friends
## For the user
* Head to the release page and download the latest release as a .zip
* Unzip the files to a location of your choosing
* Start the server by clickin the .exe file
* A terminal should open, signifying the server has started
* Two browser windows should open: the Tinder web client and the Demokratinder client
* Start by scanning the QR code with your phone
* Connect as many friends as you would like
* Change the active tab to tinder and log in if you have not till yet
* If any pop ups appear, you will have to close them manually and make sure the profile card is the focus of the browser
* Start voting!
* When you are done, kill the application by clicking the button at the bottom

## The technical details 
Demokratinders backend is written in Go and uses the [Gin Gonic](https://github.com/gin-gonic/gin) framework as web server backend. Websockets are realized using the [Gobwas/ws](https://github.com/gobwas/ws) websocket module. The frontend is plain HTML/CSS/Javascript with some Bootstrap flavouring.

The server opens a websocket with the client on connecting and awaits the users vote. When all users have voted a decision will be taken, that is executed as a keystroke mapped to the shortcuts Tinder provides for the web interface of its service.