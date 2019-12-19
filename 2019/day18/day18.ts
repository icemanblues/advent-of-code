import fs from 'fs';
import { strt } from '../util';

const dayNum: string = "18";
const dayTitle: string = "Many-Worlds Interpretation";

function read(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

const input = read('input.txt');

function print(board: string[]): void {
    board.forEach(l => console.log(l));
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

class SearchState {
    loc: [number, number];
    keys: string[];
    steps: number;

    constructor(loc: [number, number], keys: string[], steps: number) {
        this.loc = loc;
        this.keys = keys;
        this.steps = steps;
    }
}

function tile(board: string[], x: number, y: number): string {
    return board[y].charAt(x);
}

function adj(board: string[], ss: SearchState): [number, number][] {
    const neighbors: [number, number][] = [];
    const [x, y] = ss.loc;
    [[0, -1], [0, 1], [-1, 0], [1, 0]].forEach(([i, j]) => {
        const xi = x + i;
        const yj = y + j;
        if (yj >= 0 && yj < board.length && xi >= 0 && xi < board[yj].length) {
            const t = tile(board, xi, yj);
            switch (t) {
                case '@': // start is valid
                case '.': // valid
                    neighbors.push([xi, yj]);
                    break;
                case '#': // wall
                    break;
                default:
                    // key
                    if (ALL_KEYS.includes(t)) {
                        neighbors.push([xi, yj]);
                    }
                    // door
                    else if (ss.keys.includes(t.toLowerCase())) {
                        neighbors.push([xi, yj]);
                    }
                    break;
            }
        }
    });
    return neighbors;
}

function memoKey(ss: SearchState, target: [number, number]): string {
    const mkeys = [...ss.keys]; // clone
    mkeys.sort();
    return strt(ss.loc) + '|' + strt(target) + '|' + mkeys.join('');
}

const memoBfs: Map<string, number> = new Map();
// returns steps from ss.loc to target
// -1 if it is not possible
function bfs(board: string[], ss: SearchState, target: [number, number]): number {
    // check memo cache first
    const memo = memoKey(ss, target);
    if (memoBfs.has(memo)) {
        return memoBfs.get(memo);
    }

    const queue: SearchState[] = [new SearchState(ss.loc, ss.keys, 0)];
    const visited: Set<string> = new Set();
    while (queue.length !== 0) {
        const curr = queue.shift();
        if (visited.has(strt(curr.loc))) {
            continue;
        }
        visited.add(strt(curr.loc));

        if (curr.loc[0] === target[0] && curr.loc[1] === target[1]) {
            memoBfs.set(memo, curr.steps);
            return curr.steps;
        }

        const neighbors = adj(board,
            new SearchState(curr.loc, curr.keys, curr.steps + 1));
        neighbors.forEach(n => queue.push(new SearchState(n, ss.keys, curr.steps + 1)));
    }

    memoBfs.set(memo, -1);
    return -1;
}

const ALL_KEYS = 'abcdefghijklmnopqrstuvwxyz';

function allKeysShortest(board: string[]): number {
    const start = find(board, '@');
    const startState = new SearchState(start, [], 0);
    const keys: Map<string, [number, number]> = new Map();
    const doors: Map<string, [number, number]> = new Map()
    for (const c of ALL_KEYS) {
        let [x, y] = find(board, c);
        if (x !== -1) {
            keys.set(c, [x, y]);
        }

        let [i, j] = find(board, c.toUpperCase());
        if (i !== -1) {
            doors.set(c.toUpperCase(), [i, j]);
        }
    }

    // from start to all reachable keys. then from each of those to all reachable keys
    // iter and repeat until all keys have been found
    // store those results and pick the lowest one
    let count = 0;
    const win: number[] = [];
    const queue: SearchState[] = [startState];
    while (queue.length !== 0) {
        const curr = queue.shift();
        if (curr.keys.length === keys.size) {
            win.push(curr.steps);
        }

        keys.forEach((keyLoc, key) => {
            if (!curr.keys.includes(key)) {
                const c = bfs(board, curr, keyLoc);
                if (c !== -1) {
                    const ss = new SearchState(keyLoc,
                        [...curr.keys, key],
                        curr.steps + c);
                    queue.push(ss);
                }
            }
        });
    }

    if (win.length === 0) {
        return -1;
    }

    return win.reduce((acc, curr) => acc < curr ? acc : curr, Number.MAX_VALUE);
}

function part1() {
    console.log('Part 1');
    // console.log('test-8', allKeysShortest(read('test-8.txt')));
    // console.log('test-81', allKeysShortest(read('test-81.txt')));
    // console.log('test-86', allKeysShortest(read('test-86.txt')));
    // console.log('test-132', allKeysShortest(read('test-132.txt')));
    console.log('test-136', allKeysShortest(read('test-136.txt')));
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
