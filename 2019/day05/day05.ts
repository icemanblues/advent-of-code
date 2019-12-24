import { Amp, prog, parseIntcode } from '../intcode';

const dayNum: string = "05";
const dayTitle: string = "Sunny with a Chance of Asteroids";

function part1() {
    const intcode: number[] = parseIntcode('input.txt');
    const output: number[] = [];
    const amp: Amp = new Amp('part1', intcode, [1], output);
    prog(amp);
    console.log('Part 1', output);
}

function part2() {
    const intcode: number[] = parseIntcode('input.txt');
    const output: number[] = [];
    const amp: Amp = new Amp('part1', intcode, [5], output);
    prog(amp);
    console.log('Part 2', output);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
