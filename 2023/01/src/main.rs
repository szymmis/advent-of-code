fn get_sum(input: &str, parse_words: bool) -> i32 {
    let get_digit = |line: &str, (i, char): (usize, char)| -> Option<u32> {
        if char.is_ascii_digit() {
            Some(char.to_digit(10).expect("Char is an ascii digit"))
        } else if parse_words {
            let line = &line[i..];

            match line {
                line if line.starts_with("one") => Some(1),
                line if line.starts_with("two") => Some(2),
                line if line.starts_with("three") => Some(3),
                line if line.starts_with("four") => Some(4),
                line if line.starts_with("five") => Some(5),
                line if line.starts_with("six") => Some(6),
                line if line.starts_with("seven") => Some(7),
                line if line.starts_with("eight") => Some(8),
                line if line.starts_with("nine") => Some(9),
                _ => None,
            }
        } else {
            None
        }
    };

    let mut vec = Vec::new();

    for line in input.lines() {
        let mut first_digit: Option<u32> = None;
        let mut last_digit: Option<u32> = None;

        for enumeration in line.chars().enumerate() {
            let digit = get_digit(line, enumeration);
            if digit.is_some() {
                last_digit = digit;
                if first_digit.is_none() {
                    first_digit = digit;
                }
            }
        }

        vec.push((first_digit.unwrap() * 10 + last_digit.unwrap()) as i32);
    }

    vec.iter().sum()
}

fn part_one(input: &str) -> i32 {
    get_sum(input, false)
}

fn part_two(input: &str) -> i32 {
    get_sum(input, true)
}

fn main() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    println!("Part one: {}", part_one(&input));
    println!("Part two: {}", part_two(&input));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet";
        assert_eq!(part_one(input), 142);
    }

    #[test]
    fn part_two_example() {
        let input = "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen";
        assert_eq!(part_two(input), 281);
    }
}
