"""
Algorithmic approaches for finding the largest 12-digit number
by removing (n-12) digits from a string of n digits.
"""

def greedy_stack_algorithm(line, target_length=12):
    """
    OPTIMAL: Greedy approach using a stack.
    
    The idea: Use a stack to greedily keep larger digits at more significant positions.
    - We need to remove (len(line) - target_length) digits
    - For each digit, if we can still remove more digits and the current digit
      is greater than the last kept digit, remove the smaller ones first
    - This maximizes digits at the most significant positions
    
    Time: O(n), Space: O(n)
    """
    if len(line) <= target_length:
        return line
    
    to_remove = len(line) - target_length
    stack = []
    
    for digit in line:
        # While we can still remove digits and the current digit is larger
        # than what we've kept, remove the smaller ones to make room
        while stack and to_remove > 0 and digit > stack[-1]:
            stack.pop()
            to_remove -= 1
        stack.append(digit)
    
    # If we still need to remove more (happens when digits are non-increasing),
    # remove from the end (least significant positions)
    while to_remove > 0:
        stack.pop()
        to_remove -= 1
    
    return ''.join(stack)


def recursive_backtracking(line, target_length=12):
    """
    BRUTE FORCE: Try all combinations recursively.
    
    This explores all possible ways to keep target_length digits.
    Time: O(C(n, target_length)) = exponential
    Space: O(target_length) for recursion stack
    
    NOT FEASIBLE for n=40, target_length=12 (5+ billion combinations)
    """
    if target_length == 0:
        return ''
    if len(line) == target_length:
        return line
    
    # Option 1: Skip current digit
    skip = recursive_backtracking(line[1:], target_length)
    
    # Option 2: Keep current digit
    keep = line[0] + recursive_backtracking(line[1:], target_length - 1)
    
    return max(skip, keep) if skip else keep


def dynamic_programming(line, target_length=12):
    """
    DYNAMIC PROGRAMMING: DP[i][j] = largest number with j digits from first i digits.
    
    Time: O(n * target_length * n) = O(n^2 * target_length) for string comparisons
    Space: O(n * target_length)
    
    More complex than greedy, and slower. Not recommended here.
    """
    n = len(line)
    # DP[i][j] = best string of length j from first i characters
    dp = [['' for _ in range(target_length + 1)] for _ in range(n + 1)]
    
    for i in range(1, n + 1):
        for j in range(1, min(i + 1, target_length + 1)):
            # Option 1: Don't include line[i-1]
            option1 = dp[i-1][j]
            
            # Option 2: Include line[i-1]
            option2 = dp[i-1][j-1] + line[i-1]
            
            # Take the larger one (string comparison works for equal-length strings)
            dp[i][j] = max(option1, option2) if option1 else option2
    
    return dp[n][target_length]


def sliding_window_greedy(line, target_length=12):
    """
    VARIATION: Sliding window approach - at each position, decide whether to keep
    the current digit based on what's ahead.
    
    Similar to greedy but uses a different perspective:
    - For each position, check if removing current digit allows keeping
      a larger digit later
    - Time: O(n^2) in worst case
    - Space: O(n)
    """
    to_remove = len(line) - target_length
    result = []
    i = 0
    
    while i < len(line) and to_remove > 0:
        # Look ahead to see if there's a larger digit we could get by removing current
        # Find the maximum digit in the next (to_remove + 1) positions
        lookahead_end = min(i + to_remove + 1, len(line))
        max_digit = max(line[i:lookahead_end])
        max_pos = line[i:lookahead_end].index(max_digit) + i
        
        if line[i] < max_digit:
            # Skip digits until we reach the max
            to_remove -= (max_pos - i)
            i = max_pos
        else:
            result.append(line[i])
            i += 1
    
    # Add remaining digits if we haven't reached target_length
    while i < len(line) and len(result) < target_length:
        result.append(line[i])
        i += 1
    
    # Remove from end if we have too many
    while len(result) > target_length:
        result.pop()
    
    return ''.join(result)


# Test cases
if __name__ == '__main__':
    test_cases = [
        ('987654321111111', '987654321111'),
        ('811111111111119', '811111111119'),
        ('234234234234278', '434234234278'),
        ('818181911112111', '818181911112'),
    ]
    
    print("Testing Greedy Stack Algorithm (RECOMMENDED):")
    for line, expected in test_cases:
        result = greedy_stack_algorithm(line, 12)
        status = "✓" if result == expected else "✗"
        print(f"{status} {line} -> {result} (expected {expected})")
    
    print("\nAll algorithms produce optimal results, but:")
    print("1. Greedy Stack: O(n) time, optimal, recommended")
    print("2. Recursive: Exponential, not feasible for large n")
    print("3. DP: O(n^2) time, more complex, unnecessary")
    print("4. Sliding Window: O(n^2) worst case, similar to greedy but slower")

