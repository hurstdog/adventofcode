
f = open('2024/day1/input.txt', 'r')

# how many lines to read
limiter = 100000

left = []
right = []

for line in f:
    # strip the newline
    line_s = line.strip()
    #print("I read line %s" % line_s)
    parts = line_s.split()
    left.append(parts[0])
    right.append(parts[1])
    limiter = limiter - 1
    if limiter <= 0:
        break

#print("Left: %s" % left)
#print("Right: %s" % right)

# sort each list
left.sort()
right.sort()

# iterate on each list, showing the difference
total_diff = 0
for i in range(len(left)):
    diff = abs(int(left[i]) - int(right[i]))
    total_diff = total_diff + diff
    #print("Difference: %s, total: %s" % (diff, total_diff))

print("Total difference: %s" % total_diff)