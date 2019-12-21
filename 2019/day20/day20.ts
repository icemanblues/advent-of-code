import * as fs from 'fs';
import { getOrDefault, str, strt, tuple } from '../util';

const dayNum: string = "20";
const dayTitle: string = "Donut Maze";

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight().split(/\r?\n/);
}

function isLetter(s: string): boolean {
    return !('#' === s || '.' === s || ' ' === s);
}

class Donut {
    grid: Map<string, string>;
    name: Map<string, string[]>;
    warp: Map<string, string>;
    maxWidth: number;
    maxHeight: number;

    constructor(g: Map<string, string>,
        n: Map<string, string[]>,
        w: Map<string, string>,
        maxX: number, maxY: number) {
        this.grid = g;
        this.name = n;
        this.warp = w;
        this.maxWidth = maxX;
        this.maxHeight = maxY;
    }
}

function parse(input: string[]): Donut {
    const grid: Map<string, string> = new Map();   // co-ordinates, tile
    const name: Map<string, string[]> = new Map(); // text to coordinates
    const warp: Map<string, string> = new Map();   // links two labels together

    function isValid([x, y]: [number, number]): boolean {
        if (y >= 0 && y < input.length && x >= 0 && x < input[y].length) {
            return input[y].charAt(x) === '.';
        }
        return false;
    }

    function updateName(label: string, point: [number, number]) {
        if (isValid(point)) {
            const arr = getOrDefault(name, label, []);
            arr.push(strt(point));
            name.set(label, arr);
        }
    }

    const maxY = input.length;
    const maxX = input[0].length;

    for (let y = 0; y < maxY; y++) {
        for (let x = 0; x < maxX; x++) {
            const tile = input[y].charAt(x);
            switch (tile) {
                case '#':
                case '.':
                    grid.set(str(x, y), tile);
                    break;
                case ' ':
                    break;
            }

            // check for the label and tile
            if (isLetter(tile)) {
                // check left
                if (x > 0) {
                    const left = input[y].charAt(x - 1);
                    if (isLetter(left)) {
                        const right: [number, number] = [x + 1, y];
                        const label = left + tile;
                        updateName(label, right);
                    }
                }

                // check right
                if (x < input[y].length - 1) {
                    const right = input[y].charAt(x + 1);
                    if (isLetter(right)) {
                        const left: [number, number] = [x - 1, y];
                        const label = tile + right;
                        updateName(label, left);
                    }
                }

                // check up
                if (y > 0) {
                    const up = input[y - 1].charAt(x);
                    if (isLetter(up)) {
                        const down: [number, number] = [x, y + 1];
                        const label = up + tile;
                        updateName(label, down);
                    }
                }

                // check down
                if (y < input.length - 1) {
                    const down = input[y + 1].charAt(x);
                    if (isLetter(down)) {
                        const up: [number, number] = [x, y - 1];
                        const label = tile + down;
                        updateName(label, up);
                    }
                }
            }
        }
    }

    // link the labels to the warps
    name.forEach((points) => {
        if (points.length == 2) {
            warp.set(points[0], points[1]);
            warp.set(points[1], points[0]);
        }
    })

    return new Donut(grid, name, warp, maxX, maxY);
}

class Robot {
    loc: [number, number];
    steps: number;
    level: number;
    constructor(loc: [number, number], steps: number, level: number) {
        this.loc = loc;
        this.steps = steps;
        this.level = level;
    }

    str(): string {
        return `${strt(this.loc)}|${this.level}`;
    }
}

function isOutsideWarp(donut: Donut, w: [number, number]): boolean {
    // need to account for the 2 characters that the labels take up
    return w[0] === 2 || w[1] === 2 ||
        donut.maxWidth - 3 === w[0] || donut.maxHeight - 3 === w[1];
}

function bfs(donut: Donut,
    start: [number, number], end: [number, number],
    recursive: boolean): number {
    const queue: Robot[] = [new Robot(start, 0, 0)];
    const visited: Set<string> = new Set();
    while (queue.length !== 0) {
        const curr = queue.shift();
        const currStr = curr.str();
        if (visited.has(currStr)) {
            continue;
        }
        visited.add(currStr);

        if (curr.loc[0] === end[0] && curr.loc[1] === end[1]) {
            if (!recursive) {
                return curr.steps;
            }

            // recursive... make sure that we are on the top most level
            if (curr.level === 0) {
                return curr.steps;
            }
        }

        adj(donut, curr, recursive).forEach((r: Robot) => queue.push(r));
    }

    return -1;
}

function adj(donut: Donut, robot: Robot, recursive: boolean): Robot[] {
    const neighbors: Robot[] = [];
    [[0, 1], [0, -1], [1, 0], [-1, 0]].forEach(([i, j]) => {
        const [xi, yj] = [robot.loc[0] + i, robot.loc[1] + j];
        if (donut.grid.get(str(xi, yj)) === '.') {
            neighbors.push(new Robot([xi, yj], robot.steps + 1, robot.level));
        }
    });

    const loc = strt(robot.loc);
    if (donut.warp.has(loc)) {
        const ow = isOutsideWarp(donut, robot.loc);
        const to = donut.warp.get(loc);
        const r = new Robot(tuple(to), robot.steps + 1, robot.level);
        if (recursive) {
            if (ow) {
                r.level -= 1;
            } else {
                r.level += 1;
            }
        }

        // no outside warps if recursive
        if (recursive && ow && robot.level === 0) {
        } else {
            neighbors.push(r);
        }
    }

    return neighbors;
}

function short(donut: Donut, recursive: boolean = false) {
    const start: [number, number] = tuple(donut.name.get('AA')[0]);
    const end: [number, number] = tuple(donut.name.get('ZZ')[0]);
    return bfs(donut, start, end, recursive);
}

function part1() {
    console.log('Part 1');

    function run(filename: string) {
        const test = readInputSync(filename);
        const donut = parse(test);
        console.log(filename, short(donut));
    }

    run('input.txt');
}

function part2() {
    console.log('Part 2');

    function run(filename: string, ans: number): void {
        const test = readInputSync(filename);
        const donut = parse(test);
        console.log(filename, ans, short(donut, true));
    }

    run('input.txt', 490);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
