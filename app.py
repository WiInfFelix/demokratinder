from bottle import *
from enum import Enum
import socket
import qrcode
import os
import logging


class VoteEnum(Enum):
    LIKE = 1
    NOPE = 0
    NEUTRAL = 2


clientCount = 0
voteDict = {}

logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(levelname)s - %(message)s')

cookieName = "DemokratinerID"

hostname = socket.gethostname()
IPAddr = socket.gethostbyname(hostname)
viewfolder = os.path.join(os.getcwd(), 'views\\')
if(os.path.exists(os.getcwd() + r'\views\qrcode.png')):
    os.remove(os.getcwd() + r'\views\qrcode.png')
    logging.debug('Deleted old QR code')
logging.debug('Created QR code')
code = qrcode.make("http://" + IPAddr + ":8080")
code.save(viewfolder + "qrcode.png", 'PNG')


@get('/')
def index():
    try:
        if request.get_cookie(cookieName):
            clientID = request.get_cookie(cookieName)
            logging.debug("Anmeldung GET mit vorhandenem Cookie")
            if clientID in voteDict:
                logging.debug("Cookie in der Datenbank gefunden")
            else:
                logging.debug("Kein Cookie in der Datenbank gefunden")
                voteDict[clientID] = VoteEnum.NEUTRAL
                logging.debug(f"Cookie erstellt!")
            return template('index', name=clientID)
        else:
            countTinderer = str(len(voteDict) + 1)
            clientID = "TinderGuru " + countTinderer
            response.set_cookie(cookieName, clientID)
            voteDict[clientID] = VoteEnum.NEUTRAL
            logging.debug("Anmeldung ohne Cookie! Neue Datenbank: ")
            return template('index', name=clientID, css=send_static("style.css"), qrc=send_static("qrcode.png"))
    finally:
        pass


@route('<filename:path>')
def send_static(filename):
    return static_file(filename, root='')


run(host=IPAddr, port=8080, debug=True, reloader=True)
