from collections import Counter
import keyboard
import logging


class VoteManager:

    def __init__(self):
        self.client_count = 0
        self.voteDict = {}

    def decide(self, client_id: str, vote: str):

        decision = {
            'yes': self.add_vote,
            'nope': self.add_vote,
            'next': self.next_pic,
            'info': self.open_info,
            'no_info': self.close_info,
        }

        decision[vote](client_id, vote)

    def add_vote(self, cookie_id: str, vote: str):
        self.voteDict[cookie_id] = vote
        self.evaluate_votes()
        

    def evaluate_votes(self):
        if len(self.voteDict) == self.client_count:
            vote_res = Counter(self.voteDict).most_common(1)
            #logging.debug(vote_res[0][1])
            if vote_res[0][1] == 'yes':
                logging.info('Swiped right')
                self.like()
            else:
                logging.info('Swiped left')
                self.nope()
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
        logging.debug('Pressed RIGHT on Keyboard')
        keyboard.press_and_release('right')

    def nope(self):
        logging.debug('Pressed LEFT on Keyboard')
        keyboard.press_and_release('left')

    def next_pic(self, *args):
        logging.debug('Pressed SPACE on Keyboard')
        keyboard.press_and_release('space')

    def open_info(self, *args):
        logging.debug('Pressed UP on Keyboard')
        keyboard.press_and_release('up')

    def close_info(self, *args):
        logging.debug('Pressed DOWN on Keyboard')
        keyboard.press_and_release('down')

    def superlike(self, *args):
        logging.debug('Pressed ENTER on Keyboard')
        keyboard.press_and_release('enter')

    def flush(self):
        self.voteDict.clear()

    def increment_cookie_count(self):
        self.client_count += 1
        logging.info(f'Clientcount increased to {self.client_count}')

    def decrement_cookie_count(self):
        self.client_count -= 1