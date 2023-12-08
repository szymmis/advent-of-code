use std::collections::HashMap;

use regex::Regex;

extern crate regex;

#[derive(Debug)]
struct Node {
    name: String,
    left: String,
    right: String,
}

fn collect_nodes(input: &str) -> HashMap<&str, Node> {
    let re = Regex::new(r"(\w+) = \((\w+), (\w+)\)").unwrap();

    input
        .lines()
        .skip(2)
        .map(|line| {
            let (_, [name, left, right]) = re.captures(line).unwrap().extract();
            (
                name,
                Node {
                    name: name.into(),
                    left: left.into(),
                    right: right.into(),
                },
            )
        })
        .collect()
}

fn count_steps(
    instructions: &str,
    nodes: &HashMap<&str, Node>,
    start_node: &str,
    check_fn: fn(&str) -> bool,
) -> i32 {
    let mut ptr = start_node;
    let mut steps = 0;

    loop {
        for instruction in instructions.chars() {
            match instruction {
                'L' => ptr = &nodes.get(ptr).unwrap().left,
                'R' => ptr = &nodes.get(ptr).unwrap().right,
                _ => panic!("Unknown command"),
            }

            steps += 1;

            if check_fn(ptr) {
                return steps;
            }
        }
    }
}

fn calculate_lcm(values: &[i32]) -> u64 {
    fn gcd(mut a: u64, mut b: u64) -> u64 {
        if b > a {
            return gcd(b, a);
        }

        loop {
            (a, b) = (b, a % b);
            if b == 0 {
                return a;
            }
        }
    }

    values
        .iter()
        .map(|v| *v as u64)
        .reduce(|acc, val| acc * val / gcd(acc, val))
        .unwrap()
}

fn part_one(input: &str) -> i32 {
    let instructions = input.lines().next().unwrap();
    let nodes = collect_nodes(input);

    count_steps(instructions, &nodes, "AAA", |name| name == "ZZZ")
}

fn part_two(input: &str) -> u64 {
    let instructions = input.lines().next().unwrap();
    let nodes = collect_nodes(input);

    let paths: Vec<String> = nodes
        .iter()
        .filter_map(|(key, value)| {
            if key.ends_with('A') {
                Some(value.name.clone())
            } else {
                None
            }
        })
        .collect();

    let values: Vec<i32> = paths
        .iter()
        .map(|path| count_steps(instructions, &nodes, path, |name| name.ends_with('Z')))
        .collect();

    println!("{:?}", values);

    calculate_lcm(&values)
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
        let input = "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)";
        assert_eq!(part_one(input), 6);
    }

    #[test]
    fn part_two_example() {
        let input = "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)";
        assert_eq!(part_two(input), 6);
    }
}
