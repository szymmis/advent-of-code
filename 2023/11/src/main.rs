#[allow(clippy::needless_range_loop)]

fn calculate_total_distance(input: &str, factor: i64) -> i64 {
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let (width, height) = (grid[0].len(), grid.len());

    let mut galaxies = Vec::new();
    let mut empty_rows = Vec::new();
    let mut empty_cols = Vec::new();

    'outer: for x in 0..width {
        for y in 0..height {
            if grid[y][x] == '#' {
                continue 'outer;
            }
        }

        empty_cols.push(x as i64)
    }

    'outer: for y in 0..height {
        for x in 0..width {
            if grid[y][x] == '#' {
                continue 'outer;
            }
        }

        empty_rows.push(y as i64)
    }

    println!("Empty rows: {:?}", empty_rows);
    println!("Empty cols: {:?}", empty_cols);

    for y in 0..height {
        for x in 0..width {
            if grid[y][x] == '#' {
                galaxies.push((x as i64, y as i64));
            }
        }
    }

    for galaxy in galaxies.iter_mut() {
        let mut x_shift = 0;
        let mut y_shift = 0;

        for c in empty_cols.iter() {
            if *c < galaxy.0 {
                x_shift += factor - 1;
            }
        }
        for r in empty_rows.iter() {
            if *r < galaxy.1 {
                y_shift += factor - 1;
            }
        }

        galaxy.0 += x_shift;
        galaxy.1 += y_shift;
    }

    let mut total_distance = 0;

    for i in 0..galaxies.len() {
        for j in i..galaxies.len() {
            if i != j {
                let a = galaxies[i];
                let b = galaxies[j];
                let distance = (b.0 - a.0).abs() + (b.1 - a.1).abs();

                total_distance += distance;
            }
        }
    }

    total_distance
}

fn main() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    println!("Part one: {}", calculate_total_distance(&input, 2));
    println!("Part two: {}", calculate_total_distance(&input, 1000000))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....";
        assert_eq!(calculate_total_distance(input, 2), 374);
    }

    #[test]
    fn part_two_example() {
        let input = "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....";
        assert_eq!(calculate_total_distance(input, 10), 1030);
        assert_eq!(calculate_total_distance(input, 100), 8410);
    }
}
