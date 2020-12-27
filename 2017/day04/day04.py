day_num = "04"
day_title = "High-Entropy Passphrases"


def read_input(filename):
    with open(filename) as f:
        content = [x.strip('\n') for x in f.readlines()]
    return content


def isValid(passphrase: str) -> bool:
    words = passphrase.split()
    s = set()
    for w in words:
        if w in s:
            return False
        s.add(w)
    return True


def isValidNoAnagrams(passphrase: str) -> bool:
    words = passphrase.split()
    for i in range(len(words)-1):
        for j in range(i+1, len(words)):
            if isAnagram(words[i], words[j]):
                return False
    return True


def isAnagram(a: str, b: str) -> bool:
    return sorted(list(a)) == sorted(list(b))


def part1():
    passphrases = read_input('input.txt')
    count = 0
    anagram_count = 0
    for phrase in passphrases:
        if isValid(phrase):
            count += 1
        if isValidNoAnagrams(phrase):
            anagram_count += 1

    print(f"Part 1: {count}")
    print(f"Part 2: {anagram_count}")


def main():
    print("Day", day_num, ":", day_title)
    part1()


if __name__ == '__main__':
    main()
