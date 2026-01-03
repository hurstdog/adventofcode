
DEBUG = 1

TEST = '''
987654321111111
811111111111119
234234234234278
818181911112111
'''

TEST_PART = 2
TEST_MAX_JOLTAGE = [98, 89, 78, 92]
TEST_TOTAL_JOLTAGE = 357
TEST2_MAX_JOLTAGE = [987654321111, 811111111119, 434234234278, 818181911112]
TEST2_TOTAL_JOLTAGE = 987654321111 + 811111111119 + 434234234278 + 818181911112

def getBatteryBanks(s):
    return s.rstrip('\n').lstrip('\n').split()

def getPart1JoltageFromLine(line):
    l = len(line)
    # Iterate over the list, keeping the highest (a) and second-highest (b) number seen
    a = int(line[0])
    b = -1
    for i in range(1, l-1):
        cur = int(line[i])
        if cur > a:
            a = cur
            b = -1
        elif cur > b:
            b = cur
    
    if b < int(line[-1]):
        b = int(line[-1])
    
    joltage = int(str(a) + str(b))
    if DEBUG:
        print(f"joltage({line}) == {joltage}")
    return joltage

def testGetPart1JoltageFromLine():
    lines = getBatteryBanks(TEST)
    if len(lines) != len(TEST_MAX_JOLTAGE):
        raise AssertionError(f"Expected {len(TEST_MAX_JOLTAGE)} lines, got {len(lines)}")
    
    for line, expected_joltage in zip(lines, TEST_MAX_JOLTAGE):
        actual_joltage = getPart1JoltageFromLine(line)
        if actual_joltage != expected_joltage:
            raise AssertionError(f"Expected {expected_joltage}, got {actual_joltage} for line: {line}")

def dropDigits(s, d, c):
    """
    Strips up to c occurrences of digit d from the string s, starting from the left.
    Only digits matching the character d are removed; removal is left-to-right only.
    """
    if DEBUG:
        print(f"dropDigits({s}, {d}, {c})")
    out = []
    count = 0
    for ch in s:
        if ch == d and count < c:
            count += 1
            continue
        out.append(ch)
    return ''.join(out)

def testDropDigits():
    # Remove up to 3 of '1' from the left
    assert dropDigits('111234111', '1', 3) == '234111'
    # Remove up to 1 of '2'
    assert dropDigits('124242', '2', 1) == '14242'
    # Remove 0 of '4'
    assert dropDigits('4433', '4', 0) == '4433'
    # Remove more than possible
    assert dropDigits('8888', '8', 10) == ''
    # No such digit in s
    assert dropDigits('12345', '9', 2) == '12345'
    # Remove multiple, digit not at start
    assert dropDigits('123123123', '3', 2) == '1212123'
    # Remove up to c, but not all
    assert dropDigits('111222', '1', 2) == '1222'
    # Removing '0' when not present
    assert dropDigits('987654', '0', 1) == '987654'
    # Remove from middle
    assert dropDigits('abc1def1ghi', '1', 1) == 'abcdef1ghi'
    print("All tests passed for dropDigits.")

def getPart2JoltageFromLine(line):
    l = len(line)

    digit = 1
    joltagelen = 0
    workingline = line
    while len(workingline) > 12:
        maxdrop = len(workingline) - 12
        workingline = dropDigits(workingline, str(digit), maxdrop)
        digit += 1

    joltage = int(workingline)

    return joltage

def testGetPart2JoltageFromLine():
    lines = getBatteryBanks(TEST)
    if len(lines) != len(TEST2_MAX_JOLTAGE):
        raise AssertionError(f"Expected {len(TEST2_MAX_JOLTAGE)} lines, got {len(lines)}")
    
    for line, expected_joltage in zip(lines, TEST2_MAX_JOLTAGE):
        actual_joltage = getPart2JoltageFromLine(line)
        if actual_joltage != expected_joltage:
            raise AssertionError(f"Expected {expected_joltage}, got {actual_joltage} for line: {line}")

def getJoltageFromAllLines(lines):
    tot = 0
    for l in lines:
        # clean up newlines and skip any empty lines
        l = l.strip()
        if len(l) < 1:
            continue
        tot = tot + getPart1JoltageFromLine(l)

    return tot

def testGetJoltageFromAllLines():
    lines = TEST.strip().split()
    tot = 0
    tot = getJoltageFromAllLines(lines)
    expected = TEST_TOTAL_JOLTAGE
    if TEST_PART == 2:
        expected = TEST2_TOTAL_JOLTAGE

    if expected != tot:
        raise AssertionError(f"Expected {TEST_TOTAL_JOLTAGE}, got {tot} for TEST input.")

def main():
    testDropDigits()
    testGetPart1JoltageFromLine()
    testGetPart2JoltageFromLine()
    testGetJoltageFromAllLines()

    lines = open('2025/day3/input.txt', 'r').readlines()
    tot = getJoltageFromAllLines(lines)

    # Part 1 answer - 17107
    print(f"getJoltageFromAllLines(input.txt) == {tot}")

if __name__ == '__main__':
    main()