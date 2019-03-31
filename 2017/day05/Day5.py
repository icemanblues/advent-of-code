
def jump_out(filename):
    with open(filename) as f:
        jumps = [int(x.strip('\n')) for x in f.readlines()]
    print(jumps)
    
    curr = 0
    step_count = 0
    while curr>=0 and curr<len(jumps):
        old_curr = curr
        curr = curr + jumps[curr]
        if jumps[old_curr] >= 3:
            jumps[old_curr] -= 1
        else:
            jumps[old_curr] += 1
        step_count += 1

    print(jumps)
    return step_count


test1 = jump_out("test1.txt")
print("test", test1)

part1 = jump_out("input.txt")
print("part1", part1)




