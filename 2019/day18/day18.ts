import fs from 'fs';

const dayNum: string = "18";
const dayTitle: string = "Many-Worlds Interpretation";

function read(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

const input = read('input.txt');

function print(board: string[]): void {
    for (let l of board) {
        console.log(l);
    }
}

function find(board: string[], s: string): [number, number] {
    for (let y = 0; y < board.length; y++) {
        const x = board[y].indexOf(s);
        if (x !== -1) {
            return [x, y];
        }
    }

    return [-1, -1];
}


function part1() {
    console.log('Part 1');
    print(input);
    const start = find(input, '@');
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
