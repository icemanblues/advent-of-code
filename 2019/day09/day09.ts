import fs from 'fs';
import { Amp, prog } from '../intcode';

const dayNum: string = "09";
const dayTitle: string = "Sensor Boost";

function readInputSync(filename: string): number[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight()
        .split(/,/)
        .map(Number);
}

function part1() {
    console.log('Part 1');
    const lines: number[] = readInputSync('input.txt');
    const output: number[] = [];
    const amp: Amp = new Amp('part1', lines, [1], output);
    prog(amp);
    console.log('part1', output);
}

function part2() {
    console.log('Part 2');
    const lines: number[] = readInputSync('input.txt');
    const output: number[] = [];
    const amp: Amp = new Amp('part1', lines, [2], output);
    prog(amp);
    console.log('part1', output);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
