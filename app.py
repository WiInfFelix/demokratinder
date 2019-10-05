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
os.remove(viewfolder + "qrcode.png")
code = qrcode.make("https://" + IPAddr + ":8080")
code.save(viewfolder + "qrcode.png", 'PNG')

@get('/')
def index():
    try:
        if request.get_cookie(cookieName):
            clientID = request.get_cookie(cookieName)
            print("voteDict")
            if clientID in voteDict:
                print("lecker")
                return template('index',
                                name=clientID)
        else:
            countTinderer = str(len(voteDict) + 1)
            clientID = "TinderGuru " + countTinderer
            response.set_cookie(cookieName, clientID)
            voteDict[clientID] = VoteEnum.NEUTRAL
            print(voteDict)
            return template('index', name=clientID, css = send_static("style.css"), qrc = send_static("qrcode.png"))
    finally:
        pass

@route('/filename:path>')
def send_static(filename):
    print("lecker2")
    return static_file(filename, root='views/')       


run(host=IPAddr, port=8080, debug=True, reloader=True)
