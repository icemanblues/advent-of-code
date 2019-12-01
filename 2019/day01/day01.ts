//
//
//

import fs from 'fs';
import readline from 'readline';

const dayNum: string = "01";
const dayTitle: string = "The Tyranny of the Rocket Equation";
const subTitle1: string = "subtitle 1";
const subTitle2: string = "subtitle 2";

// reads entire file into memory, then splits it
function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/\r?\n/);
    return lines;
}

async function readInput(filename: string): Promise<string[]> {
    const lines: string[] = [];
    const fileStream: fs.ReadStream = fs.createReadStream(filename);
    const rl: readline.Interface = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity,
    });12

    rl.on('line', (line: string) => lines.push(line));
    return new Promise<string[]>((resolve) => {
        rl.on('close', () => resolve(lines));
    });
}

function fuelReq(mass: number): number {
    return Math.floor(mass / 3) - 2;
}

function fuelFuel(mass: number): number {
    let fuel: number = 0;
    while(mass > 0) {
        let f: number = fuelReq(mass);
        //console.log(`mass ${mass} with fuel ${f}`);
        if( f > 0 ) {
            fuel += f;
            mass = f;
        } else {
            break;
        }
    }
    return fuel;
}

async function part1() {
    console.log(`Part 1: ${subTitle1}`);
    console.log(fuelReq(12));
    console.log(fuelReq(14));
    console.log(fuelReq(1969));
    console.log(fuelReq(100756));

    const inputs: string[] = await readInput('input.txt');
    const p1: number = inputs.reduce( (acc, curr) => acc+fuelReq(Number(curr)), 0);
    console.log(`part1 ${p1}`);   
}

async function part2() {
    console.log(`Part 2: ${subTitle2}`);
    console.log(fuelFuel(12));
    console.log(fuelFuel(14));
    console.log(fuelFuel(1969));
    console.log(fuelFuel(100756));

    const inputs: string[] = await readInput('input.txt');
    const p2: number = inputs.reduce( (acc, curr) => acc+fuelFuel(Number(curr)), 0);
    console.log(`part2 ${p2}`);   
}

async function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    await part1();
    await part2();
}

main();
