
f = open('2024/day1/input.txt', 'r')

# how many lines to read
limiter = 100000

left = []
right = []
right_count = {}

for line in f:
    # strip the newline
    line_s = line.strip()
    #print("I read line %s" % line_s)
    parts = line_s.split()
    left.append(parts[0])
    right.append(parts[1])

    # add counts of the number of times we've seen numbers on the right
    if right_count.get(parts[1]) is None:
        right_count[parts[1]] = 1
    else:
        right_count[parts[1]] = right_count[parts[1]] + 1

    limiter = limiter - 1
    if limiter <= 0:
        break

#print("Left: %s" % left)
#print("Right: %s" % right)
#print("Right count: %s" % right_count)

### part 1 solution section
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

### part 2 solution section
# Iterate over the left list, multiplying by the count on the right list
similarity = 0
for i in range(len(left)):
    num = left[i]
    multiplier = 0
    if right_count.get(num) is not None:
        multiplier = right_count[num]
    #print("num: %s; multiplier: %s" % (num, multiplier))
    similarity = similarity + (int(num) * multiplier)

print("Part 2 Similarity: %s" % similarity)