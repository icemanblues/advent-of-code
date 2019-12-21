import * as fs from 'fs';
import { Amp, progAmp, prog } from '../intcode';

const dayNum: string = "21";
const dayTitle: string = "Springdroid Adventure";

function read(filename: string): number[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight().split(/,/).map(Number);
}

function readSpringscript(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight().split(/\n/);
}

const ASCII_LENGTH = 128; //0-127
const MAX_LENGTH = 15;

function toAscii(springscript: string[]): number[] {
    const r: number[] = [];
    springscript.forEach(s => {
        for (let c of s) {
            r.push(c.charCodeAt(0));
        }
        r.push(10); // new line
    });
    return r;
}

function display(output: number[]): number {
    let score: number = -1;
    let render: string[] = [];
    while (output.length !== 0) {
        const o = output.shift();
        if (o >= ASCII_LENGTH) {
            score = o;
        } else {
            render.push(String.fromCharCode(o));
        }
    }

    console.log(render.join(''));
    return score;
}

function jump(file: string) {
    const springscript: string[] = readSpringscript(file);
    const intcode = read('input.txt');
    const output: number[] = [];
    const input: number[] = toAscii(springscript);
    const amp = new Amp('Day 21', intcode, input, output);
    prog(amp);
    console.log(display(output));
}

function part1() {
    console.log('Part 1');
    jump('part1.txt');
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
