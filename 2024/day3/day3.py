import re

INPUT = "2024/day3/input.txt"

"""
Class for storing the input text, and returning the next command on each request.

Note: this is probably all just easier with regexps
"""
class Day3Reader:
    """
    all_text: str - contains the full input text from a day3 file.
    """
    def __init__(self, all_text):
        self.all_text = all_text
        self.i = 0      # pointer into the string
        self.text_len = len(all_text)
        self.sum = 0

    """
    Returns the index of the next command start after the given index,
    or -1 if we've read off the end of the input.
    """
    def findCommandStartFromIndex(self, index):
        # skip to the next mul( start.
        cur_index = index
        while ((cur_index + 4) < self.text_len and 
               self.all_text[cur_index:cur_index+4] != "mul("):
            cur_index += 1

        # Return error if we've read off the end of the string.
        if (cur_index + 4) >= self.text_len:
            return -1
        
        return cur_index

    """
    Consumes the input text and returns the next command in the string or ""
    if there is no more commands.
    """
    def returnNextCommand(self):
        # start from last iteration point, and keep reading until we find
        # the start of a command. Continue, skipping characters that are 
        # invalid, unless we see the start of another command ("mul(")
        cmd = ""
        # skip to the next mul( start.
        potential_index = self.findCommandStartFromIndex(self.i)
        if potential_index < 0:
            return ""
        self.i = potential_index

        # read until closing paren, but don't loop forever
        # this should say command length
        cmd_len = 3
        while ((self.i + cmd_len) < self.text_len and
               self.all_text[self.i+cmd_len] != ")"):
            cmd_len += 1
            #print(f"Current Command: '{self.all_text[self.i:self.i+cmd_len+1]}'")
            # if we find a new command starting, then fast forward to that
            # and restart.
            if self.all_text[self.i+cmd_len+1-4:self.i+cmd_len+1] == "mul(":
                # found a new command!
                # hack it up, and recurse
                self.i = self.i+cmd_len + 1 - 4
                return self.returnNextCommand()
        
        #print(f"i: {self.i}, cmd_len: {cmd_len}")

        # build the command
        cmd = self.all_text[self.i:self.i+cmd_len+1]
        print(f"Current Command: '{cmd}'")

        # update the next position to start
        self.i = self.i + cmd_len + 1
        
        return cmd
    
    def getProductFromCommand(self, cmd):
        match = re.search(r'mul\((\d+),(\d+)\)', cmd)
        if match:
            num1 = int(match.group(1))
            num2 = int(match.group(2))
            return num1 * num2
        return 0

    # Yes, it's bad form to have the caller update the class' instance
    # variables.
    def updateSum(self, product):
        self.sum += product

    def totalSum(self):
        return self.sum
    
    def calculateSum(self):
        cmd = self.returnNextCommand()
        while cmd != "":
            product = self.getProductFromCommand(cmd)
            self.updateSum(product)
            cmd = self.returnNextCommand()
        return self.totalSum()



def testDay3():
    input_test = ("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)" +
                  "+mul(32,64]then(mul(11,8)mul(8,5))")
    dr = Day3Reader(input_test)

    test_cmds = ["mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"]
    test_products = [8, 25, 88, 40]
    test_total = sum(test_products)
    for idx in range(len(test_cmds)):
        cmd = dr.returnNextCommand()
        assert cmd == test_cmds[idx], f"FAIL: cmd is '{cmd}', but expected '{test_cmds[idx]}'"
        product = dr.getProductFromCommand(cmd)
        assert product == test_products[idx], f"FAIL: product is '{product}', but expected '{test_products[idx]}'"
        dr.updateSum(product)
    
    assert test_total == dr.totalSum(), f"FAIL: total is '{dr.totalSum()}', but expected '{test_total}'"

def executeDay3():
    f = open(INPUT, "r")
    all_text = f.read()
    dr = Day3Reader(all_text)
    sum = dr.calculateSum()
    print(f"Total sum from text is {sum}")

testDay3()
executeDay3()

# day 1 part 1 answer: 159892596