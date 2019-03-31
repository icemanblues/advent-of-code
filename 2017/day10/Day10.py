test_list = [0, 1, 2, 3, 4]
test_inst = [3, 4, 1, 5]


def swap(list, i, j):
    temp = list[i]
    list[i] = list[j]
    list[j] = temp


def reverse(list, curr, inst):
    end = curr + inst - 1
    num_swaps = int( (inst)/2 )
    for x in range(num_swaps):
        i = (curr + x) % len(list)
        j = (end - x) % len(list)
        swap(list, i, j)


def knot_hash(l, inst):
    curr = 0
    skip_size = 0
    
    for i in inst:
        #print(l)
        reverse(l, curr, i)
        curr = curr + i + skip_size
        skip_size += 1
        #print("curr:", curr, " skip_size:", skip_size)

    return l[0] * l[1]

        
#reverse(test_list, 0, 3)
answer = knot_hash(test_list, [3, 4, 1, 5])
#print(test_list)
#print("answer:",answer)

input_list = list(range(256))
input_inst = [102,255,99,252,200,24,219,57,103,2,226,254,1,0,69,216]
answer = knot_hash(input_list, input_inst)
print("part1 answer:",answer)


my_suffix = [17, 31, 73, 47, 23]

def asciiCode(s):
    b = []
    for i in s:
        b.append(ord(i))
    return b


def toHex(s):
    return "%0.2x" % s


def dense_knot_hash(l, inst):
    a = asciiCode(inst)
    ass = a + my_suffix
    as64 = ass * 64
    knot_hash(l, as64)

    # XOR the 16th elements
    dense_hash = list()
    prev = 0
    for i in range(16,len(l)+1,16):
        #print ("xor slice", prev, i-1)
        slice = l[prev:i]
        acc = slice[0]
        for j in range(1, len(slice)):
            #print(j)
            acc = acc ^ slice[j]
        # things to do before iterating
        dense_hash.append(acc)
        prev = i

    print("length of dense_hash", len(dense_hash))
    ## convert dense hash to hex
    result = ""
    for i in dense_hash:
        result = result + toHex(i)
    return result


# test cases for part 2
test2_inst = "1,2,3"
test2_empty = ""
test2_aoc = "AoC 2017"
test2_124 = "1,2,4"
print(test2_empty, dense_knot_hash(list(range(256)), test2_empty))
print(test2_inst, dense_knot_hash(list(range(256)), test2_inst))
print(test2_aoc, dense_knot_hash(list(range(256)), test2_aoc))
print(test2_124, dense_knot_hash(list(range(256)), test2_124))


part2_list = list(range(256))
part2_question = "102,255,99,252,200,24,219,57,103,2,226,254,1,0,69,216"
answer = dense_knot_hash(part2_list, part2_question)
print("part 2 answer:", answer)
