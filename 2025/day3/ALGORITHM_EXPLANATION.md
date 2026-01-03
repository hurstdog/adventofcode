# Algorithm Options for Finding Largest 12-Digit Sequence

## Problem
Given a string of n digits (n ≤ 40), remove (n-12) digits to maximize the remaining 12-digit number.

## Algorithm Options

### 1. **Greedy Stack Algorithm** ⭐ RECOMMENDED
**Time:** O(n)  
**Space:** O(n)  
**Optimal:** Yes

**How it works:**
- Use a stack to build the result
- For each digit in the input:
  - While we can still remove digits AND the current digit > last kept digit:
    - Remove the smaller digit (pop from stack) to make room for the larger one
    - This ensures larger digits are at more significant positions
- If we haven't removed enough digits yet, remove from the end (least significant)

**Python implementation:**
```python
def greedy_max_digits(line, target_length=12):
    if len(line) <= target_length:
        return line
    
    to_remove = len(line) - target_length
    stack = []
    
    for digit in line:
        while stack and to_remove > 0 and digit > stack[-1]:
            stack.pop()
            to_remove -= 1
        stack.append(digit)
    
    # Remove any remaining excess from the end
    while to_remove > 0:
        stack.pop()
        to_remove -= 1
    
    return ''.join(stack)
```

**Why it works:**
- Maximizes digits at the most significant positions first
- Greedy choice is safe: if we can swap a smaller digit for a larger one earlier, we should
- Linear time, optimal result

---

### 2. **Recursive Backtracking** (Brute Force)
**Time:** O(C(n, 12)) = exponential  
**Space:** O(12) for recursion stack  
**Optimal:** Yes, but NOT FEASIBLE

**How it works:**
- Try all possible combinations of which digits to keep
- For each position: try keeping it OR skipping it
- Return the maximum result

**Why NOT to use:**
- For n=40, target=12: C(40,12) = 5,586,853,480 combinations
- Would take millions of seconds even at 1000 operations/second
- Exponential complexity makes it impractical

---

### 3. **Dynamic Programming**
**Time:** O(n² × 12) for string operations  
**Space:** O(n × 12)  
**Optimal:** Yes, but OVERKILL

**How it works:**
- DP[i][j] = best string of length j from first i digits
- For each position, decide: include current digit or not
- More complex than greedy, slower, unnecessary here

**Why NOT to use:**
- More complex to implement
- Slower than greedy (quadratic vs linear)
- Unnecessary complexity for this problem

---

### 4. **Your Current Heuristic** (Remove by Digit Value)
**Time:** O(n × 9) = O(n)  
**Space:** O(n)  
**Optimal:** No ❌

**How it works:**
- Iterate through digits 1, 2, 3, ..., 9
- Remove all occurrences of smaller digits first
- Continue until only 12 digits remain

**Why it's suboptimal:**
- Doesn't consider digit positions
- Example: "1234" - removing all 1s gives "234", but keeping "234" might not be optimal if there's a larger digit later
- The position of a digit matters more than just its value
- May miss better combinations

**Example where it fails:**
- Input: "12321", target length 3
- Your approach: Remove all 1s → "232"
- Optimal: Could be "321" if we consider positions differently

---

## Recommendation

**Use the Greedy Stack Algorithm (#1)** because:
1. ✅ Optimal result (guaranteed maximum)
2. ✅ Linear time O(n) - fast even for n=40
3. ✅ Simple to implement and understand
4. ✅ Standard solution for "remove k digits to maximize/minimize" problems
5. ✅ Space efficient O(n)

The greedy approach works because:
- We want larger digits at more significant (leftmost) positions
- When we see a larger digit, we should prefer it over smaller digits we've already kept
- The stack allows us to "undo" previous choices when we find something better

---

## Implementation Tips

When choosing which digits to drop in a loop:

1. **Greedy Stack (Recommended):** 
   - Loop through digits once
   - Use a stack to track what to keep
   - Compare current digit with last kept digit
   - Drop smaller digits when you find larger ones

2. **Avoid:**
   - Trying all combinations (exponential)
   - Removing digits purely by value (ignores position)
   - Complex DP (unnecessary overhead)

The key insight: **Position matters more than absolute digit value** when maximizing a number.

