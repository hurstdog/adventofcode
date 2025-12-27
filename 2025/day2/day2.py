"""
This file contains a solution to day2 of 2025's advent of code.
"""

import sys

DEBUG = 1

"""
Takes an integer and returns True if it's a valid ID (no repeated sequences of numbers)
"""
def isValidID(id):
    idstr = str(id)
    l = len(idstr)
    midpoint = int(l / 2)
    p1 = idstr[:midpoint]
    p2 = idstr[midpoint:]
    if DEBUG and 0:
        print(f"Comparing {p1} to {p2}")
    if (p1 == p2):
        return False
    return True

def testIsValidID():
    cases = [
        (11, False),
        (22, False),
        (1, True),
        (9, True),
        (123123, False),
        (1010, False),
        (1188511885, False)        
    ]

    success = True
    for (id, expected) in cases:
        valid = isValidID(id)
        if valid == expected:
            if DEBUG:
                print(f"SUCCESS: {id} is {valid}")
        else:
            success = False
            if DEBUG:
                print(f"FAIL: {id} is {valid} expected {expected}")
                
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
    ranges = [
        ("11-22", 33),
        ("95-115", 99),
        ("998-1012", 1010),
        ("1188511880-1188511890", 1188511885),
        ("222220-222224", 222222),
        ("1698522-1698528", 0),
        ("446443-446449", 446446),
        ("38593856-38593862", 38593859),
        ("565653-565659", 0),
        ("824824821-824824827", 0),
        ("2121212118-2121212124", 0)
    ]

    success = True
    for (r, s) in ranges:
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
    part1sum = sumInvalidIDsInFullRange(ranges)
    print(f"Part1 sum: {part1sum}")

if __name__ == '__main__':
    main()

