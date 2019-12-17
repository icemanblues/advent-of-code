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
                // now we need to check around it for scaffold
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
    for (let n of arr) {
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
        }
    }

    return map;
}

const MEM_LIMIT = 20;

const movement = new Map<string, number>();

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
    const input: number[] = [];
    const output: number[] = [];
    const amp = new Amp('Aft', intcode, input, output);
    prog(amp);

}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
