from bottle import route, view, run, template, request, response, get
from enum import Enum
import socket


class VoteEnum(Enum):
    LIKE = 1
    NOPE = 0
    NEUTRAL = 2


voteDict = {}

cookieName = "DemokratinerID"

hostname = socket.gethostname()
IPAddr = socket.gethostbyname(hostname)


@get('/')
def index():
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


run(host=IPAddr, port=8080, debug=True, reloader=True)
