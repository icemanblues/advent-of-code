import { str, tuple } from '../util';
import { Amp, prog, parseIntcode } from '../intcode';

const dayNum: string = "13";
const dayTitle: string = "Care Package";

function count(game: number[], tile: number): number {
    const board = gameBoard(game);
    let counter: number = 0;
    board.forEach((v, k) => {
        if (v === tile) {
            counter++;
        }
    });
    return counter;
}

function gameBoard(game: number[]): Map<string, number> {
    const board: Map<string, number> = new Map<string, number>();
    for (let i: number = 0; i < game.length - 2; i += 3) {
        const key: string = str(game[i], game[i + 1]);
        const value: number = game[i + 2];
        board.set(key, value);
    }
    return board;
}

function score(board: Map<string, number>): number {
    return board.get(str(-1, 0));
}

function paint(board: Map<string, number>) {
    let maxX: number = -1;
    let maxY: number = -1;
    for (let tile of board.keys()) {
        const [tx, ty] = tuple(tile);
        if (maxX < tx) {
            maxX = tx;
        }
        if (maxY < ty) {
            maxY = ty;
        }
    }

    for (let y: number = 0; y <= maxY; y++) {
        let line: string[] = [];
        for (let x: number = 0; x <= maxX; x++) {
            const key: string = str(x, y);
            const value: number = board.get(key);
            if (value === 0) {
                line.push(' ');
            } else {
                line.push(String(value));
            }
            line.push();
        }
        console.log(line.join(''));
    }
    console.log('score', score(board));
}

function part1() {
    const intcode: number[] = parseIntcode('input.txt');
    const output: number[] = [];
    const amp: Amp = new Amp('Part 1', intcode, [], output);
    prog(amp);
    console.log('Part 1', count(output, 2));
}

function part2() {
    const intcode: number[] = parseIntcode('input.txt');
    intcode[0] = 2;
    const output: number[] = [];
    const amp: Amp = new Amp('Part 2', intcode, [], output);

    const autopilot = () => {
        const board = gameBoard(output);
        //paint(board);

        let ball: number = -1;
        let paddle: number = -1;
        board.forEach((v, k) => {
            if (v === 4) {
                ball = Number(k.split(/,/)[0]);
            }
            if (v === 3) {
                paddle = Number(k.split(/,/)[0]);
            }
        });

        if (ball < paddle) {
            return -1;
        } else if (ball > paddle) {
            return 1;
        }
        return 0;
    };

    amp.inputCallback = autopilot;
    prog(amp);
    const board = gameBoard(output);
    console.log('Part 2', score(board));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
