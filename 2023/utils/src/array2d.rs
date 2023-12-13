use std::fmt::{Debug, Display};

pub struct Array2D<T> {
    points: Vec<Vec<T>>,
}

impl<T> Array2D<T> {
    pub fn get(&self, x: isize, y: isize) -> Option<&T> {
        match (x, y) {
            (x, _) if x < 0 => None,
            (x, _) if x >= self.width() as isize => None,
            (_, y) if y < 0 => None,
            (_, y) if y >= self.height() as isize => None,
            (x, y) => Some(self.points.get(y as usize)?.get(x as usize)?),
        }
    }

    pub fn width(&self) -> usize {
        self.points[0].len()
    }

    pub fn height(&self) -> usize {
        self.points.len()
    }

    pub fn col(&self, x: isize) -> Vec<&T> {
        (0..self.height())
            .map(|y| self.get(x, y as isize).unwrap())
            .collect()
    }

    pub fn row(&self, y: isize) -> Vec<&T> {
        (0..self.width())
            .map(|x| self.get(x as isize, y).unwrap())
            .collect()
    }
}

impl From<&str> for Array2D<char> {
    fn from(value: &str) -> Self {
        Self {
            points: value
                .lines()
                .map(|line| line.trim().chars().collect())
                .collect(),
        }
    }
}

impl<T> Debug for Array2D<T>
where
    T: Display,
{
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for y in 0..self.height() {
            for x in 0..self.width() {
                write!(f, "{}", self.points[y][x])?;
            }

            if y < self.height() - 1 {
                writeln!(f)?;
            }
        }

        Ok(())
    }
}
