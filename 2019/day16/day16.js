"use strict";
exports.__esModule = true;
var fs = require("fs");
var dayNum = "16";
var dayTitle = "Flawed Frequency Transmission";
function readInputSync(filename) {
    return fs.readFileSync(filename, "utf-8")
        .trimRight()
        .split('')
        .map(Number);
}
var BASE = [0, 1, 0, -1];
function computePattern(base, position) {
    var r = [];
    for (var _i = 0, BASE_1 = BASE; _i < BASE_1.length; _i++) {
        var b = BASE_1[_i];
        for (var i = 0; i <= position; i++) {
            r.push(b);
        }
    }
    return r;
}
function patternAt(pattern, position) {
    return pattern[(position + 1) % pattern.length];
}
function dotProduct(input, pattern, offset) {
    if (offset === void 0) { offset = 0; }
    var sum = 0;
    for (var i = 0; i < input.length; i++) {
        sum += input[i] * pattern[(i + 1 + offset) % pattern.length];
    }
    return Math.abs(sum) % 10;
}
function fft(input, phases, offset) {
    if (offset === void 0) { offset = 0; }
    var result = input;
    for (var i = 1; i <= phases; i++) {
        var temp = [];
        for (var j = 0; j < result.length; j++) {
            var pattern = computePattern(BASE, j);
            var dp = dotProduct(result, pattern, offset);
            temp.push(dp); // temp[j] =
        }
        result = temp;
    }
    return result.join('');
}
function dpOptimized(input) {
    var sum = 0;
    for (var i = input.length - 1; i >= 0; i--) {
        var pattern = computePattern(BASE, i);
        sum += input[i] * patternAt(pattern, i);
        input[i] = Math.abs(sum) % 10;
        console.log('iteration', i, 'out of ', input.length);
    }
}
function fftOptimized(input, phases) {
    var result = input;
    for (var p = 1; p <= phases; p++) {
        console.log('fft phase', p);
        dpOptimized(result);
    }
    return result.join('');
}
function part1() {
    console.log('Part 1');
    var input = readInputSync('input.txt');
    var output = fft(input, 100);
    console.log('Part 1', output.substring(0, 8));
}
function part2() {
    console.log('Part 2');
    var input = readInputSync('input.txt');
    var offset = Number(input.slice(0, 7).join(''));
    var big = [];
    for (var i = 0; i < 10000; i++) {
        for (var _i = 0, input_1 = input; _i < input_1.length; _i++) {
            var e = input_1[_i];
            big.push(e);
        }
    }
    big = big.slice(offset);
    var output = fftOptimized(big, 100);
    console.log('Part 1', output.substring(0, 8));
}
function main() {
    console.log("Day " + dayNum + " : " + dayTitle);
    part1(); // 82435530
    part2(); // 83036156
}
main();
