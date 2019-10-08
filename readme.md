 democratic experience for Tinder**

## What is Demokratinder?
Demokratinder is the attempt to enhance your Tinder
experience by inluding your friends and peers in your
partner selection.
After all, who could know you better and wish you
better than your friends?

## Prerequisites
* An account on Tinder
* A computer with a webbrowser (no mobile support)
* Friends with smartphones

## Instructions
* Download the executable
* Run the executable on your computer
* A server will start in the background and host your
session
*Open a browser to 1921681782 :8080 on your phone
*Show the QR code to your friends
*Each of you should be greeted with a different
identifier
*Get started!

## Technical stuff
Demokratinder is built in Python. The web server is
running on the minimal [Bottle]
(https://github.com/bottlepy/bottle) framework. Calls
of the clients are made via HTTP and are evaluated by
the client which then makes a keystroke action on the
opened Tinder page via the keyboard]
(https://github.com/boppreh/keyboard) library.
