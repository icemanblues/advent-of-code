import fs from 'fs';

const dayNum: string = "14";
const dayTitle: string = "Space Stoichiometry";

function getOrDefault<K, V>(m: Map<K, V>, k: K, d: V): V {
    return m.has(k) ? m.get(k) : d;
}

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

class Reaction {
    inputs: Map<string, number>;
    reagent: string;
    count: number;

    constructor(i: Map<string, number>, r: string, c: number) {
        this.inputs = i;
        this.reagent = r;
        this.count = c;
    }
}

function parseReactions(filename: string): Map<string, Reaction> {
    const rm: Map<string, Reaction> = new Map();

    readInputSync(filename).forEach(l => {
        let [input, output] = l.split(/=>/)
        function mapper(s: string): Map<string, number> {
            const m: Map<string, number> = new Map();
            s.split(/,/).forEach(i => {
                let [count, reagent] = i.trim().split(/ /);
                m.set(reagent, Number(count));
            });
            return m;
        }
        let [oc, or] = output.trim().split(/ /);
        rm.set(or, new Reaction(mapper(input), or, Number(oc)));
    });
    return rm;
}

const ORE: string = 'ORE';
const FUEL: string = 'FUEL';

function leastOre(rm: Map<string, Reaction>): number {
    const producesMap: Map<string, number> = new Map();

    const needsMap: Map<string, number> = new Map([[FUEL, 1]]);
    const queue: string[] = [FUEL];
    while (queue.length !== 0) {
        const curr: string = queue.shift();
        if (ORE === curr) {
            continue;
        }

        const currNeed: number = needsMap.get(curr);
        const r: Reaction = rm.get(curr);
        const currHave: number = getOrDefault(producesMap, curr, 0);

        const m: number = Math.ceil((currNeed - currHave) / r.count);
        const produces: number = m * r.count;
        producesMap.set(curr, getOrDefault(producesMap, curr, 0) + produces);

        const reactionNeeds: Map<string, number> = new Map();
        r.inputs.forEach((v, k) => {
            const consumes: number = v * m;
            needsMap.set(k, getOrDefault(needsMap, k, 0) + consumes);
            const idx: number = queue.indexOf(k);
            queue.push(k);
        });
    }

    return needsMap.get(ORE);
}

function part1() {
    console.log('Part 1');
    console.log('Part 1', leastOre(parseReactions('input.txt')));
}

function part2() {
    console.log('Part 2');
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
