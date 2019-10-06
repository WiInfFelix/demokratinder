import keyboard


class EventManager():

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
