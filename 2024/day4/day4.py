
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
_EX_DIAG_SE = 1
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

def matchAtIndex(text, x, m, a, s):
    return (text[x] == "X" and
            text[m] == "M" and
            text[a] == "A" and
            text[s] == "S")

def countByOffset(text, x_i, m_i, a_i, s_i):
    count = 0
    for i in range(len(text)):

        # Skip anything that would cause out of bounds
        if (i + x_i < 0 or
            i + s_i < 0):
            continue
        if (i + x_i >= len(text) or
            i + s_i >= len(text)):
            continue

        if matchAtIndex(text,
                        i + x_i,
                        i + m_i,
                        i + a_i,
                        i + s_i):
            count += 1
    return count

def countRTL(text):
    count = countByOffset(text, 0, -1, -2, -3)

    if DEBUG:
        print(f"countRTL: {count}")
        if count != _EX_RTL:
            print(f"FAIL: RTL is {count}, but expected {_EX_RTL}")

    return count

def countLTR(text):
    count = countByOffset(text, 0, 1, 2, 3)

    if DEBUG:
        print(f"countLTR: {count}")
        if count != _EX_LTR:
            print(f"FAIL: LTR is {count}, but expected {_EX_LTR}")

    return count

def countDown(text):
    # down will have each character exactly LINELEN apart, forwards
    count = countByOffset(text, 0, LINELEN, LINELEN * 2, LINELEN * 3)

    if DEBUG:
        print(f"countDown: {count}")
        if count != _EX_DOWN:
            print(f"FAIL: DOWN is {count}, but expected {_EX_DOWN}")

    return count

def countUp(text):
    # up will have each character exactly LINELEN apart, backwards
    count = countByOffset(text, LINELEN * 3, LINELEN * 2, LINELEN, 0)

    if DEBUG:
        print(f"countUp: {count}")
        if count != _EX_UP:
            print(f"FAIL: UP is {count}, but expected {_EX_UP}")

    return count

def countDiagSE(text):
    # DiagSE will have increasing down and to the right, +10, +11, +12.
    count = countByOffset(text,
                          0,
                          LINELEN + 1,
                          (2*LINELEN) + 2,
                          (3*LINELEN) + 3)

    if DEBUG:
        print(f"countDiagSE: {count}")
        if count != _EX_DIAG_SE:
            print(f"FAIL: DIAG_SE is {count}, but expected {_EX_DIAG_SE}")

    return count

def countAll(text):
    count = 0
    count += countRTL(text)
    count += countLTR(text)
    count += countDown(text)
    count += countUp(text)
    count += countDiagSE(text)
    return count

def main():
    #input = getInput()
    i = getExampleInput()
    c = countAll(i)
    print(f"count is {c}")

if __name__ == "__main__":
    main()