use itertools::Itertools;

extern crate itertools;

fn matches(input: &str, pattern: &str) -> bool {
    if pattern.len() > input.len() {
        return false;
    }

    for (pattern_char, input_char) in pattern.chars().zip(input.chars()) {
        match input_char {
            '?' => (),
            '.' => {
                if pattern_char == '#' {
                    return false;
                }
            }
            '#' => {
                if pattern_char == '.' {
                    return false;
                }
            }
            _ => panic!("Unknown char '{}' in input", input_char),
        }
    }

    true
}

fn perms(input: &str, nums: &[i64]) -> i64 {
    if nums.len() == 1 && input.chars().all(|c| c == '?') {
        return input.len() as i64 - nums[0] + 1;
    }

    let mut count = 0;
    let value = nums[0];
    let tail = &nums[1..];
    let sum = (tail.iter().fold(0, |acc, v| acc + (v + 1)) - 1).max(0);

    let pattern = |offset: usize| -> String {
        if nums.len() > 1 {
            format!("{}{}.", ".".repeat(offset), "#".repeat(value as usize))
        } else {
            format!(
                "{}{}{}",
                ".".repeat(offset),
                "#".repeat(value as usize),
                ".".repeat(if input.len() > offset + value as usize {
                    input.len() - offset - value as usize
                } else {
                    0
                })
            )
        }
    };

    let len_left =
        |offset: usize| -> isize { input.len() as isize - pattern(offset).len() as isize };

    let mut offset = 0;
    while len_left(offset) >= sum as isize {
        if matches(input, &pattern(offset)) {
            match nums.len() {
                1 => count += 1,
                _ => count += perms(&input[(offset + value as usize + 1)..], tail),
            }
        }
        offset += 1;
    }

    count
}

fn part_one(input: &str) -> i64 {
    input
        .lines()
        .map(|line| {
            let [pattern, sizes] = line.split(' ').collect::<Vec<_>>()[..] else {
                panic!()
            };
            let sizes: Vec<i64> = sizes.split(',').flat_map(str::parse).collect();
            let perms = perms(pattern, &sizes[..]);

            println!("{} {:?} -> {}", pattern, sizes, perms);

            perms
        })
        .sum()
}

fn part_two(input: &str) -> i64 {
    input
        .lines()
        .map(|line| {
            let [pattern, sizes] = line.split(' ').collect::<Vec<_>>()[..] else {
                panic!()
            };
            let sizes: Vec<i64> = sizes.split(',').flat_map(str::parse).collect();
            let perms = perms(&(0..5).map(|_| pattern).join("?"), &(sizes.repeat(5))[..]);

            println!("{} {:?} -> {}", pattern, sizes, perms);

            perms
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
        let input = "???.### 1,1,3\n.??..??...?##. 1,1,3\n?#?#?#?#?#?#?#? 1,3,1,6\n????.#...#... 4,1,1\n????.######..#####. 1,6,5\n?###???????? 3,2,1";
        assert_eq!(part_one(input), 21);
    }

    #[test]
    fn part_two_example() {
        let input = "???.### 1,1,3\n.??..??...?##. 1,1,3\n?#?#?#?#?#?#?#? 1,3,1,6\n????.#...#... 4,1,1\n????.######..#####. 1,6,5\n?###???????? 3,2,1";
        assert_eq!(part_two(input), 525152);
    }
}
