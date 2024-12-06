use std::collections::{HashMap, HashSet};

advent_of_code::solution!(6);

const DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];

fn create_grid(input: &str) -> HashMap<(i32, i32), char> {
    input
        .lines()
        .enumerate()
        .flat_map(|(y, line)| {
            line.chars()
                .enumerate()
                .map(move |(x, c)| ((y as i32, x as i32), c))
        })
        .collect()
}

fn find_guard(grid: &HashMap<(i32, i32), char>) -> (i32, i32) {
    let guard = grid
        .iter()
        .find_map(|(&pos, &c)| if c.eq(&'^') { Some(pos) } else { None })
        .expect("Expected at least one '^' character in the grid");
    guard
}

fn walk(
    mut pos: (i32, i32),
    grid: &HashMap<(i32, i32), char>,
) -> (HashSet<(i32, i32)>, bool) {
    let mut visited = HashSet::new();
    let mut dir = 0;
    while grid.contains_key(&pos) && !visited.contains(&(pos, dir)) {
        visited.insert((pos, dir));
        let next = (pos.0 + DIRS[dir].0, pos.1 + DIRS[dir].1);
        if grid.get(&next) == Some(&'#') {
            dir = (dir + 1) % 4
        } else {
            pos = next
        }
    }

    let is_loop = visited.contains(&(pos, dir));
    // Remove duplicates that were created because of different dir
    let path: HashSet<(i32, i32)> = visited.iter().map(|&x| x.0).collect();
    (path, is_loop)
}

pub fn part_one(input: &str) -> Option<u32> {
    let grid = create_grid(input);
    let guard = find_guard(&grid);
    let (path, _) = walk(guard, &grid);
    Some(path.len() as u32)
}

pub fn part_two(input: &str) -> Option<u32> {
    let mut grid = create_grid(input);
    let guard = find_guard(&grid);

    let (path, _) = walk(guard, &grid);
    let result = path
        .iter()
        .filter(|&&pos| pos != guard)
        .filter(|&&pos| {
            grid.insert(pos, '#');
            let (_, is_loop) = walk(guard, &grid);
            grid.insert(pos, '.');
            is_loop
        })
        .count();

    Some(result as u32)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(41));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(6));
    }
}
