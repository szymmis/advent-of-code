extern crate regex;

use regex::Regex;

fn part_one(input: &str) -> i32 {
    let mut lines = input.lines();
    let re = Regex::new(r"\d+").unwrap();

    let times = re
        .find_iter(lines.next().unwrap())
        .flat_map(|m| m.as_str().parse::<f64>());
    let distances = re
        .find_iter(lines.next().unwrap())
        .flat_map(|m| m.as_str().parse::<f64>());

    let races = times.zip(distances).collect::<Vec<_>>();

    let mut total_product = 1;

    for (time, distance) in races {
        println!("Time: {}, Distance: {}", time, distance);

        let delta: f64 = time * time - 4.0 * (distance + 1.0);
        let x1: f64 = ((-time + delta.sqrt()) / -2.0).ceil();
        let x2: f64 = ((-time - delta.sqrt()) / -2.0).floor();
        let count = (x2 - x1) as i64 + 1;

        println!("x1 = {}, x2 = {}, x2 - x1 = {}", x1, x2, count);

        total_product *= count;
    }

    total_product as i32
}

fn part_two(input: &str) -> i64 {
    let mut lines = input.lines();
    let re = Regex::new(r"\d+").unwrap();

    let time: f64 = re
        .find_iter(lines.next().unwrap())
        .map(|m| m.as_str())
        .collect::<String>()
        .parse()
        .unwrap();
    let distance: f64 = re
        .find_iter(lines.next().unwrap())
        .map(|m| m.as_str())
        .collect::<String>()
        .parse()
        .unwrap();

    println!("Time: {}, Distance: {}", time, distance);

    let delta: f64 = time * time - 4.0 * (distance + 1.0);
    let x1: f64 = ((-time + delta.sqrt()) / -2.0).ceil();
    let x2: f64 = ((-time - delta.sqrt()) / -2.0).floor();
    let count = (x2 - x1) as i64 + 1;

    println!("x1 = {}, x2 = {}, x2 - x1 = {}", x1, x2, count);

    count
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
        let input = "Time:      7  15   30\nDistance:  9  40  200";
        assert_eq!(part_one(input), 288);
    }

    #[test]
    fn part_two_example() {
        let input = "Time:      7  15   30\nDistance:  9  40  200";
        assert_eq!(part_two(input), 71503);
    }
}
