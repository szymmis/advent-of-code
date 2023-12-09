fn extrapolate(values: &[i32]) -> i32 {
    let mut diffs = Vec::new();

    for w in values.windows(2) {
        let [a, b] = w else { panic!() };
        diffs.push(b - a)
    }

    let next_val = if diffs.iter().skip(1).all(|v| v == diffs.first().unwrap()) {
        diffs[0]
    } else {
        extrapolate(&diffs)
    };

    values.last().unwrap() + next_val
}

fn part_one(input: &str) -> i32 {
    input
        .lines()
        .map(|line| {
            let values = line
                .split_ascii_whitespace()
                .flat_map(str::parse)
                .collect::<Vec<i32>>();

            extrapolate(&values)
        })
        .sum()
}

fn part_two(input: &str) -> i32 {
    input
        .lines()
        .map(|line| {
            let values = line
                .split_ascii_whitespace()
                .flat_map(str::parse)
                .rev()
                .collect::<Vec<i32>>();

            extrapolate(&values)
        })
        .sum()
}

fn main() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    println!("Part one: {}", part_one(&input));
    println!("Part two: {}", part_two(&input))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45";
        assert_eq!(part_one(input), 114);
    }

    #[test]
    fn part_two_example() {
        let input = "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45";
        assert_eq!(part_two(input), 2);
    }
}
