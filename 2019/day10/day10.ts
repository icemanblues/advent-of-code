import fs from 'fs';

const dayNum: string = "10";
const dayTitle: string = "Monitoring Station";

function readInputSync(filename: string): string[] {
    const contents: string = fs.readFileSync(filename, "utf-8");
    const lines: string[] = contents.trimRight().split(/\r?\n/);
    return lines;
}

function asteroids(lines: string[]): [number, number][] {
    const a: [number, number][] = [];
    lines.forEach((l, y) => {
        for (let x: number = 0; x < l.length; x++) {
            if (l.charAt(x) === '#') {
                a.push([x, y]);
            }
        }
    });
    return a;
}

function str(x: number, y: number): string {
    return `${x},${y}`;
}

function strt(t: [number, number]): string {
    return str(t[0], t[1]);
}

function asteroidDetect(filename: string): number {
    const asteroidMap: string[] = readInputSync(filename);
    const asteroidList: [number, number][] = asteroids(asteroidMap);

    function isAsteroid(x: number, y: number): boolean {
        if (x < 0 || x >= asteroidMap[0].length) {
            return false;
        }
        if (y < 0 || y >= asteroidMap.length) {
            return false;
        }
        return asteroidMap[y].charAt(x) === '#';
    }

    let maxCount: number = -1;
    let maxTuple: [number, number] = [-1, -1];

    asteroidList.forEach(a => {
        const s: Set<string> = new Set<string>();
        asteroidList.forEach(t => s.add(strt(t)));
        s.delete(strt(a));

        // current row left
        let i: number = a[0] - 1;
        let crl: number = 0;
        while (i >= 0) {
            if (isAsteroid(i, a[1])) {
                if (crl !== 0) {
                    s.delete(str(i, a[1]));
                }
                crl++;
            }
            i--;
        }


        // current row right
        i = a[0] + 1;
        crl = 0;
        while (i < asteroidMap[a[1]].length) {
            if (isAsteroid(i, a[1])) {
                if (crl !== 0) {
                    s.delete(str(i, a[1]));
                }
                crl++;
            }
            i++;
        }

        // row up
        for (let yi: number = a[1] - 1; yi >= 0; yi--) {
            const check: [number, number][] =
                asteroidList.filter(e => e[1] === yi && s.has(strt(e)));
            check.forEach(c => {
                const dx: number = c[0] - a[0];
                const dy: number = c[1] - a[1];
                const slope: number = dx / dy;

                let los: [number, number] = [a[0], a[1]];
                for (let yii: number = yi - 1; yii >= 0; yii--) {
                    const deltaY: number = yii - a[1];
                    const deltaX: number = slope * deltaY;

                    if (!Number.isInteger(deltaX)) {
                        continue;
                    }

                    los[0] = a[0] + deltaX;
                    los[1] = a[1] + deltaY;

                    if (isAsteroid(los[0], los[1])) {
                        s.delete(strt(los));
                    }
                }
            });
        }

        // row down
        for (let yi: number = a[1] + 1; yi < asteroidMap.length; yi++) {
            const check: [number, number][] =
                asteroidList.filter(e => e[1] === yi && s.has(strt(e)));
            check.forEach(c => {
                const dx: number = c[0] - a[0];
                const dy: number = c[1] - a[1];
                const slope: number = dx / dy;

                let los: [number, number] = [a[0], a[1]];
                for (let yii: number = yi + 1; yii < asteroidMap.length; yii++) {
                    const deltaY: number = yii - a[1];
                    const deltaX: number = slope * deltaY;

                    if (!Number.isInteger(deltaX)) {
                        continue;
                    }
                    los[0] = a[0] + deltaX;
                    los[1] = a[1] + deltaY;

                    if (isAsteroid(los[0], los[1])) {
                        s.delete(strt(los));
                    }
                }
            });
        }

        if (maxCount < s.size) {
            maxCount = s.size;
            maxTuple = [a[0], a[1]];
        }
    });

    console.log(filename, maxTuple);
    return maxCount;
}

