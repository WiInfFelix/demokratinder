from app import VoteEnum

class VoteManager(self):

    def __init__(self):
        self.voteDict = {}
    

    def addVote(self, cookieId: String,vote:  VoteEnum):
        #check if cookieID is already put into the dict
        #if yes:
            #add cookie vote and increment counter of votes given
        #else:
            #discard vote with no answer

        #call evaluateVotes
        pass

    def evaluateVotes(self):
        #check if number of votes is number of clients
        #if no:
            #return
        #else:
            #evaluate votes against the number of total clients
            #call the EventManager for the right action
            #call flush
        pass

    def flush(self):
        dict.clear()