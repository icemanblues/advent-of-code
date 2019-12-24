import * as fs from 'fs';
import { getOrDefault } from '../util';

const dayNum: string = "08";
const dayTitle: string = "Space Image Format";

function readInputSync(filename: string): string {
    return fs.readFileSync(filename, "utf-8").trimRight();
}

const pixels: string = readInputSync('input.txt');
const width: number = 25;
const height: number = 6;
const layerLen: number = width * height;
const layers: string[] = splitLayer(pixels, width, height);

function splitLayer(pixels: string, w: number, h: number): string[] {
    const chunks: string[] = [];
    const chunk: number = w * h;
    for (let i: number = 0; i < pixels.length; i += chunk) {
        chunks.push(pixels.slice(i, i + chunk));
    }
    return chunks;
}

function validate(layers: string[]): number {
    let minZero: number = layerLen + 1;
    let value: number = 0;
    layers.forEach(l => {
        const counts: Map<string, number> = new Map<string, number>();
        for (let i: number = 0; i < l.length; i++) {
            const digit: string = l.charAt(i);

            const c: number = getOrDefault(counts, digit, 0);
            counts.set(digit, c + 1);
        }

        const numZero: number = getOrDefault(counts, '0', 0);
        if (numZero < minZero) {
            minZero = numZero;
            value = getOrDefault(counts, '1', 0) * getOrDefault(counts, '2', 0);
        }
    });

    return value;
}

function render(layers: string[]): string {
    const s: string[] = [];
    for (let i: number = 0; i < layerLen; i++) {
        let c: string = 'X';
        for (let j: number = 0; j < layers.length; j++) {
            const p: string = layers[j].charAt(i);
            if (p !== '2') {
                c = p;
                break;
            }
        }
        s.push(c);
    }
    return s.join('');
}

function printImage(image: string, w: number): void {
    for (let i: number = 0; i < image.length; i += w) {
        let line: string = image.slice(i, i + w);
        line = line.replace(/0/g, ' ');
        console.log(line);
    }
}

function part1() {
    console.log('Part 1', validate(layers));
}

function part2() {
    const layers: string[] = splitLayer(pixels, width, height);
    const image: string = render(layers);
    printImage(image, width);
    console.log('Part 2', 'CEKUA');
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
