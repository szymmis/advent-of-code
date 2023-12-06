extern crate regex;

use regex::Regex;

#[derive(Debug)]
struct Range {
    from: u64,
    to: u64,
    length: u64,
}

impl Range {
    fn contains(&self, val: u64) -> Option<u64> {
        if val >= self.from && val < self.from + self.length {
            Some(self.to + (val - self.from))
        } else {
            None
        }
    }
}

fn extract_range_from_slice(input: &str, label: &str) -> Vec<Range> {
    let re = Regex::new(&format!(r"{}:([\d \n]+)", label)).unwrap();
    let numbers = re.captures(input).unwrap().get(1).unwrap().as_str();
    let re = Regex::new(r"(\d+) (\d+) (\d+)").unwrap();
    numbers
        .trim()
        .lines()
        .map(|line| {
            let (_, [to, from, length]) = re.captures(line).unwrap().extract();
            Range {
                to: to.parse().unwrap(),
                from: from.parse().unwrap(),
                length: length.parse().unwrap(),
            }
        })
        .collect()
}

fn value_range_mapping(ranges: &Vec<Range>, value: u64) -> u64 {
    for range in ranges {
        if let Some(val) = range.contains(value) {
            return val;
        }
    }

    value
}

fn parse_seeds(input: &str) -> Vec<u64> {
    let re = Regex::new(r"seeds:([\d \n]+)").unwrap();
    let seeds = re.captures(input).unwrap().get(1).unwrap().as_str();
    seeds
        .trim()
        .split(' ')
        .flat_map(str::parse::<u64>)
        .collect()
}

fn parse_seed_ranges(input: &str) -> Vec<(u64, u64)> {
    let re = Regex::new(r"seeds:([\d \n]+)").unwrap();
    let seeds = re.captures(input).unwrap().get(1).unwrap().as_str();
    let re = Regex::new(r"(\d+) (\d+)").unwrap();

    re.captures_iter(seeds)
        .map(|c| {
            let (_, [start, length]) = c.extract();
            (start.parse().unwrap(), length.parse().unwrap())
        })
        .collect()
}

fn parse_almanac(input: &str) -> Vec<Vec<Range>> {
    let seed_to_soil = extract_range_from_slice(input, "seed-to-soil map");
    let soil_to_fertilizer = extract_range_from_slice(input, "soil-to-fertilizer map");
    let fertilizer_to_water = extract_range_from_slice(input, "fertilizer-to-water map");
    let water_to_light = extract_range_from_slice(input, "water-to-light map");
    let light_to_temperature = extract_range_from_slice(input, "light-to-temperature map");
    let temperature_to_humidity = extract_range_from_slice(input, "temperature-to-humidity map");
    let humidity_to_location = extract_range_from_slice(input, "humidity-to-location map");

    vec![
        seed_to_soil,
        soil_to_fertilizer,
        fertilizer_to_water,
        water_to_light,
        light_to_temperature,
        temperature_to_humidity,
        humidity_to_location,
    ]
}

fn part_one(input: &str) -> u64 {
    let seeds = parse_seeds(input);
    println!("{:?}", seeds);
    let almanac = parse_almanac(input);

    let mut values = Vec::new();

    for seed in seeds {
        let mut transformed_value = seed;
        for step in &almanac {
            transformed_value = value_range_mapping(step, transformed_value);
        }
        values.push(transformed_value);
    }

    *values.iter().min().unwrap()
}

fn part_two(input: &str) -> u64 {
    let seed_ranges = parse_seed_ranges(input);
    println!("{:?}", seed_ranges);
    let almanac = parse_almanac(input);

    let mut lowest_value = u64::MAX;

    for seed_range in seed_ranges {
        for seed in seed_range.0..(seed_range.0 + seed_range.1) {
            let mut transformed_value = seed;
            for step in &almanac {
                transformed_value = value_range_mapping(step, transformed_value);
            }

            if transformed_value < lowest_value {
                lowest_value = transformed_value;
            }
        }
    }

    lowest_value
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
        let input = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4";
        assert_eq!(part_one(input), 35);
    }

    #[test]
    fn part_two_example() {
        let input = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4";
        assert_eq!(part_two(input), 46);
    }
}
