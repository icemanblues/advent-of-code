

def redistribute(l):
    visited = dict()
    count = 0

    while tuple(l) not in visited:
        visited[tuple(l)] = count
        #print(l)
        
        # find the index of the max
        m = l[0]
        index_max = 0
        for i in range(1, len(l)):
            if l[i] > m:
                index_max = i
                m = l[i]

        #print("max and index", m, index_max)
        # redistribute the wealth
        l[index_max]=0
        for i in range(m):
            index_max +=1
            l[index_max % len(l)] += 1

        # increment the step counter
        count += 1
        #print("count", count)

    first_seen = visited[tuple(l)]
    return count - first_seen
        

test_bank = [0, 2, 7, 0]
test_count = redistribute(test_bank)
print(test_bank, test_count)

with open("input.txt") as f:
    content = [x.strip('\n') for x in f.readlines()]
    input_bank = content[0].split()
    input_bank = [int(x) for x in input_bank]

print(input_bank)
part1_answer = redistribute(input_bank)
print("part1 answer", part1_answer)
