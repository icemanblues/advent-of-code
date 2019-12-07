import fs from 'fs';

const dayNum: string = "07";
const dayTitle: string = "Amplification Circuit";

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

class Amp {
    name: string;
    isHalted: boolean;
    intcodes: number[];
    i: number;
    inputs: number[];
    outputs: number[];

    constructor(name: string, ic: number[], ins: number[], outs: number[]) {
        this.name = name;
        this.intcodes = ic;
        this.inputs = ins;
        this.outputs = outs;

        this.isHalted = false;
        this.i = 0;
    }

    setValue(i: number, v: number) {
        this.intcodes[this.intcodes[i]] = v;
    }
}

function progAmp(amp: Amp) {
    if (amp.isHalted) {
        console.log('Amp is halted.', amp.name);
        return;
    }

    while (true) {
        const inst: number = amp.intcodes[amp.i];
        let op: number;
        let modes: boolean[];
        [op, modes] = opMode(inst);

        if (99 === op) {
            amp.isHalted = true;
            break;
        }
        else if (1 === op) { // addition
            const a: number = lookup(amp.intcodes, amp.i + 1, modes[0]);
            const b: number = lookup(amp.intcodes, amp.i + 2, modes[1]);
            amp.setValue(amp.i + 3, a + b);
            amp.i += 4;
        }
        else if (2 === op) { // multiplication
            const a: number = lookup(amp.intcodes, amp.i + 1, modes[0]);
            const b: number = lookup(amp.intcodes, amp.i + 2, modes[1]);
            amp.setValue(amp.i + 3, a * b);
            amp.i += 4;
        }
        else if (3 === op) { // input
            const input: number = amp.inputs.shift();
            if (!input) {
                console.log('ut-oh. no input provided for amp', amp.name, input);
            }

            amp.setValue(amp.i + 1, input);
            amp.i += 2;
        }
        else if (4 === op) { //output
            const output: number = amp.intcodes[amp.intcodes[amp.i + 1]];
            amp.outputs.push(output);
            amp.i += 2;
        }
        else if (5 === op) { // jump if not zero
            const a: number = lookup(amp.intcodes, amp.i + 1, modes[0]);
            const b: number = lookup(amp.intcodes, amp.i + 2, modes[1]);
            if (a !== 0) {
                amp.i = b;
            } else {
                amp.i += 3;
            }
        }
        else if (6 === op) { // jump if zero
            const a: number = lookup(amp.intcodes, amp.i + 1, modes[0]);
            const b: number = lookup(amp.intcodes, amp.i + 2, modes[1]);
            if (a === 0) {
                amp.i = b;
            } else {
                amp.i += 3;
            }
        }
        else if (7 === op) {
            const a: number = lookup(amp.intcodes, amp.i + 1, modes[0]);
            const b: number = lookup(amp.intcodes, amp.i + 2, modes[1]);
            const c: number = (a < b) ? 1 : 0;
            amp.setValue(amp.i + 3, c);
            amp.i += 4;
        }
        else if (8 === op) {
            const a: number = lookup(amp.intcodes, amp.i + 1, modes[0]);
            const b: number = lookup(amp.intcodes, amp.i + 2, modes[1]);
            const c: number = (a === b) ? 1 : 0;
            amp.setValue(amp.i + 3, c);
            amp.i += 4;
        }
        else {
            console.log('illegal op code', op);
            break;
        }
    }
}

// returns the most recent output
function prog(intcodes: number[], inputs: number[], outputs: number[]): number {
    let output: number = 0;

    let i: number = 0;
    while (true) {
        const inst: number = intcodes[i];
        let op: number;
        let modes: boolean[];
        [op, modes] = opMode(inst);

        if (99 === op) {
            //console.log('halting', output);
            return output;
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
            const input: number = inputs.shift();
            //console.log('input', input);
            intcodes[intcodes[i + 1]] = input;
            i += 2;
        }
        else if (4 === op) {
            output = intcodes[intcodes[i + 1]];
            outputs.push(output);
            //console.log('output', output);
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

    console.log('never should have reached here', output);
    return -1;
}

function input(): number[] {
    return readInputSync('input.txt').map(Number);
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

        //const ampOut: number = prog(intcodes, inputs, outputs);
        const thrust: number = maxPhaseSetting(amp + 1, ampOut, usedPhases);
        if (thrust > max) {
            max = thrust;
        }
        usedPhases.pop();
    }

    return max;
}


function feedback(): number {
    const phases: number[] = [9, 8, 7, 6, 5];

    const inputA: number[] = [phases[0], 0];
    const inputB: number[] = [phases[1]];
    const inputC: number[] = [phases[2]];
    const inputD: number[] = [phases[3]];
    const inputE: number[] = [phases[4]];



    return -1;
}

function part1() {
    console.log('Part 1');
    console.log(maxPhaseSetting(0, 0, []));
}

function part2() {
    console.log('Part 2');
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
