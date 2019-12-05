import fs from 'fs';

const dayNum: string = "05";
const dayTitle: string = "Sunny with a Chance of Asteroids";

// reads entire file into memory, then splits it
function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/,/);
    return lines;
}

function opMode(inst: number): [number, boolean[]] {
    const s: string = String(inst);
    const op: number = Number(s.slice(s.length - 2));
    const mode: boolean[] = [];

    for (let i = s.length - 3; i >= 0; i--) {
        mode.push(s.charAt(i) !== '0'); // true is immediate
    }
    if (mode.length !== 3) {
        mode.push(false);
    }

    return [op, mode];
}

function lookup(intcodes: number[], index: number, mode: boolean): number {
    if (mode) {
        return intcodes[index];
    }

    return intcodes[intcodes[index]];

}

function prog(intcodes: number[], input: number): number[] {
    let i: number = 0;

    while (true) {
        const inst: number = intcodes[i];
        let op: number;
        let modes: boolean[];
        [op, modes] = opMode(inst);

        if (99 === op) {
            console.log('halting');
            break;
        }
        else if (1 === op) { // addition
            const a: number = lookup(intcodes, i + 1, modes[0]);
            const b: number = lookup(intcodes, i + 2, modes[1]);
            intcodes[intcodes[i + 3]] = a + b;
            i += 4;
        }
        else if (2 === op) { // multiplication
            const a: number = lookup(intcodes, i + 1, modes[0]);
            const b: number = lookup(intcodes, i + 2, modes[1]);
            intcodes[intcodes[i + 3]] = a * b;
            i += 4;
        }
        else if (3 === op) {
            intcodes[intcodes[i + 1]] = input;
            i += 2;
        }
        else if (4 === op) {
            console.log('output', intcodes[intcodes[i + 1]]);
            i += 2;
        }
        else if (5 === op) {
            const a: number = lookup(intcodes, i + 1, modes[0]);
            const b: number = lookup(intcodes, i + 2, modes[1]);
            if (a !== 0) {
                i = b;
            } else {
                i += 3;
            }
        }
        else if (6 === op) {
            const a: number = lookup(intcodes, i + 1, modes[0]);
            const b: number = lookup(intcodes, i + 2, modes[1]);
            if (a === 0) {
                i = b;
            } else {
                i += 3;
            }
        }
        else if (7 === op) {
            const a: number = lookup(intcodes, i + 1, modes[0]);
            const b: number = lookup(intcodes, i + 2, modes[1]);
            const c: number = (a < b) ? 1 : 0;
            intcodes[intcodes[i + 3]] = c;
            i += 4;
        }
        else if (8 === op) {
            const a: number = lookup(intcodes, i + 1, modes[0]);
            const b: number = lookup(intcodes, i + 2, modes[1]);
            const c: number = (a === b) ? 1 : 0;
            intcodes[intcodes[i + 3]] = c;
            i += 4;
        }
        else {
            console.log('illegal op code', op);
            break;
        }
    }

    return intcodes;
}

function part1() {
    console.log('Part 1');
    const lines: number[] = readInputSync('input.txt').map(x => Number(x));
    prog(lines, 1);
}

function part2() {
    console.log('Part 2');
    const lines: number[] = readInputSync('input.txt').map(x => Number(x));
    prog(lines, 5);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
