import * as fs from 'fs';

const dayNum: string = "01";
const dayTitle: string = "The Tyranny of the Rocket Equation";

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

function fuelReq(mass: number): number {
    return Math.floor(mass / 3) - 2;
}

function fuelFuel(mass: number): number {
    let fuel: number = 0;
    while (mass > 0) {
        let f: number = fuelReq(mass);
        if (f > 0) {
            fuel += f;
            mass = f;
        } else {
            break;
        }
    }
    return fuel;
}

function part1() {
    const inputs: string[] = readInputSync('input.txt');
    const p1: number = inputs.reduce((acc, curr) => acc + fuelReq(Number(curr)), 0);
    console.log(`Part 1 ${p1}`);
}

function part2() {
    const inputs: string[] = readInputSync('input.txt');
    const p2: number = inputs.reduce((acc, curr) => acc + fuelFuel(Number(curr)), 0);
    console.log(`Part 2 ${p2}`);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
