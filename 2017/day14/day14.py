from day10 import dense_knot_hash
import binascii

print("Day 14")

def dkh(p):
    return dense_knot_hash(list(range(256)), p)


def toBin(hex):
    return "{0:8b}".format(int(hex,16))


puzzle = "amgozmfv"
count = 0
grid = set()
for i in range(128):
    p = puzzle + "-" + str(i)
    d = dkh(p)
    b = toBin(d)

    for x in range(len(b)):
        print(x)
        j = b[x]
        if(j == "1"):
            count += 1
            grid.add((i,j))

print("part1", count, len(grid))
print("grid", grid)
all_groups = list()
for i in range(128):
    for j in range(128):
        p = (i,j)
        if p in grid:
            print("p is in grid", p)
            # find the set within all_groups that p belongs too
            s = None
            for g in all_groups:
                if p in g:
                    s = g
            # if not found, then create a new group and add it to all_groups
            if s is None:
                s = set()
                all_groups.append(s)
            
            s.add(p)
            # check the neighbors
            n = [(i,j+1), (i,j-1), (i+1,j), (i-1,j)]
            for k in n:
                if k in grid:
                    s.add(k)
    
print("group count", len(all_groups))
