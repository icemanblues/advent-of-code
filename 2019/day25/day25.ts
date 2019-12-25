import { Amp, prog, parseIntcode } from '../intcode';
import { toAscii } from '../util';
import * as readline from 'readline-sync';
import * as fs from 'fs';

const dayNum: string = "25";
const dayTitle: string = "";

function printOutput(output: number[]): void {
    const lines: string[] = [];
    while (output.length > 0) {
        lines.push(String.fromCharCode(output.shift()));
    }
    console.log(lines.join(''));
}

function playMud(output: number[], inputFunc: () => number): void {
    const intcode = parseIntcode('input.txt');
    const amp = new Amp('santa', intcode, [], output);
    amp.inputCallback = inputFunc;
    prog(amp);
}

// interactive version to play the mud
function part1() {
    const input: number[] = [];
    const output: number[] = [];
    const interactive: () => number = () => {
        printOutput(output);

        if (input.length > 0) {
            return input.shift();
        }

        // read line from console
        const line = readline.prompt().trim();
        const command = toAscii(line);
        const c = command[0];
        const ommand = command.slice(1);
        input.push(...ommand);
        return c;
    }
    playMud(output, interactive);
    printOutput(output);
}

// uses the commands from the file to solve the puzzle
function part2() {
    const commands = fs.readFileSync('mud.txt', 'utf-8').trim().split(/\r?\n/);
    let i = 0;
    const input: number[] = [];
    const output: number[] = [];
    const bot: () => number = () => {
        output.splice(0, output.length);
        if (input.length > 0) {
            return input.shift();
        }
        const command = toAscii(commands[i++]);
        const c = command[0];
        const ommand = command.slice(1);
        input.push(...ommand);
        return c;
    }
    playMud(output, bot);
    printOutput(output);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    // part2();
}

main();
