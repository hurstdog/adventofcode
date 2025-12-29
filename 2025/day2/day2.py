"""
This file contains a solution to day2 of 2025's advent of code.
"""

import sys

DEBUG = 0
PART = 1

"""
Given a list, returns True if every entry in the list is equal, False otherwise
"""
def allEqual(l):
    if DEBUG and 0:
        print(f"Checking list {l}")

    if len(l) < 1:
        return True

    first = l[0]
    for i in l[1:]:
        if i != first:
            return False
    
    return True

"""
Takes an integer and returns True if it's a valid ID (no repeated sequences of numbers)
"""
def isValidID(id):
    idstr = str(id)
    l = len(idstr)

    # Split the string into equal segments for part 1
    if PART == 1:
        # Skip odd length integers
        if l % 2 != 0:
            return True

        parts = []
        midpoint = int(l / 2)
        parts.append(idstr[:midpoint])
        parts.append(idstr[midpoint:])
        return not allEqual(parts)
    else:
        # split the string into equal lists, starting from length 1 up through len/2
        listlen = 1
        while listlen < (int(l / 2) + 1):
            testlist = []
            for i in range(0, l, listlen):
                testlist.append(idstr[i:i+listlen])
            if allEqual(testlist):
                return False
            listlen = listlen + 1
    return True

def testIsValidID():
    # [(id, part1valid, part2valid)...]
    cases = [
        (11, False, False),
        (22, False, False),
        (1, True, True),
        (9, True, True),
        (123123, False, False),
        (1010, False, False),
        (1188511885, False, False),
        (123123123, True, False),
        (1212121212, True, False),
        (1111111, True, False)
    ]

    success = True
    for (id, expected1, expected2) in cases:
        exp = expected1
        if PART == 2:
            exp = expected2
        valid = isValidID(id)
        if valid == exp:
            if DEBUG:
                print(f"SUCCESS: {id} is {valid}")
        else:
            success = False
            if DEBUG:
                print(f"FAIL: {id} is {valid} expected {exp}")
                
    return success

def rangeToSequence(r):
    # split on the -
    (start, end) = r.split("-")

    # int the first
    istart = int(start)
    iend = int(end)

    ret = []
    i = int(start)
    while i <= iend:
        ret.append(i)
        i = i + 1

    # return the array
    return ret

def sumInvalidIDsInRange(r):
    seq = rangeToSequence(r)
    tot = 0
    for num in seq:
        if not isValidID(num):
            tot = tot + num

    return tot

def testSumInvalidIDsInRange():
    # [(id, part1sum, part2sum)...]
    ranges = [
        ("11-22", 33, 33),
        ("95-115", 99, 210),
        ("998-1012", 1010, 2009),
        ("1188511880-1188511890", 1188511885, 1188511885),
        ("222220-222224", 222222, 222222),
        ("1698522-1698528", 0, 0),
        ("446443-446449", 446446, 446446),
        ("38593856-38593862", 38593859, 38593859),
        ("565653-565659", 0, 565656),
        ("824824821-824824827", 0, 824824824),
        ("2121212118-2121212124", 0, 2121212121)
    ]

    success = True
    for (r, s1, s2) in ranges:
        s = s1
        if PART == 2:
            s = s2
        res = sumInvalidIDsInRange(r)
        if res == s:
            if DEBUG:
                print(f"SUCCESS: Sum of invalid ids in {r} is {s}")
        else:
            success = False
            if DEBUG:
                print(f"FAIL: Sum of invalid ids in {r} is {res}, expected {s}")

    return success

def sumInvalidIDsInFullRange(r):
    ranges = r.split(",")
    s = 0
    for r in ranges:
        s = s + sumInvalidIDsInRange(r)

    return s

def testSumInvalidIDsInFullRange():
    r = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
    s = sumInvalidIDsInFullRange(r)
    expected = 33 + 99 + 1010 + 1188511885 + 222222 + 446446 + 38593859
    if PART == 2:
        expected = 33 + 210 + 2009 + 1188511885 + 222222 + 446446 + 38593859 + 565656 + 824824824 + 2121212121

    if s != expected:
        if DEBUG:
            print(f"FAIL: Expected {expected} got {s}")
            return False
    else:
        if DEBUG:
            print(f"SUCCESS: testSumInvalidIDsInFullRange() got {s}")
    
    return True

def main():
    if not testIsValidID() or not testSumInvalidIDsInRange() or not testSumInvalidIDsInFullRange():
        print("**** STOPPING EXECUTION AFTER FAILURES ****")
        sys.exit(1)

    # open the file
    f = open('2025/day2/input.txt', 'r')
    ranges = f.readline().rstrip('\n')
    global PART
    for p in [1, 2]:
        PART = p
        theSum = sumInvalidIDsInFullRange(ranges)
        print(f"Part{PART} sum: {theSum}")

    # Part 1 Solution: 8576933996
    # Part 2 Solution: 25663320831

if __name__ == '__main__':
    main()

