def main():
    input = open("input.txt").read().strip().splitlines()

    print("Part One:", part_one(input))
    print("Part Two:", part_two(input))


def part_one(input):
    count = 0

    for y in range(len(input)):
        for x in range(len(input[y])):
            char = input[y][x]

            if char == "@":
                rolls_count = 0
                for i in range(-1, 2):
                    for j in range(-1, 2):
                        if (
                            (i != 0 or j != 0)
                            and x + j >= 0
                            and x + j < len(input[y])
                            and y + i >= 0
                            and y + i < len(input)
                        ):
                            if input[y + i][x + j] == "@":
                                rolls_count += 1
                if rolls_count < 4:
                    count += 1

    return count


def part_two(input: list[str]):
    grid = [list(line) for line in input]
    total_count = 0

    while True:
        count = 0
        for y in range(len(grid)):
            for x in range(len(grid[y])):
                char = grid[y][x]

                if char == "@":
                    rolls_count = 0
                    for i in range(-1, 2):
                        for j in range(-1, 2):
                            if (
                                (i != 0 or j != 0)
                                and x + j >= 0
                                and x + j < len(grid[y])
                                and y + i >= 0
                                and y + i < len(grid)
                            ):
                                if grid[y + i][x + j] == "@":
                                    rolls_count += 1
                    if rolls_count < 4:
                        count += 1
                        grid[y][x] = "."

        if count == 0:
            break
        else:
            total_count += count

    return total_count


if __name__ == "__main__":
    main()
