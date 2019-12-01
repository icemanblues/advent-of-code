import fs from 'fs';

const dayNum: string = "01";
const dayTitle: string = "The Tyranny of the Rocket Equation";

// reads entire file into memory, then splits it
function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/\r?\n/);
    return lines;
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
    console.log('Part 1');
    console.log(fuelReq(12));
    console.log(fuelReq(14));
    console.log(fuelReq(1969));
    console.log(fuelReq(100756));

    const inputs: string[] = readInputSync('input.txt');
    const p1: number = inputs.reduce((acc, curr) => acc + fuelReq(Number(curr)), 0);
    console.log(`part1 ${p1}`);
}

function part2() {
    console.log('Part 2');
    console.log(fuelFuel(12));
    console.log(fuelFuel(14));
    console.log(fuelFuel(1969));
    console.log(fuelFuel(100756));

    const inputs: string[] = readInputSync('input.txt');
    const p2: number = inputs.reduce((acc, curr) => acc + fuelFuel(Number(curr)), 0);
    console.log(`part2 ${p2}`);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
