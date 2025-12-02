def main():
    input = open("input.txt").read().strip().splitlines()

    print("Part One:", part_one(input))
    print("Part Two:", part_two(input))


def part_one(input: list[str]):
    sum = 0

    def is_id_invalid(id: str):
        length_half = int(len(id) / 2)
        for i in range(0, length_half):
            if id[i] != id[i + length_half]:
                return False
        return True

    for _range in input[0].split(","):
        start, end = _range.split("-")
        start, end = int(start), int(end)

        for i in range(start, end + 1):
            id = str(i)
            if len(id) % 2 == 0:
                if is_id_invalid(id):
                    sum += i

    return sum


def part_two(input: list[str]):
    sum = 0

    def is_matching_pattern(id: str, pattern: str):
        for search_start in range(0, len(id), len(pattern)):
            if pattern != id[search_start : search_start + len(pattern)]:
                return False
        return True

    def is_id_invalid(id: str):
        for i in reversed(range(1, round(len(id) / 2) + 1)):
            pattern = id[0:i]
            if is_matching_pattern(id, pattern):
                return True
        return False

    for _range in input[0].split(","):
        start, end = _range.split("-")
        start, end = int(start), int(end)

        for id in range(start, end + 1):
            if is_id_invalid(str(id)):
                sum += id

    return sum


if __name__ == "__main__":
    main()
