import fs from 'fs';

import { Amp, prog, progAmp, parseIntcode } from '../intcode';
import { strt, str, tuple } from '../util';

const dayNum: string = "17";
const dayTitle: string = "Set and Forget";

const ascii: Map<number, string> = new Map([
    [35, '#'],
    [46, '.'],
    [10, '\n']
]);

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
    let x = 0;
    let y = 0;

    const map: string[][] = [];
    let row: string[] = [];
    while (arr.length !== 0) {
        const n = arr.shift();
        const tile = ascii.get(n);
        switch (tile) {
            case '#':
                row.push(tile);
                x++;
                break;
            case '.':
                row.push(tile);
                x++;
                break;
            case '\n':
                map.push(row);
                y++;
                x = 0;
                row = [];
                break;
            default:
                row.push(String.fromCharCode(n));
                x++;
                break;
        }
    }

    return map;
}

function displayGrid(grid: string[][]): void {
    for (let y = 0; y < grid.length; y++) {
        console.log(grid[y].join(''));
    }
}

const MEM_LIMIT = 20;

const grammar = new Map<string, number>([
    [',', 44],
    ['\n', 10]
]);

const COMMA = 44;
const routines = new Map<string, number>([
    ['A', 65],
    ['B', 66],
    ['C', 66]
]);

const movement = new Map<string, number>([
    ['L', 76],
    ['R', 82]
]);


function go(amp: Amp) {
    while (!amp.isHalted) {
        progAmp(amp);
    }
}

function part1() {
    console.log('Part 1');
    const intcode = parseIntcode('input.txt');
    const input: number[] = [];
    const output: number[] = [];
    const amp = new Amp('Aft', intcode, input, output);
    prog(amp);
    const grid = buildMap(output);
    console.log(intersection(grid));
}

function part2() {
    console.log('Part 2');
    const intcode = parseIntcode('input.txt');
    intcode[0] = 2;


    const main = [65, 44, 66, 44, 67, 44, 66, 44, 65, 44, 67, 10];
    const amove = [82, 44, 56, 44, 82, 44, 56, 10];
    const bmove = [82, 44, 52, 44, 82, 44, 52, 44, 82, 44, 56, 10];
    const cmove = [76, 44, 54, 44, 76, 44, 50, 10];
    const video = [89, 10]; // y
    //const video = [78, 10]; // no
    const input: number[] = main.concat(amove, bmove, cmove, video);
    const output: number[] = [];
    const amp = new Amp('Aft', intcode, input, output);
    prog(amp);

    const grid = buildMap(output);
    displayGrid(grid);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    //part1();
    part2();
}

main();
