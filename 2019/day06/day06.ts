import * as fs from 'fs';

const dayNum: string = "06";
const dayTitle: string = "Universal Orbit Map";

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

function orbits(lines: string[]): Map<string, string> {
    const m: Map<string, string> = new Map<string, string>();

    lines.forEach(l => {
        const planets: string[] = l.split(')');
        m.set(planets[1], planets[0]);
    });

    return m;
}

const lines: string[] = readInputSync('input.txt');

function part1() {
    const orbitMap = orbits(lines);
    const s: Set<string> = new Set<string>();
    orbitMap.forEach((v, k) => {
        s.add(k);
        s.add(v);
    });

    let count: number = 0;
    s.forEach(k => {
        let curr: string = k;
        let currbit: string = orbitMap.get(k);
        while (currbit) {
            count++;
            curr = currbit;
            currbit = orbitMap.get(curr);
        }
    });
    console.log('Part 1', count);
}

class Move {
    planet: string;
    count: number;

    constructor(p: string, c: number) {
        this.planet = p;
        this.count = c;
    }
}

function edges(lines: string[]): Map<string, string[]> {
    const e: Map<string, string[]> = new Map<string, string[]>();

    lines.forEach(l => {
        const planets: string[] = l.split(')');
        let list: string[] = e.get(planets[0]);
        if (!list) {
            list = [];
        }
        list.push(planets[1]);
        e.set(planets[0], list);

        list = e.get(planets[1]);
        if (!list) {
            list = [];
        }
        list.push(planets[0]);
        e.set(planets[1], list);
    });

    return e;
}

function part2() {
    const e: Map<string, string[]> = edges(lines);
    const visited: Set<string> = new Set<string>();
    const start: Move = new Move('YOU', 0);
    const queue: Move[] = [start];

    let answer: number = -1;
    while (queue.length !== 0) {
        const curr: Move = queue.shift();

        if (curr.planet === 'SAN') {
            answer = curr.count - 2;
            break;
        }

        const children: string[] = e.get(curr.planet);
        if (children) {
            children.forEach(c => {
                if (!visited.has(c)) {
                    const m: Move = new Move(c, curr.count + 1);
                    queue.push(m);
                    visited.add(c);
                }
            });
        }
    }

    console.log('Part 2', answer);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
