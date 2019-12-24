import { Amp, parseIntcode, prog } from '../intcode';

const dayNum: string = "02";
const dayTitle: string = "1202 Program Alarm";

function nounVerb(noun: number, verb: number): number {
    return 100 * noun + verb;
}

function progger(n: number, v: number): number {
    const intcodes: number[] = parseIntcode('input.txt');
    intcodes[1] = n;
    intcodes[2] = v;
    const amp: Amp = new Amp('day02', intcodes, [], []);
    prog(amp);
    return intcodes[0];
}

function part1() {
    console.log('Part 1', progger(12, 2));
}

function part2() {
    const target: number = 19690720;
    let answer: number = -1;
    iter:
    for (let n: number = 0; n <= 99; n++) {
        for (let v: number = 0; v <= 99; v++) {
            if (progger(n, v) === target) {
                answer = nounVerb(n, v);
                break iter;
            }
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
