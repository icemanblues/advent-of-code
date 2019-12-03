import fs from 'fs';

const dayNum: string = "03";
const dayTitle: string = "Crossed Wires";

function readFile(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/\r?\n/);
    return lines;
}

class Point {
    x: number;
    y: number;

    constructor(x: number, y: number) {
        this.x = x;
        this.y = y;
    }

    toString(): string {
        return `${this.x},${this.y}`;
    }
}

function minManhattanDist(wire1: string[], wire2: string[]): number {
    const grid1: Set<string> = new Set<string>(wireList(wire1));
    const grid2: Set<string> = new Set<string>(wireList(wire2));

    const intersect: Set<String> = new Set([...grid1].filter(i => grid2.has(i)));

    let dist: number = 0;
    intersect.forEach(p => {
        let pp: number[] = p.split(/,/).map(s => Number(s));
        const d: number = pp[0] + pp[1];
        if (dist === 0) {
            dist = d;
        } else if (d < dist) {
            dist = d;
        }
    });

    return dist;
}

function part1() {
    console.log('Part 1');

    const lines: string[] = readFile('input.txt');
    const wire1: string[] = lines[0].split(/,/);
    const wire2: string[] = lines[1].split(/,/);
    console.log(minManhattanDist(wire1, wire2));
}

function wireList(wire: string[]): string[] {
    const l: string[] = [];
    let step: number = 0;
    let curr: Point = new Point(0, 0);

    wire.forEach(str => {
        const dir: string = str.charAt(0);
        const len: number = Number(str.slice(1));

        switch (dir) {
            case 'U':
                //y+1
                for (let i: number = 0; i < len; i++) {
                    curr = new Point(curr.x, curr.y + 1);
                    l.push(curr.toString());
                }

                break;
            case 'D':
                // y-1
                for (let i: number = 0; i < len; i++) {
                    curr = new Point(curr.x, curr.y - 1);
                    l.push(curr.toString());
                }
                break;
            case 'L':
                // x-1
                for (let i: number = 0; i < len; i++) {
                    curr = new Point(curr.x - 1, curr.y);
                    l.push(curr.toString());
                }
                break;
            case 'R':
                // x+1
                for (let i: number = 0; i < len; i++) {
                    curr = new Point(curr.x + 1, curr.y);
                    l.push(curr.toString());
                }
                break;

        }
    });

    return l;
}


function stepCount(wire1: string[], wire2: string[]): number {
    const l1: string[] = wireList(wire1);
    const l2: string[] = wireList(wire2);
    const grid1: Set<string> = new Set(l1);
    const grid2: Set<string> = new Set(l2);

    const intersect: Set<string> = new Set([...grid1].filter(i => grid2.has(i)));

    let stepCount: number = 0;
    intersect.forEach(p => {
        const s1: number = l1.indexOf(p);
        const s2: number = l2.indexOf(p);
        const c: number = s1 + s2;
        if (stepCount === 0) {
            stepCount = c;
        } else if (c < stepCount) {
            stepCount = c;
        }
    });

    // the +2 is that I am not counting (0,0) for both wires
    return stepCount + 2;
}

function part2() {
    console.log('Part 2');

    const lines: string[] = readFile('input.txt');
    const wire1: string[] = lines[0].split(/,/);
    const wire2: string[] = lines[1].split(/,/);
    console.log(stepCount(wire1, wire2));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle} `);

    part1();
    part2();
}

main();
