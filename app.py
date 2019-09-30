from bottle import route, run, template, request, response, get
from enum import Enum
import socket

class VoteEnum(Enum):
    LIKE = 1
    NOPE = 0
    NEUTRAL = 2

clientCount = 0
voteDict = {}

cookieName = "DemokratinerID"

hostname = socket.gethostname()
IPAddr = socket.gethostbyname(hostname)


@get('/')
def index():
    if request.get_cookie(cookieName):
        clientID = request.get_cookie(cookieName)
        clientCount = clientCount + 1
        print(voteDict)
        if clientID in voteDict:
            return template('<b>Hello {{name}}, du neidischer Schwanz! Du Fanboi! Schön, dass du wieder hier bist</b>!', name=clientID)
    else:
        countTinderer = str(len(voteDict) + 1)
        clientID = "TinderGuru " + countTinderer
        response.set_cookie(cookieName, clientID)
        voteDict[clientID] = VoteEnum.NEUTRAL
        print(voteDict)
        return template('<b>Hello {{name}}, du neidischer Schwanz! Du Fanboi! Du bist jetzt dabei!</b>!', name=clientID)

run(host=IPAddr, port=8080, debug=True)
