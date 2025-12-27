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
    if DEBUG:
        print(f"Comparing {p1} to {p2}")
    if (p1 == p2):
        return False
    return True

def countInvalidIDsInRange(range):
    return 0

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

def testCountInvalidIDsInRange():
    return 0

def main():
    if not testIsValidID():
        print("**** STOPPING EXECUTION AFTER FAILURES ****")
        sys.exit(1)

if __name__ == '__main__':
    main()

