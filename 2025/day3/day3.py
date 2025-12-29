
DEBUG = 1

TEST = '''
987654321111111
811111111111119
234234234234278
818181911112111
'''

TEST_MAX_JOLTAGE = [98, 89, 78, 92]
TEST_TOTAL_JOLTAGE = 357

def getBatteryBanks(s):
    return s.rstrip('\n').lstrip('\n').split()

def getJoltageFromLine(line):
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

def testGetJoltageFromLine():
    lines = getBatteryBanks(TEST)
    if len(lines) != len(TEST_MAX_JOLTAGE):
        raise AssertionError(f"Expected {len(TEST_MAX_JOLTAGE)} lines, got {len(lines)}")
    
    for line, expected_joltage in zip(lines, TEST_MAX_JOLTAGE):
        actual_joltage = getJoltageFromLine(line)
        if actual_joltage != expected_joltage:
            raise AssertionError(f"Expected {expected_joltage}, got {actual_joltage} for line: {line}")

def getJoltageFromAllLines(lines):
    tot = 0
    for l in lines:
        # clean up newlines and skip any empty lines
        l = l.strip()
        if len(l) < 1:
            continue
        tot = tot + getJoltageFromLine(l)

    return tot

def testGetJoltageFromAllLines():
    lines = TEST.strip().split()
    tot = getJoltageFromAllLines(lines)
    if TEST_TOTAL_JOLTAGE != tot:
        raise AssertionError(f"Expected {TEST_TOTAL_JOLTAGE}, got {tot} for TEST input.")

def main():
    testGetJoltageFromLine()
    testGetJoltageFromAllLines()

    lines = open('2025/day3/input.txt', 'r').readlines()
    tot = getJoltageFromAllLines(lines)

    # Part 1 answer - 17107
    print(f"getJoltageFromAllLines(input.txt) == {tot}")

if __name__ == '__main__':
    main()