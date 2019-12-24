import { getOrDefault, tuple } from '../util';
import fs from 'fs';

export function parseIntcode(filename: string): number[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight()
        .split(/,/)
        .map(Number);
}

function opMode(inst: number): [number, number[]] {
    const s: string = String(inst);
    const op: number = Number(s.slice(s.length - 2));
    const mode: number[] = [];

    for (let i = s.length - 3; i >= 0; i--) {
        mode.push(Number(s.charAt(i)));
    }
    while (mode.length !== 3) {
        mode.push(0);
    }

    return [op, mode];
}

export class Amp {
    name: string;
    isHalted: boolean;
    intcodes: number[];
    i: number;
    inputs: number[];
    outputs: number[];
    mem: Map<number, number>;
    relBase: number;

    inputCallback: () => number;
    outputCallback: (n: number) => void;

    constructor(name: string, ic: number[], ins: number[], outs: number[]) {
        this.name = name;
        this.intcodes = ic;
        this.inputs = ins;
        this.outputs = outs;

        this.isHalted = false;
        this.mem = new Map<number, number>();
        this.i = 0;
        this.relBase = 0;
    }

    _idx(k: number, mode: number): number {
        let idx: number = 0;

        if (mode === 0) { // position
            idx = this.intcodes[k];
        } else if (mode === 1) { // immediate
            idx = k;
        } else if (mode === 2) { // relative
            idx = this.relBase + this.intcodes[k];
        } else {
            console.log('this should never have happened');
        }

        return idx;
    }

    getValue(k: number, mode: number): number {
        const idx: number = this._idx(k, mode);
        if (idx < this.intcodes.length) {
            return this.intcodes[idx];
        }

        return getOrDefault(this.mem, idx, 0);
    }

    setValue(k: number, v: number, mode: number) {
        const idx: number = this._idx(k, mode);

        if (idx < this.intcodes.length) {
            this.intcodes[idx] = v;
        } else {
            this.mem.set(idx, v);
        }
    }
}

// run until it halts, or (outputs, inputs) conditionally
export function progAmp(amp: Amp, isOutput: boolean = true, isInput: boolean = false) {
    if (amp.isHalted) {
        return;
    }

    while (true) {
        const inst: number = amp.intcodes[amp.i];
        let op: number;
        let modes: number[];
        [op, modes] = opMode(inst);

        if (99 === op) {
            amp.isHalted = true;
            break;
        }
        else if (1 === op) { // addition
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = a + b;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (2 === op) { // multiplication
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = a * b;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (3 === op) { // input
            let input: number;
            if (amp.inputCallback) {
                input = amp.inputCallback();
            } else {
                input = amp.inputs.shift();
            }

            amp.setValue(amp.i + 1, input, modes[0]);
            amp.i += 2;

            if (isInput) {
                return;
            }
        }
        else if (4 === op) { //output
            const output: number = amp.getValue(amp.i + 1, modes[0]);
            if (amp.outputCallback) {
                amp.outputCallback(output);
            } else {
                amp.outputs.push(output);
            }
            amp.i += 2;

            if (isOutput) {
                return;
            }
        }
        else if (5 === op) { // jump if not zero
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            if (a !== 0) {
                amp.i = b;
            } else {
                amp.i += 3;
            }
        }
        else if (6 === op) { // jump if zero
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            if (a === 0) {
                amp.i = b;
            } else {
                amp.i += 3;
            }
        }
        else if (7 === op) {
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = (a < b) ? 1 : 0;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (8 === op) {
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            const b: number = amp.getValue(amp.i + 2, modes[1]);
            const c: number = (a === b) ? 1 : 0;
            amp.setValue(amp.i + 3, c, modes[2]);
            amp.i += 4;
        }
        else if (9 === op) {
            const a: number = amp.getValue(amp.i + 1, modes[0]);
            amp.relBase = amp.relBase + a;
            amp.i += 2;
        }
        else {
            console.log('illegal op code', op);
            break;
        }
    }
}

export function prog(amp: Amp) {
    while (!amp.isHalted) {
        progAmp(amp);
    }
}
