import fs from 'fs';

const dayNum: string = "16";
const dayTitle: string = "Flawed Frequency Transmission";

function readInputSync(filename: string): number[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight()
        .split('')
        .map(Number);
}

const BASE: number[] = [0, 1, 0, -1];

function dotProduct(input: number[], pattern: number[]): number {
    let sum: number = 0;
    for (let i = 0; i < input.length; i++) {
        sum += input[i] * pattern[i];
    }
    return Math.abs(sum) % 10;
};

function computePattern(base: number[], position: number, len: number): number[] {
    const r: number[] = [];

    let i = 0;
    while (r.length <= len) {
        for (let j = 0; j <= position; j++) {
            r.push(base[i % base.length]);
        }
        i++;
    }


    r.shift();
    return r;
}

function fft(input: number[], phases: number): string {
    let result: number[] = input;
    for (let i = 1; i <= phases; i++) {
        let temp: number[] = [];
        for (let j = 0; j < result.length; j++) {
            const pattern = computePattern(BASE, j, result.length);
            const dp = dotProduct(result, pattern);
            temp.push(dp); // temp[j] =
        }
        result = temp;
    }

    //return result;
    return result.join('');
}

function part1() {
    console.log('Part 1');
    const input = readInputSync('input.txt');
    const output = fft(input, 100)

    console.log('Part 1', output.substring(0, 8));
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
