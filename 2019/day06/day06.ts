import fs from 'fs';

const dayNum: string = "XX";
const dayTitle: string = "Title";

// reads entire file into memory, then splits it
function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/\r?\n/);
    return lines;
}

function part1() {
    console.log('Part 1');

    const lines: string[] = readInputSync('input.txt');
    console.log(lines);
}

function part2() {
    console.log('Part 2');

    const lines: string[] = readInputSync('input.txt');
    console.log(lines);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
