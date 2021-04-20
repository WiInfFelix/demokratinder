# Demokratinder

A democratic client to your Tinder experience.

## How does it the work?
### For the user
* Start the client by clickin the .exe file
* Two browser windows should open: the Tinder web client and the Demokratinder client
* Start by scanning the QR code with your phone
* Connect as many friends as you would like
* Change the active tab to tinder and log in if you have not till yet
* Start voting!
* When you are down, kill the application by clicking the button at the bottom

### The specifics
Demokratinder is written in Go and uses the [Gin Gonic](https://github.com/gin-gonic/gin) framework as web server backend. Websockets are realized using the [Melody](https://github.com/olahol/melody) websocket module.