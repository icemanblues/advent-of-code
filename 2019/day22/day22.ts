import * as fs from 'fs';

const dayNum: string = "22";
const dayTitle: string = "Slam Shuffle";

function readInput(filename: string): string[] {
    return fs.readFileSync(filename, "utf-8")
        .trimRight().split(/\r?\n/);
}

const DECK_LENGTH = 10007;

class Deck {
    cards: number[];

    constructor(cards: number[]) {
        this.cards = cards;
    }
}

function newDeck(size: number): Deck {
    const cards = new Array<number>(size);
    for (let i = 0; i < size; i++) {
        cards[i] = i;
    }
    return new Deck(cards);
}

function dealNewStack(deck: Deck) {
    deck.cards.reverse();
}

function cutCards(deck: Deck, n: number) {
    let p = n;
    if (p < 0) {
        p = deck.cards.length + n;
    }

    deck.cards = [...deck.cards.slice(p), ...deck.cards.slice(0, p)];
}

function dealIncrement(deck: Deck, n: number) {
    const inc = new Array<number>(deck.cards.length);
    for (let i = 0; i < deck.cards.length; i++) {
        const idx = (n * i) % deck.cards.length;
        inc[idx] = deck.cards[i];
    }
    deck.cards = inc;
}

function shuffle(deck: Deck, commands: string[]) {
    commands.forEach(l => {
        if ("deal into new stack" === l) {
            dealNewStack(deck);
        }
        else if (l.startsWith('deal with increment')) {
            const n = Number(l.split(/deal with increment /)[1]);
            dealIncrement(deck, n);
        }
        else if (l.startsWith('cut')) {
            const n = Number(l.split(/cut /)[1]);
            cutCards(deck, n);
        }
        else {
            console.log('unknown command', l);
        }
    });
}

function test1() {
    console.log('Test 1');

    let deck = newDeck(10);
    dealNewStack(deck);
    console.log('test new stack', '9 8 7 6 5 4 3 2 1 0', deck.cards);

    deck = newDeck(10);
    cutCards(deck, 3);
    console.log('test cut', '3 4 5 6 7 8 9 0 1 2', deck.cards);

    deck = newDeck(10);
    cutCards(deck, -4);
    console.log('test cut neg', '6 7 8 9 0 1 2 3 4 5', deck.cards);

    deck = newDeck(10);
    dealIncrement(deck, 3);
    console.log('test increment', '0 7 4 1 8 5 2 9 6 3', deck.cards);

    deck = newDeck(10);
    let commands: string[] = [
        'deal with increment 7',
        'deal into new stack',
        'deal into new stack'
    ];
    shuffle(deck, commands);
    console.log('0 3 6 9 2 5 8 1 4 7', deck.cards);

    deck = newDeck(10);
    commands = [
        'cut 6',
        'deal with increment 7',
        'deal into new stack'
    ];
    shuffle(deck, commands);
    console.log('3 0 7 4 1 8 5 2 9 6', deck.cards);

    deck = newDeck(10);
    commands = [
        'deal with increment 7',
        'deal with increment 9',
        'cut -2'
    ];
    shuffle(deck, commands);
    console.log('6 3 0 7 4 1 8 5 2 9', deck.cards);

    deck = newDeck(10);
    commands = [
        'deal into new stack',
        'cut -2',
        'deal with increment 7',
        'cut 8',
        'cut -4',
        'deal with increment 7',
        'cut 3',
        'deal with increment 9',
        'deal with increment 3',
        'cut -1'
    ];
    shuffle(deck, commands);
    console.log('9 2 5 8 1 4 7 0 3 6', deck.cards);
}

function part1() {
    console.log('Part 1');
    const deck = newDeck(DECK_LENGTH);
    const commands = readInput('input.txt');
    shuffle(deck, commands);
    console.log(deck.cards.indexOf(2019));
}

function part2() {
    console.log('Part 2');
    const BIG_DECK = 119315717514047;
    const deck = newDeck(BIG_DECK);
    const commands = readInput('input.txt');
    const BIG_SHUFFLE = 101741582076661;
    for (let i = 0; i < BIG_SHUFFLE; i++) {
        shuffle(deck, commands);
    }
    console.log(deck.cards[2020]);

}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);
    part1();
    part2();
}

main();
