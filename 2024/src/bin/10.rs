use std::collections::{HashMap, HashSet};

advent_of_code::solution!(10);

const DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];

type Coord = (i32, i32);
fn create_grid(input: &str) -> HashMap<Coord, usize> {
    input
        .lines()
        .enumerate()
        .flat_map(|(y, line)| {
            line.chars().enumerate().filter(|(_, c)| c.is_digit(10)).map(move |(x, c)| {
                (
                    (y as i32, x as i32),
                    c.to_digit(10).expect("Char is not usize.") as usize,
                )
            })
        })
        .collect()
}

fn filter_starts(grid: &HashMap<Coord, usize>) -> Vec<Coord> {
    let starts: Vec<Coord> = grid
        .iter()
        .filter(|(_, &h)| h == 0)
        .map(|(&c, _)| c)
        .collect();
    starts
}

fn find_trail<F>(pos: Coord, grid: &HashMap<Coord, usize>, add_to_visited: &mut F) where F: FnMut((i32, i32))
{
    let pos_value = grid.get(&pos).unwrap();
    if *pos_value == 9 {
        add_to_visited(pos);
    }

    for dir in DIRS {
        let next = (pos.0 + dir.0, pos.1 + dir.1);
        if grid.contains_key(&next) {
            let next_value = grid.get(&next).unwrap();
            if pos_value + 1 == *next_value {
                find_trail(next, &grid, add_to_visited);
            }
        }
    }
}

pub fn part_one(input: &str) -> Option<u32> {
    let grid = create_grid(input);
    let starts = filter_starts(&grid);

    let sum = starts.iter().map(|&x| {
        let mut result = HashSet::new();
        find_trail(x, &grid, &mut |c| { result.insert(c); });
        result.len() as u32
    }).sum();

    Some(sum)
}

pub fn part_two(input: &str) -> Option<u32> {
    let grid = create_grid(input);
    let starts = filter_starts(&grid);

    let sum = starts.iter().map(|&x| {
        let mut result = Vec::new();
        find_trail(x, &grid, &mut |c| { result.push(c); });
        result.len() as u32
    }).sum();

    Some(sum)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(36));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(81));
    }
}
