import { Amp, prog, parseIntcode } from '../intcode';

const dayNum: string = "19";
const dayTitle: string = "Tractor Beam";

function checkPoint(x: number, y: number): boolean {
    const input: number[] = [];
    const output: number[] = [];
    const intcode = parseIntcode('input.txt');
    input.push(x);
    input.push(y);
    const amp = new Amp('Part 1', intcode, input, output);
    prog(amp);
    return output.shift() === 1;
}

function displayMap(width: number, height: number) {
    for (let y = 0; y < height; y++) {
        const line: string[] = [];
        for (let x = 0; x < width; x++) {
            if (checkPoint(x, y)) {
                line.push('#');
            } else {
                line.push('.');
            }
        }
        console.log(line.join(''));
    }
}

function count(width: number, height: number): number {
    let count = 0;
    for (let x = 0; x < width; x++) {
        for (let y = 0; y < height; y++) {
            if (checkPoint(x, y)) {
                count++
            }
        }
    }
    return count;
}

function santaFits(x: number, y: number): boolean {
    return checkPoint(x, y) && checkPoint(x + 99, y) && checkPoint(x, y + 99);
}

function santaSearch(): number {
    let [x, y] = [4, 6];

    while (!santaFits(x, y)) {
        y++;
        while (!checkPoint(x, y)) {
            x++;
        }
        //(x,y) is at the start on the new line
        // lets see how wide this line is
        let i = x;
        while (checkPoint(i, y)) {
            i++;
        }

        // check if santa fits
        if (santaFits(i - 100, y)) {
            return solve(i - 100, y);
        }
    }

    return -1;
}

function solve(x: number, y: number): number {
    return x * 10000 + y;
}

function part1() {
    console.log('Part 1', count(50, 50));
}

function part2() {
    //displayMap(50, 50);
    console.log('Part 2', santaSearch());
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
