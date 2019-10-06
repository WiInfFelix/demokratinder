from bottle import *
from enum import Enum
import socket
import qrcode
import os


class VoteEnum(Enum):
    LIKE = 1
    NOPE = 0
    NEUTRAL = 2


voteDict = {}

cookieName = "DemokratinerID"

hostname = socket.gethostname()
IPAddr = socket.gethostbyname(hostname)
current_dir = os.path.dirname(os.path.realpath(__file__))
viewfolder = os.path.join(current_dir, 'views\\')
#os.remove(viewfolder + "qrcode.png")
code = qrcode.make("http://" + IPAddr + ":8080")
code.save(viewfolder + "qrcode.png", 'PNG')

@get('/')
def index():
    try:
        if request.get_cookie(cookieName):
            clientID = request.get_cookie(cookieName)
            print("Anmeldung GET mit vorhandenem Cookie ")
            if clientID in voteDict:
                print("Cookie in der Datenbank")
            else:
                print("Cookie nicht in der Datenbank -> neu hinzugefügt")
                voteDict[clientID] = VoteEnum.NEUTRAL
            return template('index', name=clientID)
        else:
            countTinderer = str(len(voteDict) + 1)
            clientID = "TinderGuru " + countTinderer
            response.set_cookie(cookieName, clientID)
            voteDict[clientID] = VoteEnum.NEUTRAL
            print("Anmeldung ohne Coockie. Neue Datenbank: " + voteDict)
            return template('index', name=clientID, css = send_static("style.css"), qrc = send_static("qrcode.png"))
    finally:
        pass

@route('<filename:path>')
def send_static(filename):
    return static_file(filename, root = '')


run(host=IPAddr, port=8080, debug=True, reloader=True)
