import fs from 'fs';
import { Amp, prog } from '../intcode';

const dayNum: string = "05";
const dayTitle: string = "Sunny with a Chance of Asteroids";

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
    console.log(output);
}

function part2() {
    console.log('Part 2');
    const lines: number[] = readInputSync('input.txt');
    const output: number[] = [];
    const amp: Amp = new Amp('part1', lines, [5], output);
    prog(amp);
    console.log(output);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
