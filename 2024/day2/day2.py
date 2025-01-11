
# open the file
f = open("2024/day2/input.txt", "r")

stoppoint = 3
safe_count = 0
unsafe_count = 0

# iterate over all of the lines
for line in f:
    #print(line)
    # for each line, split it into the parts
    parts = line.split()

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
        unsafe_count += 1
    elif unsafe_delta:
        #print("unsafe")
        unsafe_count += 1
    else:
        #print("safe")
        safe_count += 1

    # stop after 10 lines
    stoppoint -= 1
    if stoppoint <= 0:
        break

print("Safe lines: %d" % safe_count)
print("Unsafe lines: %d" % unsafe_count)