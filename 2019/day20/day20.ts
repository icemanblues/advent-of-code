import * as fs from 'fs';
import { getOrDefault, str, strt, tuple } from '../util';

const dayNum: string = "20";
const dayTitle: string = "Donut Maze";

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight().split(/\r?\n/);
}

function isLetter(s: string): boolean {
    if ('#' === s || '.' === s || ' ' === s) {
        return false;
    }
    return true;
}

class Donut {
    grid: Map<string, string>;
    name: Map<string, string[]>;
    warp: Map<string, string>;

    constructor(g: Map<string, string>, n: Map<string, string[]>, w: Map<string, string>) {
        this.grid = g;
        this.name = n;
        this.warp = w;
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

    for (let y = 0; y < input.length; y++) {
        for (let x = 0; x < input[y].length; x++) {
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
            function updateName(label: string, point: [number, number]) {
                if (isValid(point)) {
                    const arr = getOrDefault(name, label, []);
                    arr.push(strt(point));
                    name.set(label, arr);
                }
            }

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
    name.forEach((points, label) => {
        console.log(label, points);
        if (points.length == 2) {
            warp.set(points[0], points[1]);
            warp.set(points[1], points[0]);
        }
    })

    return new Donut(grid, name, warp);
}

class Robot {
    loc: [number, number];
    steps: number;
    constructor(loc: [number, number], steps: number) {
        this.loc = loc;
        this.steps = steps;
    }
}

function bfs(donut: Donut, start: [number, number], end: [number, number]): number {
    const queue: Robot[] = [new Robot(start, 0)];
    const visited: Set<string> = new Set();
    while (queue.length !== 0) {
        const curr = queue.shift();
        const currStr = strt(curr.loc);
        if (visited.has(currStr)) {
            continue;
        }
        visited.add(currStr);

        if (curr.loc[0] === end[0] && curr.loc[1] === end[1]) {
            return curr.steps;
        }

        adj(donut, curr.loc).forEach((p: [number, number]) => queue.push(new Robot(p, curr.steps + 1)));
    }

    return -1;
}

function adj(donut: Donut, [x, y]: [number, number]): [number, number][] {
    const neighbors: [number, number][] = [];
    [[0, 1], [0, -1], [1, 0], [-1, 0]].forEach(([i, j]) => {
        const [xi, yj] = [x + i, y + j];
        if (donut.grid.get(str(xi, yj)) === '.') {
            neighbors.push([xi, yj]);
        }
    });

    if (donut.warp.has(str(x, y))) {
        neighbors.push(tuple(donut.warp.get(str(x, y))));
    }

    return neighbors;
}

function short(donut: Donut) {
    const start: [number, number] = tuple(donut.name.get('AA')[0]);
    const end: [number, number] = tuple(donut.name.get('ZZ')[0]);
    return bfs(donut, start, end);
}

function part1() {
    console.log('Part 1');

    function run(filename: string) {
        const test = readInputSync(filename);
        const donut = parse(test);
        console.log(filename, short(donut));
    }

    run('test-small-23.txt');
    run('test-large-58.txt');
    run('input.txt');
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
