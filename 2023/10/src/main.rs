enum ComingFrom {
    Left,
    Right,
    Up,
    Down,
}

fn next_step(char: char, (x, y): (usize, usize), from: ComingFrom) -> ((usize, usize), ComingFrom) {
    match char {
        '|' => match from {
            ComingFrom::Up => ((x, y + 1), from),
            ComingFrom::Down => ((x, y - 1), from),
            _ => panic!(),
        },
        '-' => match from {
            ComingFrom::Left => ((x + 1, y), from),
            ComingFrom::Right => ((x - 1, y), from),
            _ => panic!(),
        },
        'L' => match from {
            ComingFrom::Up => ((x + 1, y), ComingFrom::Left),
            ComingFrom::Right => ((x, y - 1), ComingFrom::Down),
            _ => panic!(),
        },
        'J' => match from {
            ComingFrom::Up => ((x - 1, y), ComingFrom::Right),
            ComingFrom::Left => ((x, y - 1), ComingFrom::Down),
            _ => panic!(),
        },
        '7' => match from {
            ComingFrom::Down => ((x - 1, y), ComingFrom::Right),
            ComingFrom::Left => ((x, y + 1), ComingFrom::Up),
            _ => panic!(),
        },
        'F' => match from {
            ComingFrom::Down => ((x + 1, y), ComingFrom::Left),
            ComingFrom::Right => ((x, y + 1), ComingFrom::Up),
            _ => panic!(),
        },
        _ => panic!("Unknown char {}", char),
    }
}

fn part_one(input: &str) -> i32 {
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let start_pos = {
        let index = input.find('S').unwrap();
        (index % (grid[0].len() + 1), index / (grid[0].len() + 1))
    };
    println!("Start position: {:?}", start_pos);

    let mut step = {
        let (x, y) = (start_pos.0, start_pos.1);

        if x > 0 && ['-', 'L', 'F'].contains(&grid[y][x - 1]) {
            ((x - 1, y), ComingFrom::Right)
        } else if ['-', 'J', '7'].contains(&grid[y][x + 1]) {
            ((x + 1, y), ComingFrom::Left)
        } else if y > 0 && ['|', 'F', '7'].contains(&grid[y - 1][x]) {
            ((x, y - 1), ComingFrom::Down)
        } else if ['|', 'J', 'L'].contains(&grid[y - 1][x]) {
            ((x, y + 1), ComingFrom::Up)
        } else {
            panic!("Cannot find starting point connection")
        }
    };

    let mut visited_nodes: Vec<(usize, usize)> = Vec::new();

    loop {
        let ((x, y), from) = step;
        let char = grid[y][x];

        visited_nodes.push((x, y));

        match char {
            'S' => break,
            char => step = next_step(char, (x, y), from),
        }
    }

    (visited_nodes.len() / 2) as i32
}

fn main() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    println!("Part one: {}", part_one(&input));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = ".....\n.S-7.\n.|.|.\n.L-J.\n.....";
        assert_eq!(part_one(input), 4);
        let input = "..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...";
        assert_eq!(part_one(input), 8);
    }
}
