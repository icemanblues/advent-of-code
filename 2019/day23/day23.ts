import * as fs from 'fs';
import { Amp, prog, progAmp } from '../intcode';

const dayNum: string = "23";
const dayTitle: string = "Category Six";

function readInput(filename: string): number[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/,/).map(Number);
}

class Network {
    comps: Amp[];
    msgIn: Map<string, number[]>;
    msgOut: Map<string, number[]>;


    constructor(n: number) {
        this.msgIn = new Map();
        this.msgOut = new Map();

        this.comps = new Array<Amp>(n);
        for (let i = 0; i < n; i++) {
            const intcode = readInput('input.txt');
            const input: number[] = [i];
            const output: number[] = [];
            const amp = new Amp(String(i), intcode, input, output);
            this.comps[i] = (amp);

            this.msgIn.set(amp.name, input);
            this.msgOut.set(amp.name, output);

            amp.inputCallback = () => {
                const inner = this.msgIn.get(amp.name);
                let v = -1;
                if (inner.length > 0) {
                    v = inner.shift();
                }
                return v;
            };

            amp.outputCallback = (n) => {
                const outer = this.msgOut.get(amp.name);
                outer.push(n);

                while (outer.length >= 3) {
                    const addr = outer.shift();
                    const x = outer.shift();
                    const y = outer.shift();

                    if (addr === 255) {
                        console.log(`addr ${addr} x ${x} y ${y}`);
                        // do I set all of them to halted?
                        this.comps.forEach(a => a.isHalted = true);
                        return;
                    }

                    const rcvr = this.msgIn.get(String(addr));
                    rcvr.push(x);
                    rcvr.push(y);
                }
            };
        }
    }

    run() {
        while (!this.comps.reduce((acc, curr) => acc && curr.isHalted, true)) {
            for (let i = 0; i < this.comps.length; i++) {
                progAmp(this.comps[i]);
            }
        }
    }
}

function part1() {
    console.log('Part 1');
    const numComps = 50;
    const network = new Network(numComps);

    network.run(); // 19937
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
