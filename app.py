from bottle import route, run, template
from enum import Enum

class VoteEnum(Enum):
    LIKE = 1
    NOPE = 0


votedict = {}

@route('/hello/<name>')
def index(name):
    return template('<b>Hello {{name}}, du neidischer Schwanz! Du Fanboi!</b>!', name=name)


run(host='192.168.178.183', port=8080, debug=True)
