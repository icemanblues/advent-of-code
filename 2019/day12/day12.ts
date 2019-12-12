import fs from 'fs';

const dayNum: string = "12";
const dayTitle: string = "The N-Body Problem";

function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/\r?\n/);
    return lines;
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

    toString(): string {
        return `${this.pos.x},${this.pos.y},${this.pos.z}|${this.vel.x},${this.vel.y},${this.vel.z}`;
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

function moonsKey(moons: Moon[]): string {
    return moons.map(m => m.toString()).join('#');
}

function loopDetection(moons: Moon[]): number {
    const start: string = moonsKey(moons);
    //const s: Set<string> = new Set<string>();

    let steps: number = 1;
    step(moons);
    //while (!s.has(moonsKey(moons))) {
    while (start !== moonsKey(moons)) {
        //s.add(moonsKey(moons));
        step(moons);
        steps++;

        if (steps % 10000 === 0) {
            console.log(steps);
        }
    }

    return steps;
}

function part1() {
    console.log('Part 1');

    const moons: Moon[] = parse('input.txt');
    console.log(motion(moons, 1000));
}

function part2() {
    console.log('Part 2');

    //const moons: Moon[] = parse('test-100-1940.txt');
    const moons: Moon[] = parse('input.txt');
    console.log(loopDetection(moons));

}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    //part1();
    part2();
}

main();
