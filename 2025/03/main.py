def main():
    input = open("input.txt").read().strip().splitlines()

    print("Part One:", part_one(input))
    print("Part Two:", part_two(input))


def compute_max_number(line: str, digits_count: int):
    sum = 0
    digits_ptrs = [i for i in range(digits_count)]

    for char_idx in range(0, len(line)):
        char_value = int(line[char_idx])

        for idx, ptr in enumerate(digits_ptrs):
            if char_idx > ptr and char_idx <= len(line) - (digits_count - idx):
                if char_value > int(line[ptr]):
                    for offset, i in enumerate(range(idx, digits_count)):
                        digits_ptrs[i] = char_idx + offset
                    break

    for base, ptr in enumerate(reversed(digits_ptrs)):
        sum += pow(10, base) * int(line[ptr])

    return sum


def part_one(input: list[str]):
    sum = 0
    for line in input:
        sum += compute_max_number(line, digits_count=2)
    return sum


def part_two(input):
    sum = 0
    for line in input:
        sum += compute_max_number(line, digits_count=12)
    return sum


if __name__ == "__main__":
    main()
