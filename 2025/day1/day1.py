#!/usr/local/bin/python3

DEBUG = 1

"""
Given a dial and a rotation, apply the rotation and return the new dial.

Rotation is a string starting with 'L' or 'R' followed by a number.

If the rotation is 'L', subtract the number from the dial.
If the rotation is 'R', add the number to the dial.

Return the new dial with the count of times the dial passed 0.
"""
def rotate(dial, rot):
    newDial = dial
    zeroPasses = 0
    if rot.startswith("L"):
        newDial = dial - int(rot[1:])
        while newDial < 0:
            newDial = newDial + 100
            zeroPasses += 1

        # if we started at 0, don't count the first pass through negative
        # as a zero pass.
        if dial == 0:
            zeroPasses -= 1
    elif rot.startswith("R"):
        newDial = dial + int(rot[1:])
        while newDial > 99:
            newDial = newDial - 100
            zeroPasses += 1

        # if ends on 0, don't count the last rotation
        if newDial == 0:
            zeroPasses -= 1
    else:
        raise ValueError(f"Invalid rotation: {rot}")
    
    if DEBUG:
        print(f"DEBUG: rotate({dial}, {rot}) -> {newDial}, {zeroPasses}")
    return (newDial, zeroPasses)

"""
Given a dial and a rotation, apply the rotation and return True if rotate() 
ends at the same number, and False otherwise.
"""
def testRotate(dial, rot, expected, zeroPasses):
    (newDial, zp) = rotate(dial, rot)
    return (newDial == expected and zp == zeroPasses)

"""
Goes through a set of known test cases to ensure the expected result, reporting on
any failures through error logs.

Returns True if it passes, False otherwise.
"""
def testAllRotations():
    # Array of tuples of (dial, rotation, expected, zeroPasses)
    testCases = [
        (50, "L10", 40, 0),
        (50, "R10", 60, 0),
        (50, "L50", 0, 0),
        (50, "R50", 0, 0),
        (50, "L68", 82, 1),
        (82, "L30", 52, 0),
        (52, "R48", 0, 0),
        (0, "L5", 95, 0),
        (95, "R60", 55, 1),
        (55, "L55", 0, 0),
        (0, "L1", 99, 0),
        (99, "L99", 0, 0),
        (99, "R1", 0, 0),
        (99, "R101", 0, 1),
        (0, "R14", 14, 0),
        (14, "L82", 32, 1),
        (50, "L1000", 50, 10),
        (0, "L100", 0, 0)
    ]
    allPass = True
    for dial, rot, expected, zeroPasses in testCases:
        result = testRotate(dial, rot, expected, zeroPasses)
        if not result:
            if DEBUG:
                print(f"FAIL: {dial} {rot} -> {expected} != {result}, {zeroPasses}")
            allPass = False
        else:
            if DEBUG:
                print(f"pass: {dial} {rot} -> {expected} == {result}, {zeroPasses}")

    return allPass

def main():
    print("Starting!")
    
    if testAllRotations():
        print("PASS all test cases")
    else:
        print("FAIL test cases!")
        return -1

    dial = 50
    part1ZeroCount = 0
    part2ZeroCount = 0
    
    # open the file
    f = open('2025/day1/input.txt', 'r')

    # apply each rotation
    for l in f:
        l = l.strip()
        (dial, zp) = rotate(dial, l)
        if dial == 0:
            part1ZeroCount += 1
        part2ZeroCount += zp

    # Part 1 solution: 1102
    print(f"Part 1 Zero count: {part1ZeroCount}")

    part2Solution = part1ZeroCount + part2ZeroCount

    # Part 2 solution: 6635 (first try). Answer too high.
    # 5733 - Answer too low.
    # 6276 - wrong
    # 6175 - Fixing the math - Right answer!
    print(f"Part 2 Zero count: {part2ZeroCount}")
    print(f"Part 2 Solution: {part2Solution}")

if __name__ == '__main__':
    main()