//
//
//

import fs from 'fs';
import readline from 'readline';

const dayNum: string = "XX";
const dayTitle: string = "Title";
const subTitle1: string = "subtitle 1";
const subTitle2: string = "subtitle 2";

// reads entire file into memory, then splits it
function readInputSync(filename: string): string[] {
    const contents = fs.readFileSync(filename, "utf-8");
    const lines = contents.split('\n')
    return lines;
}

const readInput = async (filename: string): Promise<string[]> => {

    const lines: string[] = [];
    const fileStream = fs.createReadStream(filename);
    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity,
    });
    
    for await (const line of rl) {
        lines.push(line);
    }

    return lines;
}

const part1 = async function() {
    console.log(`Part 1: ${subTitle1}`);

    const synclines: string[] = readInputSync('testXX.txt');
    console.log(synclines);

    const lines: string[] = await readInput('testXX.txt');
    console.log(lines);

    // for (let s of lines) {
    //    console.log(s);
    // }
}

const part2 = function() {
    console.log(`Part 2: ${subTitle2}`);
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
