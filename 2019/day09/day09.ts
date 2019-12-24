import { Amp, prog, parseIntcode } from '../intcode';

const dayNum: string = "09";
const dayTitle: string = "Sensor Boost";

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
    const amp: Amp = new Amp('part1', intcode, [2], output);
    prog(amp);
    console.log('Part 2', output);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
