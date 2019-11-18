//
//
//

import fs from 'fs';
import readline from 'readline';
import { setupMaster } from 'cluster';

const dayNum: string = "01";
const dayTitle: string = "Inverse Captcha";
const subTitle1: string = "subtitle 1";
const subTitle2: string = "subtitle 2";

// reads entire file into memory, then splits it
function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/\r?\n/);
    return lines;
}

async function readInput(filename: string): Promise<string[]> {
    const lines: string[] = [];
    const fileStream: fs.ReadStream = fs.createReadStream(filename);
    const rl: readline.Interface = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity,
    });
    
    for await (const line of rl) {
        lines.push(line);
    }

    return lines;
}

function circularCaptcha(s: string): number {
    let sum: number = 0;
    for(let i=0; i<s.length-1; i++) {
        const num1: number = parseInt(s.charAt(i));
        const num2: number = parseInt(s.charAt(i+1));
        if(num1 === num2) {
            sum += num1;
        }
    }

    // compare the last with the first - the circular part
    const numFirst: number = parseInt(s.charAt(0));
    const numLast: number = parseInt(s.charAt(s.length-1));
    if(numFirst === numLast) {
        sum += numLast;
    }

    return sum;
}

function halfwayCaptcha(s: string): number {
    let sum: number = 0;

    const halfway: number = s.length / 2;
    for(let i=0; i<s.length; i++) {
        const num1: number = parseInt(s.charAt(i));

        const j: number = (i+halfway) % s.length;
        const num2: number = parseInt(s.charAt(j));

        if(num1 === num2) {
            sum += num1;
        }
    }

    return sum;
}

async function part1() {
    console.log(`Part 1: ${subTitle1}`);

    // tests
    console.log(3, circularCaptcha('1122'));
    console.log(4, circularCaptcha('1111'));
    console.log(0, circularCaptcha('1234'));
    console.log(9, circularCaptcha('91212129'));

    // real deal
    const t: string = fs.readFileSync('input.txt', "utf-8").trimRight();
    console.log('real deal', circularCaptcha(t));
}

async function part2() {
    console.log(`Part 2: ${subTitle2}`);

    // tests
    console.log(6, halfwayCaptcha('1212'));
    console.log(0, halfwayCaptcha('1221'));
    console.log(4, halfwayCaptcha('123425'));
    console.log(12, halfwayCaptcha('123123'));
    console.log(4, halfwayCaptcha('12131415'));

    // real deal
    const t: string = fs.readFileSync('input.txt', "utf-8").trimRight();
    console.log('real deal', halfwayCaptcha(t));
}

async function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    await part1();
    await part2();
}

main();
