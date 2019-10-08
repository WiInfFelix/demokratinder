from app import app
from collections import Counter
import keyboard


class VoteManager:
    voteDict = {}

    def __init__(self):
        self.voteDict = {}

    def add_vote(self, cookieid: str, vote: app.VoteEnum):
        self.voteDict[cookieid] = vote
        self.evaluateVotes()
        '''
        check if cookieID is already put into the dict
        if yes:
            add cookie vote and increment counter of votes given
        else:
            discard vote with no answer

        call evaluateVotes
        '''

    def evaluate_votes(self):
        if len(self.voteDict) == app.clientCount:
            vote_res = Counter(self.voteDict).most_common(1)

            if vote_res[0][1] == app.VoteEnum.LIKE:
                # swipe to the right
                pass
            else:
                # swipe to the left
                pass

            self.flush()
        else:
            return
        '''
        check if number of votes is number of clients
        if no:
            return
        else:
            evaluate votes against the number of total clients
            call the EventManager for the right action
            call flush
        '''

    def like(self):
        keyboard.press_and_release('rigth')

    def nope(self):
        keyboard.press_and_release('left')

    def next_pic(self):
        keyboard.press_and_release('space')

    def open_info(self):
        keyboard.press_and_release('down')

    def close_info(self):
        keyboard.press_and_release('up')

    def superlike(self):
        keyboard.press_and_release('enter')

    def flush(self):
        self.voteDict.clear()