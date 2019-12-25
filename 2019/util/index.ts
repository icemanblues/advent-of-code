export function getOrDefault<K, V>(m: Map<K, V>, k: K, d: V) {
    if (!k) {
        console.log('undefined?', k, m);
    }
    return m.has(k) ? m.get(k) : d;
}

export function str(x: number, y: number): string {
    return `${x},${y}`;
}

export function strt(tup: [number, number]): string {
    return str(tup[0], tup[1]);
}

export function tuple(s: string): [number, number] {
    const nums: number[] = s.split(/,/).map(Number);
    return [nums[0], nums[1]];
}

export function toAsciiMulti(script: string[]): number[] {
    let r: number[] = [];
    script.forEach(s => {
        r.push(...toAscii(s));
    });
    return r;
}

export function toAscii(command: string): number[] {
    const r: number[] = [];
    for (let c of command) {
        r.push(c.charCodeAt(0));
    }
    r.push(10); // new line
    return r;
}
