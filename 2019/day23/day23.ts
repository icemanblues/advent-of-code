import { Amp, progAmp, parseIntcode } from '../intcode';

const dayNum: string = "23";
const dayTitle: string = "Category Six";

class Network {
    private comps: Amp[];
    private msgIn: Map<string, number[]>;
    private msgOut: Map<string, number[]>;

    private result: number = -1;

    private useNat: boolean = false;
    private natx: number = -1;
    private naty: number = -1;
    private natyPrev: number = -2;
    private idleCount: Set<string> = new Set();


    constructor(n: number) {
        this.msgIn = new Map();
        this.msgOut = new Map();

        this.comps = new Array<Amp>(n);
        for (let i = 0; i < n; i++) {
            const intcode = parseIntcode('input.txt');
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
                    this.idleCount.delete(amp.name);
                } else {
                    if (this.useNat && this.isIdle()) {
                        this.idleCount.add(amp.name);
                        if (this.idleCount.size === this.comps.length) {
                            if (this.naty === this.natyPrev && this.naty !== -1) {
                                this.halt(this.naty);
                            }

                            this.natyPrev = this.naty;
                            const inZero = this.msgIn.get(String(0));
                            inZero.push(this.natx);
                            inZero.push(this.naty);
                        }
                    }
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
                        if (this.useNat) {
                            this.natx = x;
                            this.naty = y;
                        } else {
                            this.halt(y);
                        }
                    } else {
                        const rcvr = this.msgIn.get(String(addr));
                        rcvr.push(x);
                        rcvr.push(y);
                    }
                }
            };
        }
    }

    run(enableNat: boolean = false): number {
        this.useNat = enableNat;
        while (!this.comps.reduce((acc, curr) => acc && curr.isHalted, true)) {
            for (let i = 0; i < this.comps.length; i++) {
                progAmp(this.comps[i], true, true);
            }
        }
        return this.result;
    }

    private isIdle(): boolean {
        let count = 0;
        this.msgIn.forEach((inputs, cname) => count += inputs.length);
        return count === 0;
    }

    private halt(n: number) {
        this.result = n;
        this.comps.forEach(a => a.isHalted = true);
    }
}

const NUM_COMPS = 50;

function part1() {
    const network = new Network(NUM_COMPS);
    console.log('Part 1', network.run()); // 19937
}

function part2() {
    const network = new Network(NUM_COMPS);
    console.log('Part 2', network.run(true)); // 13758
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
