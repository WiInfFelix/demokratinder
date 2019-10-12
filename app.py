from bottle import *
from enum import Enum
import socket
import qrcode
import os
import logging
import webbrowser
from VoteManager import VoteManager


class VoteEnum(Enum):
    LIKE = 1
    NOPE = 0
    NEUTRAL = 2


def main():
    vtn = VoteManager()
    voteDict = {}

    logging.basicConfig(level=logging.DEBUG,
                        format='%(asctime)s - %(levelname)s - %(message)s')

    def browser_init():
        webbrowser.open_new('https://www.tinder.com')
        webbrowser.open_new_tab(f'http://{IPAddr}:8080')

    cookieName = "DemokratinerID"
    hostname = socket.gethostname()
    IPAddr = socket.gethostbyname(hostname)

    viewfolder = os.path.join(os.getcwd(), 'views\\')
    if (os.path.exists(os.getcwd() + r'\views\qrcode.png')):
        os.remove(os.getcwd() + r'\views\qrcode.png')
        logging.debug('Deleted old QR code')
    logging.debug('Created new QR code')
    code = qrcode.make("http://" + IPAddr + ":8080")
    code.save(viewfolder + "qrcode.png", 'PNG')
    browser_init()

    @get('/')
    def index():
        try:
            if request.get_cookie(cookieName):
                client_id = request.get_cookie(cookieName)
                logging.debug("Anmeldung GET mit vorhandenem Cookie")
                if client_id in voteDict:
                    logging.debug("Cookie in der Datenbank gefunden")
                else:
                    logging.debug("Kein Cookie in der Datenbank gefunden")
                    voteDict[client_id] = VoteEnum.NEUTRAL
                    vtn.increment_cookie_count()
                    logging.debug(f"Cookie erstellt!")
                return template('index', name=client_id)
            else:
                count_tinderer = str(len(voteDict) + 1)
                client_id = "TinderGuru " + count_tinderer
                response.set_cookie(cookieName, client_id)
                vtn.increment_cookie_count()
                voteDict[client_id] = VoteEnum.NEUTRAL
                logging.debug("Anmeldung ohne Cookie!")
                return template('index', name=client_id, css=send_static("style.css"), qrc=send_static("qrcode.png"))
        finally:
            pass

    @route('<filename:path>')
    def send_static(filename):
        return static_file(filename, root='')

    @get('/voting')
    def get_voting_page():
        if request.get_cookie(cookieName) in voteDict:
            clientID = request.get_cookie(cookieName)
            return template('voting', name=clientID)
        else:
            index()

    @post('/voting')
    def vote():
        client_id = request.get_cookie(cookieName)
        user_vote = request.forms.get('vote')
        logging.debug(f'{user_vote} has been posted!')
        vtn.decide(client_id, user_vote)
        return template('voting', name=client_id)

    run(host=IPAddr, port=8080, debug=True)


if __name__ == '__main__':
    main()
