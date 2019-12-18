import fs from 'fs';
import { Amp, progAmp } from '../intcode';

const dayNum: string = "07";
const dayTitle: string = "Amplification Circuit";

function readInputSync(filename: string): number[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight()
        .split(/,/)
        .map(Number);
}

function input(): number[] {
    return readInputSync('input.txt');
}

function maxPhaseSetting(amp: number, ampIn: number, usedPhases: number[]): number {
    if (amp == 5) {
        return ampIn;
    }

    let max: number = -1;
    for (let phase: number = 0; phase < 5; phase++) {
        if (usedPhases.indexOf(phase) !== -1) {
            continue;
        }

        usedPhases.push(phase);
        const intcodes: number[] = input();
        const inputs: number[] = [phase, ampIn];
        const outputs: number[] = [];
        const a: Amp = new Amp(`${amp}`, intcodes, inputs, outputs);
        progAmp(a);
        const ampOut: number = outputs[0];
        const thrust: number = maxPhaseSetting(amp + 1, ampOut, usedPhases);
        if (thrust > max) {
            max = thrust;
        }
        usedPhases.pop();
    }

    return max;
}

function feedback(phases: number[]): number {
    const inputA: number[] = [phases[0], 0];
    const inputB: number[] = [phases[1]];
    const inputC: number[] = [phases[2]];
    const inputD: number[] = [phases[3]];
    const inputE: number[] = [phases[4]];

    const ampA: Amp = new Amp('A', input(), inputA, inputB);
    const ampB: Amp = new Amp('B', input(), inputB, inputC);
    const ampC: Amp = new Amp('C', input(), inputC, inputD);
    const ampD: Amp = new Amp('D', input(), inputD, inputE);
    const ampE: Amp = new Amp('E', input(), inputE, inputA);

    const amps: Amp[] = [ampA, ampB, ampC, ampD, ampE];

    let i: number = 0;
    while (amps.reduce((acc, a) => acc && !a.isHalted, true)) {
        const a: Amp = amps[i % 5];
        progAmp(a);
        i++;
    }

    return inputA[0];
}

function maxFeedback(phases: number[]): number {
    if (phases.length == 5) {
        return feedback(phases);
    }

    let max: number = -1;
    for (let phase: number = 5; phase < 10; phase++) {
        if (phases.indexOf(phase) !== -1) {
            continue;
        }

        phases.push(phase);
        const t: number = maxFeedback(phases);
        if (t > max) {
            max = t;
        }
        phases.pop();
    }

    return max;
}

function part1() {
    console.log('Part 1');
    console.log(maxPhaseSetting(0, 0, []));
}

function part2() {
    console.log('Part 2');
    console.log(maxFeedback([]));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
