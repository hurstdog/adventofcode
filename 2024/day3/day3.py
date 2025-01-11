

INPUT = "2024/day3/input.txt"

"""
Class for storing the input text, and returning the next command on each request.
"""
class Day3Reader:
    """
    all_text: str - contains the full input text from a day3 file.
    """
    def __init__(self, all_text):
        self.all_text = all_text
        self.i = 0      # pointer into the string
        self.text_len = len(all_text)

    def returnNextCommand(self):
        # start from last iteration point, and keep reading until we find
        # the start of a command. Continue, skipping characters that are 
        # invalid, unless we see the start of another command ("mul(")
        cmd = ""
        # skip to the next mul( start.
        while ((self.i + 4) < self.text_len and 
               self.all_text[self.i:self.i+4] != "mul("):
            self.i += 1

        # read until closing paren, but don't loop forever
        # this should say command length
        cmd_len = 3
        while ((self.i + cmd_len) < self.text_len and
               self.all_text[self.i+cmd_len] != ")"):
            cmd_len += 1
            print(f"Current Command: '{self.all_text[self.i:self.i+cmd_len+1]}'")
        
        print(f"i: {self.i}, cmd_len: {cmd_len}")

        # build the command
        cmd = self.all_text[self.i:self.i+cmd_len+1]
        print(f"Current Command: '{cmd}'")

        # update the next position to start
        self.i = self.i + cmd_len + 1
        
        return cmd
    
    def totalSum(self):
        return 0


def testDay3():
    input_test = ("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)" +
                  "+mul(32,64]then(mul(11,8)mul(8,5))")
    dr = Day3Reader(input_test)

    tests = ["mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"]
    for t in tests:
        cmd = dr.returnNextCommand()
        assert cmd == t, f"FAIL: cmd is '{cmd}', but expected '{t}'"

def executeDay3():
    f = open(INPUT, "r")
    all_text = f.read()
    dr = Day3Reader(all_text)
    print(dr.totalSum())

testDay3()