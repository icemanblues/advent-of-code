import * as fs from 'fs';
import bigInt from 'big-integer';

const dayNum: string = "22";
const dayTitle: string = "Slam Shuffle";

function readInput(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8").trimRight().split(/\r?\n/);
}

interface Deck {
    dealNewStack(): void;
    cutCards(n: number): void;
    dealIncrement(n: number): void;
}

class DeckCards implements Deck {
    cards: number[];

    constructor(cards: number[]) {
        this.cards = cards;
    }

    dealNewStack() {
        this.cards.reverse();
    }

    cutCards(n: number) {
        let p = n;
        if (p < 0) {
            p = this.cards.length + n;
        }

        this.cards = [...this.cards.slice(p), ...this.cards.slice(0, p)];
    }

    dealIncrement(n: number) {
        const inc = new Array<number>(this.cards.length);
        for (let i = 0; i < this.cards.length; i++) {
            const idx = (n * i) % this.cards.length;
            inc[idx] = this.cards[i];
        }
        this.cards = inc;
    }
}

function newDeck(size: number): DeckCards {
    const cards = new Array<number>(size);
    for (let i = 0; i < size; i++) {
        cards[i] = i;
    }
    return new DeckCards(cards);
}

function shuffle(deck: Deck, commands: string[]) {
    commands.forEach(l => {
        if ("deal into new stack" === l) {
            deck.dealNewStack();
        }
        else if (l.startsWith('deal with increment')) {
            const n = Number(l.split(/deal with increment /)[1]);
            deck.dealIncrement(n);
        }
        else if (l.startsWith('cut')) {
            const n = Number(l.split(/cut /)[1]);
            deck.cutCards(n);
        }
        else {
            console.log('unknown command', l);
        }
    });
}

class D implements Deck {
    length: bigInt.BigInteger;
    offset: bigInt.BigInteger;
    increment: bigInt.BigInteger;

    constructor(length: bigInt.BigInteger) {
        this.length = length;
        this.offset = bigInt(0);
        this.increment = bigInt(1);
    }

    dealNewStack(): void {
        this.increment = this.increment.multiply(-1);
        this.offset = this.offset.add(this.increment);
    }

    cutCards(n: number): void {
        this.offset = this.offset.add(this.increment.multiply(n));
    }

    dealIncrement(n: number): void {
        this.increment = this.increment.multiply(
            inv(bigInt(n), this.length)
        );
    }

    index(n: number): bigInt.BigInteger {
        return this.increment.multiply(n).add(this.offset).mod(this.length);
    }
}

// inv(n) = pow(n, MOD-2, MOD)
function inv(n: bigInt.BigInteger, mod: bigInt.BigInteger): bigInt.BigInteger {
    const modMinusTwo = mod.minus(2);
    return n.modPow(modMinusTwo, mod);
}

function part1() {
    const DECK_LENGTH = 10007;
    const deck = newDeck(DECK_LENGTH);
    const commands = readInput('input.txt');
    shuffle(deck, commands);
    console.log('Part 1', deck.cards.indexOf(2019));
}

function part2() {
    const BIG_DECK = bigInt(119315717514047);
    const BIG_SHUFFLE = bigInt(101741582076661);
    const commands = readInput('input.txt');
    const deck = new D(BIG_DECK);

    shuffle(deck, commands);
    const offset_diff = deck.offset;
    const increment_mul = deck.increment;

    const increment = increment_mul.modPow(BIG_SHUFFLE, BIG_DECK);

    const offset_diff_multipler = bigInt(1).minus(
        increment_mul.modPow(BIG_SHUFFLE, BIG_DECK)
    );
    const offset_diff_inv = inv(bigInt(1).minus(increment_mul), BIG_DECK);
    const offset = offset_diff
        .multiply(offset_diff_multipler)
        .multiply(offset_diff_inv);

    // value at location n
    const value = offset.add(increment.multiply(2020));
    console.log('Part 2', '2020', BIG_DECK.add(value.mod(BIG_DECK)));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
