use std::fs;

fn get_sum(input: &str, parse_words: bool) -> i32 {
    let get_digit = |line: &str, (i, char): (usize, char)| -> Option<u32> {
        if char.is_ascii_digit() {
            Some(char.to_digit(10).expect("Char is an ascii digit"))
        } else if parse_words {
            let line = &line[i..];

            if line.starts_with("one") {
                Some(1)
            } else if line.starts_with("two") {
                Some(2)
            } else if line.starts_with("three") {
                Some(3)
            } else if line.starts_with("four") {
                Some(4)
            } else if line.starts_with("five") {
                Some(5)
            } else if line.starts_with("six") {
                Some(6)
            } else if line.starts_with("seven") {
                Some(7)
            } else if line.starts_with("eight") {
                Some(8)
            } else if line.starts_with("nine") {
                Some(9)
            } else {
                None
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

fn main() {
    let file_input = fs::read_to_string("day_01/src/input.txt").unwrap();

    println!("{}", get_sum(&file_input, false));
    println!("{}", get_sum(&file_input, true));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = r"1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

        assert_eq!(get_sum(input, false), 142)
    }

    #[test]
    fn part_two_example() {
        let input = r"two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

        assert_eq!(get_sum(input, true), 281)
    }
}
