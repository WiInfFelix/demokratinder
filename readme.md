# Demokratinder

Welcome to our project of shoddy code and bad ideas. This project might interest you,
if you are often times single, drunk, horny and with friends. Ideally your are all of the aforementioned.

## What is Demokratinder?
Demokratinder is the attempt to enhance your Tinder
experience by inluding your friends and peers in your
partner selection.
After all, who could know better to select a potential mate, 
then the people who wish you nothing but the best.

## Prerequisites
* An account on Tinder
* A computer with a webbrowser (no mobile support)
* Friends with smartphones

## Instructions
* Download the executable
* Run the executable on your computer
* A server will start in the background and host your
session
* Scan the QR code on the host page
* Show the QR code to your friends
* Each of you should be greeted with a different
identifier
* Get started!

## Technical stuff
Demokratinder is built in Python. The web server is
running on the minimal [Bottle]
(https://github.com/bottlepy/bottle) framework. Calls
of the clients are made via HTTP and are evaluated by
the client which then makes a keystroke action on the
opened Tinder page via the keyboard]
(https://github.com/boppreh/keyboard) library.
