from bottle import route, view, run, template, request, response, get
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
code = qrcode.make(IPAddr)
code.save(viewfolder + "qrcode.png", 'PNG')

@get('/')
def index():
    try:
        if request.get_cookie(cookieName):
            clientID = request.get_cookie(cookieName)
            print(voteDict)
            if clientID in voteDict:
                return template('index',
                                name=clientID)
        else:
            countTinderer = str(len(voteDict) + 1)
            clientID = "TinderGuru " + countTinderer
            response.set_cookie(cookieName, clientID)
            voteDict[clientID] = VoteEnum.NEUTRAL
            print(voteDict)
            return template('index', name=clientID)
    finally:
        pass
        #os.remove(viewfolder + "qrcode.png")


run(host=IPAddr, port=8080, debug=True, reloader=True)
