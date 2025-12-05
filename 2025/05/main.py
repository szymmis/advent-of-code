def main():
    input = open("input.txt").read().strip().splitlines()

    print("Part One:", part_one(input))
    print("Part Two:", part_two(input))


def part_one(input: list[str]):
    count = 0
    ranges = []

    for line in input:
        if "-" in line:
            ranges.append(list(map(int, line.split("-"))))
        elif len(line) > 0:
            value = int(line)

            for range in ranges:
                if value >= range[0] and value <= range[1]:
                    count += 1
                    break

    return count


def part_two(input):
    fresh_ids = set()

    for line in input:
        if "-" in line:
            start, end = list(map(int, line.split("-")))
            for i in range(start, end + 1):
                fresh_ids.add(i)

    return len(fresh_ids)


if __name__ == "__main__":
    main()
