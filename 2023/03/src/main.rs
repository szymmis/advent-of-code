#![feature(stmt_expr_attributes)]

use std::collections::HashMap;

#[derive(Debug)]
struct Part {
    val: i32,
    symbol: Symbol,
}

#[derive(Debug)]
struct Symbol {
    char: char,
    pos: (usize, usize),
}

fn is_valid_symbol(char: char) -> bool {
    !char.is_ascii_digit() && char != '.'
}

fn is_part_valid(x: usize, y: usize, chars: &[Vec<char>]) -> bool {
    (x > 0 && is_valid_symbol(chars[y][x - 1]))
        || (x < chars[0].len() - 1 && is_valid_symbol(chars[y][x + 1]))
        || (y > 0 && is_valid_symbol(chars[y - 1][x]))
        || (y < chars.len() - 1 && is_valid_symbol(chars[y + 1][x]))
        || (x > 0 && y > 0 && is_valid_symbol(chars[y - 1][x - 1]))
        || (x < chars[0].len() - 1 && y > 0 && is_valid_symbol(chars[y - 1][x + 1]))
        || (x > 0 && y < chars.len() - 1 && is_valid_symbol(chars[y + 1][x - 1]))
        || (x < chars[0].len() - 1 && y < chars.len() - 1 && is_valid_symbol(chars[y + 1][x + 1]))
}

fn get_part_symbol(x: usize, y: usize, chars: &[Vec<char>]) -> Symbol {
    #[rustfmt::skip]
    let (char, pos) = match (x, y) {
        (x, y) if (x > 0 && is_valid_symbol(chars[y][x - 1])) => (chars[y][x - 1], (x - 1, y)),
        (x, y) if (x < chars[0].len() - 1 && is_valid_symbol(chars[y][x + 1])) => (chars[y][x + 1], (x + 1, y)),
        (x, y) if (y > 0 && is_valid_symbol(chars[y - 1][x])) => (chars[y - 1][x], (x, y - 1)),
        (x, y) if (y < chars.len() - 1 && is_valid_symbol(chars[y + 1][x])) => (chars[y + 1][x],(x, y + 1)),
        (x, y) if (x > 0 && y > 0 && is_valid_symbol(chars[y - 1][x - 1])) => (chars[y - 1][x - 1], (x-1,y-1)),
        (x, y) if (x < chars[0].len() - 1 && y > 0 && is_valid_symbol(chars[y - 1][x + 1])) => (chars[y - 1][x + 1], (x+1,y-1)),
        (x, y) if (x > 0 && y < chars.len() - 1 && is_valid_symbol(chars[y + 1][x - 1])) => (chars[y + 1][x - 1], (x-1,y+1)),
        (x, y) if (x < chars[0].len() - 1 && y < chars.len() - 1 && is_valid_symbol(chars[y + 1][x + 1])) => (chars[y + 1][x + 1], (x+1,y+1)),
        _ => panic!("{} {}", x, y),
    };

    Symbol { char, pos }
}

fn part_one(input: &str) -> i32 {
    let chars: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let mut numbers: Vec<i32> = Vec::new();

    for (i, line) in chars.iter().enumerate() {
        let mut num_buffer = String::new();
        let mut is_num_valid = false;

        for (j, char) in line.iter().enumerate() {
            if char.is_ascii_digit() {
                num_buffer.push(*char);
                is_num_valid = is_num_valid || is_part_valid(j, i, &chars);
            } else if !num_buffer.is_empty() {
                if is_num_valid {
                    numbers.push(num_buffer.clone().parse().unwrap());
                }
                num_buffer.clear();
                is_num_valid = false;
            }
        }

        if !num_buffer.is_empty() && is_num_valid {
            numbers.push(num_buffer.clone().parse().unwrap());
        }
    }

    println!("{:?}", &numbers);

    numbers.iter().sum()
}

fn part_two(input: &str) -> i32 {
    let chars: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let mut parts: Vec<Part> = Vec::new();

    for (i, line) in chars.iter().enumerate() {
        let mut num_buffer = String::new();
        let mut symbol: Option<Symbol> = None;

        for (j, char) in line.iter().enumerate() {
            if char.is_ascii_digit() {
                num_buffer.push(*char);
                if is_part_valid(j, i, &chars) {
                    symbol = Some(get_part_symbol(j, i, &chars));
                }
            } else if !num_buffer.is_empty() {
                if let Some(symbol) = symbol {
                    parts.push(Part {
                        val: num_buffer.clone().parse().unwrap(),
                        symbol,
                    });
                }
                num_buffer.clear();
                symbol = None;
            }
        }

        if !num_buffer.is_empty() {
            if let Some(symbol) = symbol {
                parts.push(Part {
                    val: num_buffer.clone().parse().unwrap(),
                    symbol,
                });
            }
        }
    }

    let engine_parts = parts
        .iter()
        .filter(|part| part.symbol.char == '*')
        .collect::<Vec<_>>();

    let mut engine_map: HashMap<(usize, usize), Vec<i32>> = HashMap::new();

    // println!("{:#?}", &engine_parts);

    for part in engine_parts {
        engine_map
            .entry(part.symbol.pos)
            .or_default()
            .push(part.val);
    }

    println!("{:?}", engine_map);

    engine_map
        .values()
        .filter(|v| v.len() > 1)
        .map(|v| v.iter().product::<i32>())
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
        let input = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..";
        assert_eq!(part_one(input), 4361);
    }

    #[test]
    fn part_two_example() {
        let input = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..";
        assert_eq!(part_two(input), 467835);
    }
}
