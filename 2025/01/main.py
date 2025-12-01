import math


def main():
    input = open("input.txt").read().strip().splitlines()

    print("Part One:", part_one(input))
    print("Part Two:", part_two(input))


def part_one(input: list[str]):
    position = 50
    zeros_count = 0

    for line in input:
        move = int(line[1:])
        move = -move if line[0] == "L" else move

        position = (position + 100 + move) % 100
        if position == 0:
            zeros_count += 1

    return zeros_count


def part_two(input: list[str]):
    position = 50
    zeros_count = 0

    for line in input:
        move = int(line[1:])
        cycles = math.floor(move / 100)
        move = -(move % 100) if line[0] == "L" else move % 100

        if (position + move) >= 100:
            cycles += 1
        elif position > 0 and (position + move) <= 0:
            cycles += 1

        position = (position + 100 + move) % 100
        zeros_count += cycles

    return zeros_count


if __name__ == "__main__":
    main()
