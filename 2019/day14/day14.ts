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
const TRILLION: number = 1000000000000;

function leastOre(rm: Map<string, Reaction>, fuel: number): number {
    const needsMap: Map<string, number> = new Map([[FUEL, fuel]]);
    const producesMap: Map<string, number> = new Map();

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

        r.inputs.forEach((v, k) => {
            const consumes: number = v * m;
            needsMap.set(k, getOrDefault(needsMap, k, 0) + consumes);
            queue.push(k);
        });
    }

    return needsMap.get(ORE);
}

function mostFuel(rm: Map<string, Reaction>): number {
    let [lower, upper] = [0, TRILLION];
    let f: number = Math.round((lower + upper) / 2);
    let ore: number = 0;

    let bestFuel: number = -1;
    while (ore !== TRILLION) {
        const ore: number = leastOre(rm, f);

        if (ore === TRILLION) {
            return f;
        }
        else if (ore < TRILLION) {
            if (f == lower) {
                break;
            }

            [lower, upper] = [f, upper];
            bestFuel = f;
        } else {
            if (f == upper) {
                break;
            }

            [lower, upper] = [lower, f];
        }
        f = Math.round((lower + upper) / 2);
    }

    return bestFuel;
}

function part1() {
    console.log('Part 1');
    console.log('Part 1', leastOre(parseReactions('input.txt'), 1));
}

function part2() {
    console.log('Part 2');
    console.log('Part 2', mostFuel(parseReactions('input.txt')));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
