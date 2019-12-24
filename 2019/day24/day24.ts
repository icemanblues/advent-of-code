import * as fs from 'fs';
import { str, strt, tuple } from '../util';

const dayNum: string = "24";
const dayTitle: string = "Planet of Discord";

function readInput(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

function printGrid(grid: Set<string>): void {
    for (let y = 0; y < 5; y++) {
        const line: string[] = [];
        for (let x = 0; x < 5; x++) {
            if (grid.has(str(x, y))) {
                line.push('#');
            } else {
                line.push('.');
            }
        }
        console.log(line.join(''));
    }
}

function makeGrid(lines: string[]): Set<string> {
    const s: Set<string> = new Set();
    for (let y = 0; y < lines.length; y++) {
        for (let x = 0; x < lines[y].length; x++) {
            if (lines[y].charAt(x) === '#') {
                s.add(str(x, y));
            }
        }
    }
    return s;
}

function nextMinute(grid: Set<string>): Set<string> {
    const s: Set<string> = new Set();
    for (let x = 0; x < 5; x++) {
        for (let y = 0; y < 5; y++) {
            let bugCount: number = 0;
            [[0, 1], [0, -1], [1, 0], [-1, 0]].forEach(t => {
                let [xi, yi] = [x + t[0], y + t[1]];
                if (grid.has(str(xi, yi))) {
                    bugCount++;
                }
            });

            if (grid.has(str(x, y))) { // has a bug
                if (bugCount === 1) {
                    s.add(str(x, y));
                }
            } else { // is empty, no bug
                if (bugCount === 1 || bugCount === 2) {
                    s.add(str(x, y));
                }
            }
        }
    }
    return s;
}

function bioScore(grid: Set<string>): number {
    let score: number = 0;
    grid.forEach(s => {
        const [x, y] = tuple(s);
        const idx = y * 5 + x;
        score += Math.pow(2, idx);
    });
    return score;
}

function findRepeat(grid: Set<string>): number {
    const bio: Set<number> = new Set();
    let score = -1;
    while (!bio.has(score)) {
        bio.add(score);
        score = bioScore(grid);
        grid = nextMinute(grid);
    }
    return score;
}

function part1() {
    console.log('Part 1', findRepeat(makeGrid(readInput('input.txt'))));
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
