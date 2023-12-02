extern crate regex;

use regex::Regex;

#[derive(Debug, PartialEq, Clone, Copy)]
enum CubeColor {
    Red,
    Green,
    Blue,
}

#[derive(Debug, Clone, Copy)]
struct Cube(CubeColor, i32);

impl From<(&str, &str)> for Cube {
    fn from((count, color): (&str, &str)) -> Self {
        let count: i32 = count.parse().unwrap();

        match color {
            "red" => Self(CubeColor::Red, count),
            "green" => Self(CubeColor::Green, count),
            "blue" => Self(CubeColor::Blue, count),
            _ => panic!("Unsupported color {}", color),
        }
    }
}

impl Cube {
    fn is_valid(&self) -> bool {
        match self.0 {
            CubeColor::Red => self.1 <= 12,
            CubeColor::Green => self.1 <= 13,
            CubeColor::Blue => self.1 <= 14,
        }
    }
}

fn part_one(input: &str) -> i32 {
    let game_re = Regex::new(r"Game (?P<id>\d+)").unwrap();
    let cubes_re = Regex::new(r"(\d+) (red|green|blue)").unwrap();

    let mut valid_games: Vec<i32> = Vec::new();

    for line in input.lines() {
        let game_id: i32 = game_re.captures(line).unwrap()["id"].parse().unwrap();

        let cubes: Vec<Cube> = cubes_re
            .captures_iter(line)
            .map(|c| {
                let (_, [color, count]) = c.extract();
                Cube::from((color, count))
            })
            .collect();

        println!("Game {game_id}");
        println!("{:?}", cubes);

        if cubes.iter().all(|cube| cube.is_valid()) {
            valid_games.push(game_id);
        }
    }

    println!("{:?}", valid_games);

    valid_games.iter().sum()
}

fn find_max_cube_of_color(cubes: &[Cube], color: CubeColor) -> &Cube {
    cubes
        .iter()
        .filter(|cube| cube.0 == color)
        .reduce(|acc, cube| {
            dbg!(cube);
            if cube.1 > acc.1 {
                cube
            } else {
                acc
            }
        })
        .unwrap()
}

fn part_two(input: &str) -> i32 {
    let cubes_re = Regex::new(r"(\d+) (red|green|blue)").unwrap();

    let mut powers = Vec::new();

    for line in input.lines() {
        let cubes: Vec<Cube> = cubes_re
            .captures_iter(line)
            .map(|c| {
                let (_, [color, count]) = c.extract();
                Cube::from((color, count))
            })
            .collect();

        let Cube(_, max_r) = find_max_cube_of_color(&cubes, CubeColor::Red);
        let Cube(_, max_g) = find_max_cube_of_color(&cubes, CubeColor::Green);
        let Cube(_, max_b) = find_max_cube_of_color(&cubes, CubeColor::Blue);

        println!("r: {}, g: {}, b: {}", max_r, max_g, max_b);
        let power = max_r * max_g * max_b;
        println!("{power}");

        powers.push(power);
    }

    powers.iter().sum()
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
        let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";
        assert_eq!(part_one(input), 8);
    }

    #[test]
    fn part_two_example() {
        let input = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";
        assert_eq!(part_two(input), 2286);
    }
}
