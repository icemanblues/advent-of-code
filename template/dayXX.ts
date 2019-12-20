import * as fs from 'fs';

const dayNum: string = "XX";
const dayTitle: string = "Title";

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight().split(/\r?\n/);
}

function part1() {
    console.log('Part 1');
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
