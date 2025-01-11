

INPUT = "2024/day3/input.txt"

"""
Class for storing the input text, and returning the next command on each request.
"""
class Day3Reader:
    def __init__(self, all_text):
        self.all_text = all_text

    def returnNextCommand(self):
        return "mul(3,2)"
    
    def totalSum(self):
        return 0


def testDay3():
    input_test = ("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)" +
                  "+mul(32,64]then(mul(11,8)mul(8,5))")
    dr = Day3Reader(input_test)

    tests = ["mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"]
    for t in tests:
        cmd = dr.returnNextCommand()
        assert cmd == t, f"Assertion failed: cmd is {cmd}, but expected {t}"

def executeDay3():
    f = open(INPUT, "r")
    all_text = f.read()
    dr = Day3Reader(all_text)
    print(dr.totalSum())

testDay3()