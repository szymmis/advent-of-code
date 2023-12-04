fn parse_numbers(line: &str) -> (Vec<i32>, Vec<i32>) {
    let mut split = line.split('|');

    let winning_numbers = split
        .next()
        .unwrap()
        .split(':')
        .nth(1)
        .unwrap()
        .split(' ')
        .filter_map(|str| {
            if !str.trim().is_empty() {
                Some(str.parse::<i32>())
            } else {
                None
            }
        })
        .flatten()
        .collect::<Vec<_>>();

    let available_numbers = split
        .next()
        .unwrap()
        .split(' ')
        .filter_map(|str| {
            if !str.trim().is_empty() {
                Some(str.parse::<i32>())
            } else {
                None
            }
        })
        .flatten()
        .collect::<Vec<_>>();

    (winning_numbers, available_numbers)
}

fn overlapping_count(winning: &[i32], available: &[i32]) -> u32 {
    available.iter().fold(0, |acc, val| {
        if winning.iter().any(|v| v == val) {
            acc + 1
        } else {
            acc
        }
    })
}

fn part_one(input: &str) -> i32 {
    input
        .lines()
        .map(|line| {
            let (winning_numbers, available_numbers) = parse_numbers(line);
            let matching_count = overlapping_count(&winning_numbers, &available_numbers);

            if matching_count > 0 {
                2_i32.pow(matching_count - 1)
            } else {
                0
            }
        })
        .sum()
}

fn part_two(input: &str) -> i32 {
    let mut points = vec![1; input.lines().count()];

    input.lines().enumerate().for_each(|(i, line)| {
        let (winning_numbers, available_numbers) = parse_numbers(line);
        let matching_count = overlapping_count(&winning_numbers, &available_numbers);

        for offset in 0..matching_count {
            points[i + 1 + offset as usize] += points[i];
        }
    });

    points.iter().sum()
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
        let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";
        assert_eq!(part_one(input), 13);
    }

    #[test]
    fn part_two_example() {
        let input = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";
        assert_eq!(part_two(input), 30);
    }
}
