import * as fs from 'fs';

const dayNum: string = "12";
const dayTitle: string = "The N-Body Problem";

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

class Point3D {
    x: number;
    y: number;
    z: number;

    constructor(x: number, y: number, z: number) {
        this.x = x;
        this.y = y;
        this.z = z;
    }

    energy(): number {
        return Math.abs(this.x) + Math.abs(this.y) + Math.abs(this.z);
    }
}

class Moon {
    pos: Point3D;
    vel: Point3D;

    constructor(p: Point3D) {
        this.pos = p;
        this.vel = new Point3D(0, 0, 0);
    }

    energy(): number {
        return this.pos.energy() * this.vel.energy();
    }
}

function parse(filename: string): Moon[] {
    const moons: Moon[] = [];

    const lines: string[] = readInputSync(filename);
    lines.forEach(l => {
        const p: Point3D = new Point3D(0, 0, 0);
        l.slice(1, l.length - 1)
            .split(/,/).forEach(w => {
                const v: string[] = w.split(/=/);
                switch (v[0].trim()) {
                    case 'x':
                        p.x = Number(v[1].trim());
                        break;
                    case 'y':
                        p.y = Number(v[1].trim());
                        break;
                    case 'z':
                        p.z = Number(v[1].trim());
                        break;
                }
            });
        moons.push(new Moon(p));
    });

    return moons;
}

function gravity(m1: Moon, m2: Moon) {
    if (m1.pos.x > m2.pos.x) {
        m1.vel.x -= 1;
        m2.vel.x += 1;
    } else if (m1.pos.x < m2.pos.x) {
        m1.vel.x += 1;
        m2.vel.x -= 1;
    }

    if (m1.pos.y > m2.pos.y) {
        m1.vel.y -= 1;
        m2.vel.y += 1;
    } else if (m1.pos.y < m2.pos.y) {
        m1.vel.y += 1;
        m2.vel.y -= 1;
    }

    if (m1.pos.z > m2.pos.z) {
        m1.vel.z -= 1;
        m2.vel.z += 1;
    } else if (m1.pos.z < m2.pos.z) {
        m1.vel.z += 1;
        m2.vel.z -= 1;
    }
}

function velocity(m: Moon) {
    m.pos.x += m.vel.x;
    m.pos.y += m.vel.y;
    m.pos.z += m.vel.z;
}

function step(moons: Moon[]) {
    for (let i: number = 0; i < moons.length - 1; i++) {
        for (let j: number = i + 1; j < moons.length; j++) {
            gravity(moons[i], moons[j]);
        }
    }

    moons.forEach(velocity);
}
function motion(moons: Moon[], steps: number): number {
    for (let i: number = 0; i < steps; i++) {
        step(moons);
    }

    return moons.reduce((acc, m) => acc + m.energy(), 0);
}

function loopDetection(moons: Moon[]): number {
    const startX: number[] = [];
    const startY: number[] = [];
    const startZ: number[] = [];

    moons.forEach(m => {
        startX.push(m.pos.x);
        startY.push(m.pos.y);
        startZ.push(m.pos.z);
    });

    function atOrig(start: number[], accessor: (m: Moon) => number): boolean {
        for (let i = 0; i < start.length; i++) {
            if (start[i] !== accessor(moons[i])) {
                return false;
            }
        }
        return true;
    }

    const intervals: number[] = [0, 0, 0];
    let steps: number = 1;
    while (intervals.indexOf(0) !== -1) { // loop until there are no zeroes
        step(moons);
        steps++;

        if (intervals[0] === 0 && atOrig(startX, (m: Moon) => m.pos.x)) {
            intervals[0] = steps;
        }
        if (intervals[1] === 0 && atOrig(startY, (m: Moon) => m.pos.y)) {
            intervals[1] = steps;
        }
        if (intervals[2] === 0 && atOrig(startZ, (m: Moon) => m.pos.z)) {
            intervals[2] = steps;
        }
    }

    return computeLCM(intervals);
}

function computeLCM(nums: number[]): number {
    function gcd(a: number, b: number): number {
        return !b ? a : gcd(b, a % b);
    }

    function lcm(a: number, b: number): number {
        return (a * b) / gcd(a, b);
    }

    let n: number = nums[0];
    return nums.reduce((acc, curr) => lcm(acc, curr), n);
}

function part1() {
    const moons: Moon[] = parse('input.txt');
    console.log('Part 1', motion(moons, 1000));
}

function part2() {
    const moons: Moon[] = parse('input.txt');
    console.log('Part 2', loopDetection(moons));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
