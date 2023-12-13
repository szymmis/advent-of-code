use utils::array2d::Array2D;

fn part_one(input: &str) -> i32 {
    let v = input
        .split("\n\n")
        .map(|pattern| {
            let grid: Array2D<char> = pattern.into();

            'outer: for i in 1..grid.width() {
                for j in 0..(i.min(grid.width() - i)) {
                    if grid.col((i - 1 - j) as isize) != grid.col((i + j) as isize) {
                        continue 'outer;
                    }
                }

                return i as i32;
            }

            'outer: for i in 1..grid.height() {
                for j in 0..(i.min(grid.height() - i)) {
                    if grid.row((i - 1 - j) as isize) != grid.row((i + j) as isize) {
                        continue 'outer;
                    }
                }

                return i as i32 * 100;
            }

            panic!("Could not find a reflection")
        })
        .collect::<Vec<_>>();

    println!("{:?}", v);

    v.iter().sum()
}

fn part_two(input: &str) -> i32 {
    fn vec_diff<T>(a: &[T], b: &[T]) -> usize
    where
        T: PartialEq,
    {
        a.iter().zip(b.iter()).filter(|(a, b)| a != b).count()
    }

    let v = input
        .split("\n\n")
        .map(|pattern| {
            let grid: Array2D<char> = pattern.into();

            'outer: for i in 1..grid.width() {
                let mut was_different = false;

                for j in 0..(i.min(grid.width() - i)) {
                    match vec_diff(&grid.col((i - 1 - j) as isize), &grid.col((i + j) as isize)) {
                        0 => continue,
                        1 if !was_different => {
                            was_different = true;
                            continue;
                        }
                        _ => continue 'outer,
                    }
                }

                if was_different {
                    return i as i32;
                }
            }

            'outer: for i in 1..grid.height() {
                let mut was_different = false;

                for j in 0..(i.min(grid.height() - i)) {
                    match vec_diff(&grid.row((i - 1 - j) as isize), &grid.row((i + j) as isize)) {
                        0 => continue,
                        1 if !was_different => {
                            was_different = true;
                            continue;
                        }
                        _ => continue 'outer,
                    }
                }

                if was_different {
                    return i as i32 * 100;
                }
            }

            panic!("Could not find a reflection")
        })
        .collect::<Vec<_>>();

    println!("{:?}", v);

    v.iter().sum()
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
        let input =
            "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#";
        assert_eq!(part_one(input), 405);
    }

    #[test]
    fn part_two_example() {
        let input = "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#";
        assert_eq!(part_two(input), 400);
    }
}
