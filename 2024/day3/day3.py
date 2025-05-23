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

        self.part1sum = 0
        self.part2sum = 0

        # find mul(\d+,\d+), or do(), or don't()
        pattern = re.compile(r'(mul\(\d+,\d+\)|do\(\)|don\'t\(\))')
        self.matches = pattern.findall(all_text)
        self.match_idx = 0

    """
    Consumes the input text and returns the next command in the string or ""
    if there is no more commands.
    """
    def returnNextMulCommand(self):
        skipMul = False
        while self.match_idx < len(self.matches):
            cmd = self.matches[self.match_idx]
            self.match_idx += 1
            print(f"Current Command: '{cmd}'")

            # if we found a mul, return it
            if cmd.startswith("mul") and not skipMul:
                return cmd
            
            # if we found a do, we can start returning mul() again
            if cmd.startswith("do"):
                print(f"do'ing mul again")
                skipMul = False

            # if we found a don't, we should stop returning mul() commands
            if cmd.startswith("don't"):
                print(f"NO MORE mul")
                skipMul = True

        return ""

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
        self.part1sum += product

    def totalSum(self):
        return self.part1sum
    
    def calculateSum(self):
        cmd = self.returnNextMulCommand()
        while cmd != "":
            product = self.getProductFromCommand(cmd)
            self.updateSum(product)
            cmd = self.returnNextMulCommand()
        return self.totalSum()



def testDay3():
    input_test = ("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)" +
                  "+mul(32,64]then(mul(11,8)mul(8,5))")
    dr = Day3Reader(input_test)

    test_cmds = ["mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"]
    test_products = [8, 25, 88, 40]
    test_total = sum(test_products)
    for idx in range(len(test_cmds)):
        cmd = dr.returnNextMulCommand()
        assert cmd == test_cmds[idx], f"FAIL: cmd is '{cmd}', but expected '{test_cmds[idx]}'"
        product = dr.getProductFromCommand(cmd)
        assert product == test_products[idx], f"FAIL: product is '{product}', but expected '{test_products[idx]}'"
        dr.updateSum(product)
    
    assert test_total == dr.totalSum(), f"FAIL: total is '{dr.totalSum()}', but expected '{test_total}'"

def testDay3Part2():
    input_test = ("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64]" +
                  "(mul(11,8)undo()?mul(8,5))")
    dr = Day3Reader(input_test)

    test_cmds = ["mul(2,4)", "mul(8,5)"]
    test_products = [8, 40]
    test_total = sum(test_products)
    for idx in range(len(test_cmds)):
        cmd = dr.returnNextMulCommand()
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
testDay3Part2()
executeDay3()

# day 3 part 1 answer: 159892596
# day 3 part 2 answer: 92626942