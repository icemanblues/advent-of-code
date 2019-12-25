import { Amp, prog, parseIntcode } from '../intcode';
import { toAscii } from '../util';

const dayNum: string = "17";
const dayTitle: string = "Set and Forget";

function intersection(grid: string[][]): number {
    let sum = 0;
    for (let y = 0; y < grid.length; y++) {
        for (let x = 0; x < grid[y].length; x++) {
            if (grid[y][x] === '#') {
                let b: boolean = true;
                if (x > 0) {
                    b = b && grid[y][x - 1] === '#';
                }
                if (y > 0) {
                    b = b && grid[y - 1][x] === '#';
                }
                if (x < grid[y].length - 1) {
                    b = b && grid[y][x + 1] === '#';
                }
                if (y < grid.length - 1) {
                    b = b && grid[y + 1][x] === '#';
                }

                if (b) {
                    sum += x * y;
                }
            }
        }
    }
    return sum;
}

function buildMap(arr: number[]): string[][] {
    const map: string[][] = [];
    let row: string[] = [];
    while (arr.length !== 0) {
        const n = arr.shift();
        const tile = String.fromCharCode(n);
        switch (tile) {
            case '\n':
                map.push(row);
                row = [];
                break;
            default:
                row.push(tile);
                break;
        }
    }

    return map;
}

function part1() {
    const intcode = parseIntcode('input.txt');
    const input: number[] = [];
    const output: number[] = [];
    const amp = new Amp('Aft', intcode, input, output);
    prog(amp);
    const grid = buildMap(output);
    console.log('Part 1', intersection(grid));
}

function part2() {
    const intcode = parseIntcode('input.txt');
    intcode[0] = 2;

    // solved this by hand (emacs), not using a zip hoffman code compression algo
    const main = toAscii('A,B,A,C,A,B,C,B,C,B');
    const amove = toAscii('L,10,R,8,L,6,R,6');
    const bmove = toAscii('L,8,L,8,R,8');
    const cmove = toAscii('R,8,L,6,L,10,L,10');
    const video = toAscii('n');

    const input: number[] = main.concat(amove, bmove, cmove, video);
    const output: number[] = [];
    const amp = new Amp('Aft', intcode, input, output);
    prog(amp);
    console.log('Part 2', output[output.length - 1]);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
