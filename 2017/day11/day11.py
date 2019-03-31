

print('Day 11')

def step_count(steps):
    n = 0
    nw = 0
    sw = 0
    s = 0
    se = 0
    ne = 0
    for step in steps:
        if step == 'n':
            n += 1
        elif step == 'nw':
            nw += 1
        elif step == 'sw':
            sw += 1
        elif step == 's':
            s += 1
        elif step == 'se':
            se += 1
        elif step == 'ne':
            ne +=1
        else:
            print("known step", step)

    ### now reduce them
    min_ns = 1
    min_nwse = 1
    min_nesw = 1
    min_nes = 1
    min_nws = 1
    min_sen = 1
    min_swn = 1
    min_swse = 1
    min_nwne = 1
    while min_ns!=0 and min_nwse!=0 and min_nesw!=0 and min_nes!=0 and min_nws!=0 and min_sen!=0 and min_swn!=0 and min_swse!=0 and min_nwne!=0:
        ## example 2
        # north south are opposites
        min_ns = min(n, s)
        n -= min_ns
        s -= min_ns
    
        # nw and se are opposites
        min_nwse = min(nw, se)
        nw -= min_nwse
        se -= min_nwse

        # ne and sw are opposites
        min_nesw = min(ne, sw)
        ne -= min_nesw
        sw -= min_nesw

        ## example 3
        # ne + s = se
        min_nes = min(ne, s)
        ne -= min_nes
        s -= min_nes
        se += min_nes
    
        # nw + s = sw
        min_nws = min(nw, s)
        nw -= min_nws
        s -= min_nws
        sw += min_nws
    
        # se + n = ne
        min_sen = min(se, n)
        se -= min_sen
        n -= min_sen
        ne += min_sen
    
        # sw + n = nw
        min_swn = min(sw, n)
        sw -= min_swn
        n -= min_swn
        nw += min_swn
    
        ## example 4
        # sw + se = s
        min_swse = min(sw,se)
        sw -= min_swse
        se -= min_swse
        s =+ min_swse

        # nw + ne = n
        min_nwne = min(nw, ne)
        nw -= min_nwne
        ne -= min_nwne
        n += min_nwne
    
    return n + nw + sw + s + se + ne


ex1 = "ne,ne,ne".split(',')
ex2 = "ne,ne,sw,sw".split(',')
ex3 = "ne,ne,s,s".split(',')
ex4 = "se,sw,se,sw,sw".split(',')
print("3:", step_count(ex1))
print("0:", step_count(ex2))
print("2:", step_count(ex3))
print("3:", step_count(ex4))


with open("input.txt") as f:
    path = [x.strip('\n') for x in f.readlines()]
    path = path[0].split(',')
    #print(path)


answer = step_count(path)
print("part1 answer", answer)

max_dist = 0
for i in range(1, len(path)):
    iter_path = path[:i]
    print(len(iter_path))
    dist = step_count(iter_path)
    if(dist > max_dist):
        max_dist = dist

print("furtherest away", max_dist)