class Ast {
    x: number;
    y: number;
    dx: number;
    dy: number;

    constructor(x: number, y: number, dx: number, dy: number) {
        this.x = x;
        this.y = y;
        this.dx = dx;
        this.dy = dy;
    }

    slope(): number {
        return this.dy / this.dx;
    }

    dist(point: [number, number]): number {
        return Math.abs(this.x - point[0]) + Math.abs(this.y - point[1]);
    }
}

function fireLaser(filename: string, start: [number, number]): number {
    let t: Ast;
    let count: number = 0;

    const asteroidMap: string[] = readInputSync(filename);
    const asteroidList: [number, number][] = asteroids(asteroidMap);

    const targets: Ast[] = asteroidList
        .filter(a => a[0] !== start[0] || a[1] !== start[1])
        .map(a => {
            return new Ast(a[0], a[1], a[0] - start[0], a[1] - start[1]);
        });

    const distSort = (n1: Ast, n2: Ast) => n1.dist(start) - n2.dist(start);
    const steepSlopeSort = (n1: Ast, n2: Ast) => {
        const diff: number = n1.slope() - n2.slope()
        if (diff !== 0) {
            return diff;
        }
        return distSort(n1, n2);
    };
    const flatSlopeSort = (n1: Ast, n2: Ast) => {
        const diff: number = n2.slope() - n1.slope()
        if (diff !== 0) {
            return diff;
        }
        return distSort(n1, n2);
    };

    const fire = (a: Ast) => {
        count++;
        targets.splice(targets.indexOf(a), 1);
        if (count === 200) {
            t = a;
        }
    }

    const first = (direction: Ast[]) => {
        if (direction.length > 0) {
            fire(direction[0]);
        }
    };

    const notSameSlope = (direction: Ast[]) => {
        let slopers: number;
        for (let i: number = 0; i < direction.length; i++) {
            const a: Ast = direction[i];
            if (slopers === a.slope()) {
                continue;
            }
            fire(a);
            slopers = a.slope();
        }
    };

    while (count < 200) {
        // up
        const uppers: Ast[] = targets
            .filter(t => t.x === start[0] && t.y < start[1])
            .sort(distSort);
        first(uppers);

        // positive slope (-dy +dx)
        const upright: Ast[] = targets
            .filter(t => t.dx > 0 && t.dy < 0)
            .sort(steepSlopeSort);
        notSameSlope(upright);

        // right
        const right: Ast[] = targets
            .filter(t => t.y === start[1] && t.x > start[0])
            .sort(distSort);
        first(right);

        // negative slope(+dy, +dx)
        const downright: Ast[] = targets
            .filter(t => t.dx > 0 && t.dy > 0)
            .sort(flatSlopeSort);
        notSameSlope(downright);

        // down
        const down: Ast[] = targets
            .filter(t => t.x === start[0] && t.y > start[1])
            .sort(distSort);
        first(down);

        // positive slope (+dy -dx)
        const downleft: Ast[] = targets
            .filter(t => t.dx < 0 && t.dy > 0)
            .sort(steepSlopeSort);
        notSameSlope(downleft);

        // left
        const left: Ast[] = targets
            .filter(t => t.y === start[1] && t.x < start[0])
            .sort(distSort);
        first(left);

        // negative slope (-dy, -dx)
        const upleft: Ast[] = targets
            .filter(t => t.dx < 0 && t.dy < 0)
            .sort(steepSlopeSort);
        notSameSlope(upleft);
    }

    return 100 * t.x + t.y;
}

function part1() {
    console.log('Part 1', asteroidDetect('input.txt'));
}

function part2() {
    console.log('Part 2', fireLaser('input.txt', [28, 29]));
}

function main() {
    console.log(`Day ${dayNum} : ${dayTitle}`);

    part1();
    part2();
}

main();
