import fs from 'fs';

const dayNum: string = "02";
const dayTitle: string = "1202 Program Alarm";

// reads entire file into memory, then splits it
function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/,/);
    return lines;
}

function prog(intcodes: number[]): number[] {
    let i: number = 0;

    while (true) {
        let op: number = intcodes[i];

        if (99 === op) {
            break;
        }
        else if (1 === op) { // addition
            intcodes[intcodes[i + 3]] = intcodes[intcodes[i + 1]] + intcodes[intcodes[i + 2]];
            i += 4;
        }
        else if (2 === op) { // multiplication
            intcodes[intcodes[i + 3]] = intcodes[intcodes[i + 1]] * intcodes[intcodes[i + 2]];
            i += 4;
        }
        else {
            console.log('illegal op code', op);
            break;
        }
    }

    return intcodes;
}

function part1() {
    console.log('Part 1');

    const test1: number[] = [1, 0, 0, 0, 99];
    console.log(prog(test1), [2, 0, 0, 0, 99]);

    const test2: number[] = [2, 3, 0, 3, 99];
    console.log(prog(test2), [2, 3, 0, 6, 99]);

    const test3: number[] = [2, 4, 4, 5, 99, 0];
    console.log(prog(test3), [2, 4, 4, 5, 99, 9801]);

    const test4: number[] = [1, 1, 1, 4, 99, 5, 6, 0, 99];
    console.log(prog(test4), [30, 1, 1, 4, 2, 5, 6, 0, 99]);

    // the real thing
    const inputcodes: number[] = readInputSync('input.txt').map(x => Number(x));
    inputcodes[1] = 12;
    inputcodes[2] = 2;
    prog(inputcodes);
    console.log('Part1:', inputcodes[0]);

}

function nounVerb(noun: number, verb: number): number {
    return 100 * noun + verb;
}


function part2() {
    console.log('Part 2');

    const target: number = 19690720;
    const lines: string[] = readInputSync('input.txt');

    iter:
    for (let n: number = 0; n <= 99; n++) {
        for (let v: number = 0; v <= 99; v++) {
            const inputs: number[] = lines.map(x => Number(x));
            inputs[1] = n;
            inputs[2] = v;
            prog(inputs);
            if (inputs[0] === target) {
                console.log('part2:', nounVerb(n, v));
                break iter;
            }
        }
    }
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
