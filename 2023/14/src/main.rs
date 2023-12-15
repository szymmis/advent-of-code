use utils::array2d::Array2D;

#[derive(Debug)]
struct Stop {
    x: usize,
    y: usize,
    count: i32,
}

fn part_one(input: &str) -> i32 {
    let grid: Array2D<char> = input.into();

    let mut stops = Vec::new();

    for x in 0..grid.width() {
        let mut count = 0;
        for (y, shape) in grid.col(x as isize).iter().enumerate().rev() {
            match shape {
                '.' => (),
                'O' => count += 1,
                '#' => {
                    stops.push(Stop { x, y: y + 1, count });
                    count = 0;
                }
                _ => panic!("Unknown shape '{shape}'"),
            }
        }
        if count > 0 {
            stops.push(Stop { x, y: 0, count })
        }
    }

    // println!("{:?}", stops);

    stops
        .iter()
        .map(|Stop { x: _, y, count }| {
            ((grid.height() - y - *count as usize + 1)..=(grid.height() - y)).sum::<usize>() as i32
        })
        .sum()
}

// <<FN_2>>

fn main() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    println!("Part one: {}", part_one(&input));
    // <<PRINT_2>>
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#....";
        assert_eq!(part_one(input), 136);
    }

    // <<TEST_2>>
}
