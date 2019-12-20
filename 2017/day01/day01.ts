import * as fs from 'fs';

const dayNum: string = "01";
const dayTitle: string = "Inverse Captcha";

function readInputSync(filename: string): string {
    return fs.readFileSync(filename, "utf-8").trimRight();
}

function circularCaptcha(s: string): number {
    let sum: number = 0;
    for (let i = 0; i < s.length - 1; i++) {
        const num1 = Number(s.charAt(i));
        const num2 = Number(s.charAt(i + 1));
        if (num1 === num2) {
            sum += num1;
        }
    }

    // compare the last with the first - the circular part
    const numFirst = Number(s.charAt(0));
    const numLast = Number(s.charAt(s.length - 1));
    if (numFirst === numLast) {
        sum += numLast;
    }

    return sum;
}

function halfwayCaptcha(s: string): number {
    let sum: number = 0;

    const halfway: number = s.length / 2;
    for (let i = 0; i < s.length; i++) {
        const num1: number = Number(s.charAt(i));

        const j: number = (i + halfway) % s.length;
        const num2: number = Number(s.charAt(j));

        if (num1 === num2) {
            sum += num1;
        }
    }

    return sum;
}

function part1() {
    console.log(`Part 1`);

    // tests
    console.log(3, circularCaptcha('1122'));
    console.log(4, circularCaptcha('1111'));
    console.log(0, circularCaptcha('1234'));
    console.log(9, circularCaptcha('91212129'));

    // real deal
    const t: string = readInputSync('input.txt');
    console.log('real deal', circularCaptcha(t));
}

function part2() {
    console.log(`Part 2`);

    // tests
    console.log(6, halfwayCaptcha('1212'));
    console.log(0, halfwayCaptcha('1221'));
    console.log(4, halfwayCaptcha('123425'));
    console.log(12, halfwayCaptcha('123123'));
    console.log(4, halfwayCaptcha('12131415'));

    // real deal
    const t: string = readInputSync('input.txt');
    console.log('real deal', halfwayCaptcha(t));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
