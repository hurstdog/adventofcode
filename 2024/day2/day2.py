import time

start_time = time.time()

# open the file
f = open("2024/day2/input.txt", "r")

stoppoint = 10000
safe_count = 0
unsafe_count = 0

"""
Takes an array of strings representing numbers, and returns if it's 
safe or not.
Safe is all increasing or descreasing gradually, no more than 3 or < 1.
"""
def isItSafe(parts):
    prev_num = 0   # warning: if there's a zero in the input, this will fail
    line_increase = False
    line_decrease = False
    unsafe_delta = False

    for p in parts:
        #print("found number %d" % int(p))
        
        if prev_num != 0:
            # now I can compare the jump
            diff = abs(prev_num - int(p))
            if diff > 3 or diff < 1:
                #print("unsafe delta")
                unsafe_delta = True
                break

            # now I can compare the direction
            if prev_num < int(p):
                #print("increasing")
                line_increase = True
            elif prev_num > int(p):
                #print("decreasing")
                line_decrease = True

        prev_num = int(p)
    
    # is it safe?
    # gradually decreasing or increasing
    # change at least one, never more than three
    if line_increase and line_decrease:
        #print("unsafe")
        return False
    elif unsafe_delta:
        #print("unsafe")
        return False
    else:
        #print("safe")
        return True         # yay!

# iterate over all of the lines
for line in f:
    #print(line)
    # for each line, split it into the parts
    parts = line.split()

    all_safe = isItSafe(parts)
    if all_safe:
        safe_count += 1
    else:
        # uh oh, try to remove one and find a safe way
        safe = False
        for i in range(0, len(parts)):
            #print("trying to remove %d" % i)
            new_parts = parts[:i] + parts[i+1:]
            #print("new parts: %s" % new_parts)
            safe = isItSafe(new_parts)
            if safe:
                safe_count += 1
                break

        # if we still didn't find a safe way, it's unsafe
        if not safe:
            unsafe_count += 1

    # stop after stopppoint lines
    stoppoint -= 1
    if stoppoint <= 0:
        break

# Part 1 solution: 224 safe, 776 unsafe
# Part 2 solution: 293 safe, 707 unsafe
print("Safe lines: %d" % safe_count)
print("Unsafe lines: %d" % unsafe_count)

end_time = time.time()
print("Runtime: %f" % (end_time - start_time))