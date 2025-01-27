
# find all versions of XMAS in the text.

import re

INPUT = "2024/day4/input.txt"
LINELEN = 140       # number of characters per line
DEBUG = True

_EXAMPLE_TEXT = """
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
"""

_EX_LTR = 3
_EX_RTL = 2
_EX_DOWN = 1
_EX_UP = 2
_EX_DIAG_FWD = 6  # wrong
_EX_DIAG_BWD = 0
_EX_TOTAL = 18

def getFileInput():
    f = open(INPUT, "r")
    l = f.read()
    return l.join("").replace("\n", "")

def getExampleInput():
    # if we're using example input, reset the global line length to 10
    global LINELEN
    LINELEN = 10

    x = _EXAMPLE_TEXT
    x.join("")
    x = x.replace("\n", "")
    return x

def countRTL(text):
    count = 0
    for i in range(len(text) - 3):
        if text[i:i+4] == "SAMX":
            count += 1

    if DEBUG:
        print(f"countRTL: {count}")
        if count != _EX_RTL:
            print(f"FAIL: RTL is {count}, but expected {_EX_RTL}")

    return count

def countLTR(text):
    count = 0
    for i in range(len(text) - 3):
        if text[i:i+4] == "XMAS":
            count += 1

    if DEBUG:
        print(f"countLTR: {count}")
        if count != _EX_LTR:
            print(f"FAIL: LTR is {count}, but expected {_EX_LTR}")

    return count

def countDown(text):
    # down will have each character exactly LINELEN apart, forwards
    count = 0
    for i in range(len(text) - 3):
        if i + (LINELEN * 3) < len(text):
            x = text[i]
            m = text[i + LINELEN]
            a = text[i + (LINELEN * 2)]
            s = text[i + (LINELEN * 3)]
            if x == "X" and m == "M" and a == "A" and s == "S":
                count += 1

    if DEBUG:
        print(f"countDown: {count}")
        if count != _EX_DOWN:
            print(f"FAIL: DOWN is {count}, but expected {_EX_DOWN}")

    return count

def countUp(text):
    # up will have each character exactly LINELEN apart, backwards
    count = 0
    for i in range(len(text) - 3):
        if i + (LINELEN * 3) < len(text):
            s = text[i]
            a = text[i + LINELEN]
            m = text[i + (LINELEN * 2)]
            x = text[i + (LINELEN * 3)]
            if x == "X" and m == "M" and a == "A" and s == "S":
                count += 1

    if DEBUG:
        print(f"countUp: {count}")
        if count != _EX_UP:
            print(f"FAIL: UP is {count}, but expected {_EX_UP}")

    return count

def countAll(text):
    count = 0
    count += countRTL(text)
    count += countLTR(text)
    count += countDown(text)
    count += countUp(text)
    return count

def main():
    #input = getInput()
    i = getExampleInput()
    c = countAll(i)
    print(f"count is {c}")

if __name__ == "__main__":
    main()