//
//
//

import fs from 'fs';
import readline from 'readline';

const dayNum: string = "XX";
const dayTitle: string = "Title";
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
    });

    rl.on('line', (line: string) => lines.push(line));
    return new Promise<string[]>((resolve) => {
        rl.on('close', () => resolve(lines));
    });
}

async function part1() {
    console.log(`Part 1: ${subTitle1}`);

    const lines: string[] = await readInput('input.txt');
    console.log(lines);
}

async function part2() {
    console.log(`Part 2: ${subTitle2}`);

    const lines: string[] = readInputSync('input.txt');
    console.log(lines);
}

async function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    await part1();
    await part2();
}

main();
