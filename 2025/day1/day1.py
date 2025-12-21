#!/usr/local/bin/python3

DEBUG = 1

"""
Given a dial and a rotation, apply the rotation and return the new dial.

Rotation is a string starting with 'L' or 'R' followed by a number.

If the rotation is 'L', subtract the number from the dial.
If the rotation is 'R', add the number to the dial.

Return the new dial.
"""
def rotate(dial, rot):
    newDial = dial
    if rot.startswith("L"):
        newDial = dial - int(rot[1:])
        while newDial < 0:
            newDial = newDial + 100
    elif rot.startswith("R"):
        newDial = dial + int(rot[1:])
        while newDial > 99:
            newDial = newDial - 100
    else:
        raise ValueError(f"Invalid rotation: {rot}")
    
    if DEBUG:
        print(f"DEBUG: rotate({dial}, {rot}) -> {newDial}")
    return newDial

"""
Given a dial and a rotation, apply the rotation and return True if rotate() 
ends at the same number, and False otherwise.
"""
def testRotate(dial, rot, expected):
    newDial = rotate(dial, rot)
    return newDial == expected

"""
Goes through a set of known test cases to ensure the expected result, reporting on
any failures through error logs.
"""
def testAllRotations():
    testCases = [
        (50, "L10", 40),
        (50, "R10", 60),
        (50, "L50", 0),
        (50, "R50", 0),
        (50, "L68", 82),
        (82, "L30", 52),
        (52, "R48", 0),
        (0, "L5", 95),
        (95, "R60", 55),
        (55, "L55", 0),
        (0, "L1", 99),
        (99, "L99", 0),
        (0, "R14", 14),
        (14, "L82", 32),
    ]
    for dial, rot, expected in testCases:
        result = testRotate(dial, rot, expected)
        if not result:
            print(f"FAIL: {dial} {rot} -> {expected} != {result}")
        else:
            print(f"pass: {dial} {rot} -> {expected} == {result}")

def main():
    print("Starting!")
    
    #testAllRotations()

    dial = 50
    zeroCount = 0

    # open the file
    f = open('2025/day1/input.txt', 'r')

    # apply each rotation
    for l in f:
        l = l.strip()
        dial = rotate(dial, l)
        if dial == 0:
            zeroCount += 1

    # Part 1 solution: 1102
    print(f"Zero count: {zeroCount}")

if __name__ == '__main__':
    main()