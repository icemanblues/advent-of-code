import * as fs from 'fs';
import { str, strt, tuple } from '../util';

const dayNum: string = "24";
const dayTitle: string = "Planet of Discord";

function readInput(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
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

function bugRules(grid: Set<string>, x: number, y: number, s: Set<string>, bugs: number): void {
    if (grid.has(str(x, y))) { // has a bug
        if (bugs === 1) {
            s.add(str(x, y));
        }
    } else { // is empty, no bug
        if (bugs === 1 || bugs === 2) {
            s.add(str(x, y));
        }
    }
}

function nextMinute(grid: Set<string>): Set<string> {
    const s: Set<string> = new Set();
    for (let x = 0; x < 5; x++) {
        for (let y = 0; y < 5; y++) {
            let bugCount: number = 0;
            DIRECTIONS.forEach(t => {
                let [xi, yi] = [x + t[0], y + t[1]];
                if (grid.has(str(xi, yi))) {
                    bugCount++;
                }
            });

            bugRules(grid, x, y, s, bugCount);
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

const MIDDLE: string = str(2, 2);

const INNER_MAP: Map<string, string[]> = new Map([
    [str(0, 1), [str(0, 0), str(1, 0), str(2, 0), str(3, 0), str(4, 0)]], // down
    [str(0, -1), [str(0, 4), str(1, 4), str(2, 4), str(3, 4), str(4, 4)]], // up
    [str(1, 0), [str(0, 0), str(0, 1), str(0, 2), str(0, 3), str(0, 4)]], // right
    [str(-1, 0), [str(4, 0), str(4, 1), str(4, 2), str(4, 3), str(4, 4)]], // left
]);

function lookInner(t: [number, number], level: number, universe: Map<number, Set<string>>): number {
    let count = 0;
    const grid = universe.get(level + 1);
    if (grid) {
        INNER_MAP.get(strt(t)).forEach(s => {
            if (grid.has(s)) {
                count++;
            }
        });
    }
    return count;
}

const DIRECTIONS: [number, number][] = [[0, 1], [0, -1], [1, 0], [-1, 0]];

function apply(universe: Map<number, Set<string>>, level: number): Set<string> {
    const grid = universe.get(level);
    const s: Set<string> = new Set();

    for (let x = 0; x < 5; x++) {
        for (let y = 0; y < 5; y++) {
            if (x === 2 && y === 2) { // skip the middle, it can never have a bug
                continue;
            }


            let bugCount: number = 0;

            // in addition to the old ways, if we are on an edge (0 or 4), check outergrid
            const outerGrid = universe.get(level - 1);
            if (outerGrid) {
                if (x === 0) {
                    if (outerGrid.has(str(1, 2))) {
                        bugCount++;
                    }
                }
                else if (x === 4) {
                    if (outerGrid.has(str(3, 2))) {
                        bugCount++;
                    }
                }
                if (y === 0) {
                    if (outerGrid.has(str(2, 1))) {
                        bugCount++;
                    }
                } else if (y === 4) {
                    if (outerGrid.has(str(2, 3))) {
                        bugCount++;
                    }
                }
            }

            DIRECTIONS.forEach(t => {
                let [xi, yi] = [x + t[0], y + t[1]];
                const adjPoint = str(xi, yi);

                // if this is the middle, we can to figure out the 5 tiles to bring back
                if (adjPoint === MIDDLE) {
                    // need to check the 5 tiles in the inner grid
                    bugCount += lookInner(t, level, universe);
                } else { // old way
                    if (grid.has(adjPoint)) {
                        bugCount++;
                    }
                }
            });

            bugRules(grid, x, y, s, bugCount);
        }
    }

    return s;
}

function applyMinutes(universe: Map<number, Set<string>>, minutes: number): Map<number, Set<string>> {
    let min = 0;
    let max = 0;

    for (let i = 0; i < minutes; i++) {
        const iterverse: Map<number, Set<string>> = new Map();

        universe.forEach((grid, level) => {
            const s = apply(universe, level);
            iterverse.set(level, s);

            if (level < min) {
                min = level;
            }
            if (level > max) {
                max = level;
            }
        });

        // need to create add and apply them too
        universe.set(min - 1, new Set());
        universe.set(max + 1, new Set());
        const minGrid = apply(universe, min - 1);
        const maxGrid = apply(universe, max + 1);
        iterverse.set(min - 1, minGrid);
        iterverse.set(max + 1, maxGrid);

        universe = iterverse;
    }

    return universe;
}

function bugCount(universe: Map<number, Set<string>>): number {
    let sum: number = 0;
    universe.forEach((grid) => {
        sum += grid.size;
    });
    return sum;
}

function recurse(filename: string, minutes: number): number {
    const grid = makeGrid(readInput(filename));
    let universe: Map<number, Set<string>> = new Map();
    universe.set(0, grid);
    universe = applyMinutes(universe, minutes);
    return bugCount(universe);
}

function part1() {
    console.log('Part 1', findRepeat(makeGrid(readInput('input.txt'))));
}

function part2() {
    console.log('Part 2', recurse('input.txt', 200));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
