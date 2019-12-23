const dayNum: string = "04";
const dayTitle: string = "Secure Container";

const [INPUT_MIN, INPUT_MAX] = [197487, 673251];

function sixDigit(s: string): boolean {
    return s.length === 6;
}

function withinRange(s: string): boolean {
    const n = Number(s);
    return n >= INPUT_MIN && n <= INPUT_MAX;
}

function adjacentSame(s: string): boolean {
    for (let i = 0; i < s.length - 1; i++) {
        if (s.charAt(i) === s.charAt(i + 1)) {
            return true;
        }
    }
    return false;
}

function neverDecrease(s: string): boolean {
    let n = Number(s.charAt(0));
    for (let i = 1; i < s.length; i++) {
        const j = Number(s.charAt(i));
        if (j < n) {
            return false;
        }
        n = j;
    }

    return true;
}

function isValid(s: string): boolean {
    return sixDigit(s) && withinRange(s) && adjacentSame(s) && neverDecrease(s);
}

function password(valid: (s: string) => boolean): number {
    let count = 0;
    for (let i = INPUT_MIN; i <= INPUT_MAX; i++) {
        const s = String(i);
        if (valid(s)) {
            count++;
        }
    }
    return count;
}

function part1() {
    console.log('Part 1');
    console.log(password(isValid));
}

function adjOnlyTwo(s: string): boolean {
    let c = s.charAt(0);
    let count = 1;

    let i = 1;
    while (i < s.length) {
        if (c === s.charAt(i)) {
            count++;
        } else {
            if (count === 2) {
                return true;
            }
            count = 1;
            c = s.charAt(i);
        }

        i++;
    }

    return count === 2;
}

function isValidTwo(s: string): boolean {
    return sixDigit(s) && withinRange(s) && adjOnlyTwo(s) && neverDecrease(s);
}

function part2() {
    console.log('Part 2');
    console.log(password(isValidTwo));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
