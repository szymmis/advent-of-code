use std::collections::HashMap;

use regex::Regex;

#[derive(Debug)]
struct Step {
    label: String,
    op: Op,
}

#[derive(Debug)]
enum Op {
    Add(i32),
    Remove,
}

#[derive(Debug)]
struct Lens(String, i32);

impl From<(&str, &str, &str)> for Step {
    fn from((label, operator, value): (&str, &str, &str)) -> Self {
        Self {
            label: label.to_string(),
            op: match operator {
                "=" => Op::Add(value.parse().unwrap()),
                "-" => Op::Remove,
                _ => panic!("Unknown operator '{}'", operator),
            },
        }
    }
}

fn hash(input: &str) -> i32 {
    let mut curr_val = 0;
    for char in input.chars() {
        curr_val += char as u32;
        curr_val *= 17;
        curr_val %= 256;
    }
    curr_val as i32
}

fn part_one(input: &str) -> i32 {
    input.trim().split(',').map(hash).sum()
}

fn part_two(input: &str) -> i32 {
    let steps: Vec<Step> = input
        .trim()
        .split(',')
        .map(|pattern| {
            let re = Regex::new(r"(\w+)(=|-)(\d*)").unwrap();
            let (_, [label, operator, value]) = re.captures(pattern).unwrap().extract();
            (label, operator, value).into()
        })
        .collect();

    let mut map: HashMap<i32, Vec<Lens>> = HashMap::new();

    for Step { label, op } in steps {
        let vec = map.entry(hash(&label)).or_default();

        match op {
            Op::Add(value) => {
                if let Some(pos) = vec.iter().position(|lens| lens.0 == label) {
                    vec[pos].1 = value;
                } else {
                    vec.push(Lens(label, value))
                }
            }
            Op::Remove => {
                if let Some(pos) = vec.iter().position(|lens| lens.0 == label) {
                    vec.remove(pos);
                }
            }
        }
    }

    let sum: i32 = map
        .values()
        .map(|vec| {
            vec.iter().enumerate().fold(0, |acc, (i, lens)| {
                acc + ((hash(&lens.0) + 1) * (i as i32 + 1) * lens.1)
            })
        })
        .sum();

    sum
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
        let input = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7";
        assert_eq!(part_one(input), 1320);
    }

    #[test]
    fn part_two_example() {
        let input = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7";
        assert_eq!(part_two(input), 145);
    }
}
