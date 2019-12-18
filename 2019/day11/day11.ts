import fs from 'fs';
import { getOrDefault, str } from '../util';
import { Amp, progAmp } from '../intcode';

const dayNum: string = "11";
const dayTitle: string = "Space Police";

function readInputSync(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/,/);
}

function progPaint(amp: Amp, pb: Paintbot): number {
    amp.inputCallback = () => pb.scan();

    let i: number = 0;
    while (!amp.isHalted) {
        progAmp(amp);
        const o: number = amp.outputs.shift();

        if (i % 2 === 0) { // paint
            pb.paint(o);
        } else { // turn and step
            pb.move(o);
        }
        i++;
    }

    return pb.paintMap.size;
}

enum Direction {
    Up,
    Right,
    Down,
    Left,
}

const directionOffset: Map<Direction, [number, number]> = new Map([
    [Direction.Up, [0, -1]],
    [Direction.Right, [1, 0]],
    [Direction.Down, [0, 1]],
    [Direction.Left, [-1, 0]],
]);

function step(start: [number, number], dir: Direction): [number, number] {
    const offset: [number, number] = directionOffset.get(dir);
    return [start[0] + offset[0], start[1] + offset[1]];
}

enum Turn {
    Left,
    Right,
}

function turn90(dir: Direction, turn: Turn): Direction {
    if (turn === Turn.Left) {
        switch (dir) {
            case Direction.Up:
                return Direction.Left;
            case Direction.Left:
                return Direction.Down;
            case Direction.Down:
                return Direction.Right;
            case Direction.Right:
                return Direction.Up;
        }
    } else if (turn === Turn.Right) {
        switch (dir) {
            case Direction.Up:
                return Direction.Right;
            case Direction.Right:
                return Direction.Down;
            case Direction.Down:
                return Direction.Left;
            case Direction.Left:
                return Direction.Up;
        }
    }

    console.log('unknown Turn', turn);
    return Direction.Up;
}

class Paintbot {
    x: number;
    y: number;
    dir: Direction;
    paintMap: Map<string, number>

    constructor() {
        this.x = 0;
        this.y = 0;
        this.dir = Direction.Up;
        this.paintMap = new Map<string, number>();
    }

    scan(): number {
        return getOrDefault(this.paintMap, str(this.x, this.y), 0);
    }

    paint(color: number) {
        this.paintMap.set(str(this.x, this.y), color);
    }

    move(turn: number) {
        this.dir = turn90(this.dir, turn);
        [this.x, this.y] = step([this.x, this.y], this.dir);
    }

    print() {
        for (let y: number = 0; y < 6; y++) {
            const line: string[] = [];
            for (let x: number = 0; x < 40; x++) {
                if (this.paintMap.get(str(x, y)) === 1) {
                    line[x] = '#';
                } else {
                    line[x] = ' ';
                }
            }
            console.log(line.join(''));
        }
    }
}

function part1() {
    console.log('Part 1');
    const lines: number[] = readInputSync('input.txt').map(Number);
    const amp: Amp = new Amp('part1', lines, [], []);
    const paintbot: Paintbot = new Paintbot();
    console.log(progPaint(amp, paintbot));
}

function part2() {
    console.log('Part 2');
    const lines: number[] = readInputSync('input.txt').map(Number);
    const amp: Amp = new Amp('part2', lines, [], []);
    const paintbot: Paintbot = new Paintbot();

    paintbot.paint(1);
    progPaint(amp, paintbot)

    paintbot.print();
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
