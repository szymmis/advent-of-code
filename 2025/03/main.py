def main():
    input = open("input.txt").read().strip().splitlines()

    print("Part One:", part_one(input))
    print("Part Two:", part_two(input))


def part_one(input: list[str]):
    sum = 0

    for line in input:
        first_ptr = 0
        second_ptr = 1
        for i in range(1, len(line) - 1):
            value = int(line[i])

            if value > int(line[first_ptr]):
                first_ptr = i
                second_ptr = i + 1
            elif value > int(line[second_ptr]):
                second_ptr = i

        if int(line[-1]) > int(line[second_ptr]):
            second_ptr = len(line) - 1

        sum += 10 * int(line[first_ptr]) + int(line[second_ptr])

    return sum


def part_two(input):
    return -1


if __name__ == "__main__":
    main()
