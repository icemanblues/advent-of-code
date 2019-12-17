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

function dotProduct(input: number[], pattern: number[], offset: number = 0): number {
    let sum: number = 0;
    for (let i = 0; i < input.length; i++) {
        sum += input[i] * pattern[(i + 1 + offset) % pattern.length];
    }
    return Math.abs(sum) % 10;
};

function computePattern(base: number[], position: number): number[] {
    const r: number[] = [];


    for (let b of BASE) {
        for (let i = 0; i <= position; i++) {
            r.push(b);
        }
    }

    return r;
}

function fft(input: number[], phases: number, offset: number = 0): string {
    let result: number[] = input;
    for (let i = 1; i <= phases; i++) {
        let temp: number[] = [];
        for (let j = 0; j < result.length; j++) {
            const pattern = computePattern(BASE, j);
            const dp = dotProduct(result, pattern, offset);
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
    const input = readInputSync('input.txt');
    const offset = Number(input.slice(0, 7).join(''));

    let big: number[] = [];
    for (let i = 0; i < 10000; i++) {
        big = big.concat(input);
    }
    big = big.slice(offset);
    const output = fft(big, 100)
    console.log('Part 1', output.substring(0, 8));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    //part2();
}

main();
