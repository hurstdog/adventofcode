# find all versions of XMAS in the text.

import re

INPUT = "2024/day4/input.txt"

# Return file as array of arrays, each character it's own element
def getFileInput():
    f = open(INPUT, "r")
    l = f.read()
    arr = []
    for line in l.splitlines():
        if len(line) == 0:
            continue
        arr.append(list(line))
    return arr

def matchCharsInArray(input_arr,
                      x_r, x_c, x_char,
                      m_r, m_c, m_char,
                      a_r, a_c, a_char,
                      s_r, s_c, s_char):
    # Don't go below zero
    if (x_r < 0 or x_c < 0 or
        m_r < 0 or m_c < 0 or
        a_r < 0 or a_c < 0 or
        s_r < 0 or s_c < 0):
        return False
    # if we bounce off the end of the main arrays, return false
    if (x_r >= len(input_arr) or
        m_r >= len(input_arr) or
        a_r >= len(input_arr) or
        s_r >= len(input_arr)):
        return False
    # if we bounce off the end of the subarrays, return false
    if (x_c >= len(input_arr[x_r]) or
        m_c >= len(input_arr[m_r]) or
        a_c >= len(input_arr[a_r]) or
        s_c >= len(input_arr[s_r])):
        return False
    
    return (input_arr[x_r][x_c] == x_char and
            input_arr[m_r][m_c] == m_char and
            input_arr[a_r][a_c] == a_char and
            input_arr[s_r][s_c] == s_char)

    return False

"""
Looks for XMAS at the given positions. Put into a separate function to
minimize error handling code elsewhere.

Returns True|False
"""
def matchInArray(input_arr, x_r, x_c, m_r, m_c, a_r, a_c, s_r, s_c):
    return matchCharsInArray(input_arr,
                             x_r, x_c, "X",
                             m_r, m_c, "M",
                             a_r, a_c, "A",
                             s_r, s_c, "S")

"""
Given a two-dimensional array of characters, goes through every element and looks for "XMAS" starting at the given element. The offsets specify where to find MAS, assuming the first element is X. This is why we only have the offsets for MAS, and not X.
"""
def countByOffsetInArray(input_arr, m_r, m_c, a_r, a_c, s_r, s_c):
    count = 0
    # for each row
    for r in range(len(input_arr)):
        # for each column
        for c in range(len(input_arr[r])):
            if input_arr[r][c] == "X":
                if matchInArray(input_arr,
                                r, c,               # x position
                                r + m_r, c + m_c,   # m position
                                r + a_r, c + a_c,   # a position
                                r + s_r, c + s_c):  # s position
                    count += 1
    return count

"""
a.b
.A.
c.d
Four positions, the only thing that changes is the letters. So build a function that takes a starting index and calls matchInArray for all combinations of crossing MAS.
Given an index and the example above, a will be at -1, -1, b at 1 , +1, c at -1, +1, and d at 1, 1.
"""
def xMasAtIndex(input_arr, r, c):
    count = 0

    if input_arr[r][c] != "A":
        return False

    # M.S
    # .A.
    # M.S
    if matchCharsInArray(input_arr,
                         r - 1, c - 1, "M",
                         r - 1, c + 1, "S",
                         r + 1, c - 1, "M",
                         r + 1, c + 1, "S"):
        return True
    
    # S.M
    # .A.
    # S.M
    if matchCharsInArray(input_arr,
                         r - 1, c - 1, "S",
                         r - 1, c + 1, "M",
                         r + 1, c - 1, "S",
                         r + 1, c + 1, "M"):
        return True
    
    # M.M
    # .A.
    # S.S
    if matchCharsInArray(input_arr,
                         r - 1, c - 1, "M",
                         r - 1, c + 1, "M",
                         r + 1, c - 1, "S",
                         r + 1, c + 1, "S"):
        return True
    # S.S
    # .A.
    # M.M
    if matchCharsInArray(input_arr,
                         r - 1, c - 1, "S",
                         r - 1, c + 1, "S",
                         r + 1, c - 1, "M",
                         r + 1, c + 1, "M"):
        return True

    return False

def countLTR(input_arr):
    # only increment the column counter
    return countByOffsetInArray(input_arr,
                                 0, 1,
                                 0, 2,
                                 0, 3)

def countRTL(input_arr):
    # only decrement the column counter
    return countByOffsetInArray(input_arr,
                                 0, -1,
                                 0, -2,
                                 0, -3)

def countDown(input_arr):
    # only increment the row counter
    return countByOffsetInArray(input_arr,
                                 1, 0,
                                 2, 0,
                                 3, 0)

def countUp(input_arr):
    # only decrement the row counter
    return countByOffsetInArray(input_arr,
                                 -1, 0,
                                 -2, 0,
                                 -3, 0)

def countDiagSE(input_arr):
    # DiagSE will have increasing down and to the right.
    return countByOffsetInArray(input_arr,
                                 1, 1,
                                 2, 2,
                                 3, 3)

def countDiagSW(input_arr):
    # DiagSW will have increasing down, decreasing columns.
    return countByOffsetInArray(input_arr,
                                 1, -1,
                                 2, -2,
                                 3, -3)

def countDiagNE(input_arr):
    # Row goes down, column goes up
    return countByOffsetInArray(input_arr,
                                 -1, 1,
                                 -2, 2,
                                 -3, 3)

def countDiagNW(input_arr):
    # Row goes down, column goes down
    return countByOffsetInArray(input_arr,
                                 -1, -1,
                                 -2, -2,
                                 -3, -3)

def countAll(input_arr):
    count = 0
    count += countLTR(input_arr)
    count += countRTL(input_arr)
    count += countDown(input_arr)
    count += countUp(input_arr)
    count += countDiagSE(input_arr)
    count += countDiagNE(input_arr)
    count += countDiagSW(input_arr)
    count += countDiagNW(input_arr)
    return count

def countXmas(input_arr):
    count = 0

    # for each row
    for r in range(len(input_arr)):
        # for each column
        for c in range(len(input_arr[r])):
            if xMasAtIndex(input_arr, r, c):
                count += 1

    return count

def main():
    i = getFileInput()              
    c = countAll(i)                 # Part 1 - 2358
    print(f"count is {c}")
    c = countXmas(i)                # Part 2 - 1737
    print(f"countXmas is {c}")

if __name__ == "__main__":
    main()