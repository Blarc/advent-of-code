use std::collections::HashSet;

advent_of_code::solution!(6);

const DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];

fn parse(input: &str) -> (Vec<Vec<char>>, (usize, usize, usize)) {
    let mut guard = (0, 0, 0);
    let grid = input
        .lines()
        .enumerate()
        .map(|(y, line): (usize, &str)| {
            line.chars()
                .enumerate()
                .map(|(x, c): (usize, char)| {
                    if c.eq(&'^') {
                        guard = (y, x, 0)
                    }
                    c
                })
                .collect()
        })
        .collect();
    (grid, guard)
}

pub fn is_loop(mut guard: (usize, usize, usize), grid: &Vec<Vec<char>>) -> (bool, HashSet<(usize, usize)>) {
    let mut pos = HashSet::new();
    let mut pos_with_dir = HashSet::new();
    loop {
        pos.insert((guard.0, guard.1));
        let next = ((guard.0 as i32 + DIRS[guard.2].0) as usize, (guard.1 as i32 + DIRS[guard.2].1) as usize, guard.2);
        if next.0 < grid.len() && next.1 < grid.len() {

            if pos_with_dir.contains(&next) {
                return (true, pos)
            }

            if grid[next.0][next.1].eq(&'#') {
                pos_with_dir.insert(guard);
                guard.2 += 1;
                guard.2 %= 4;
            } else {
                guard = next;
            }
        } else {
            return (false, pos);
        }
    }
}

pub fn part_one(input: &str) -> Option<u32> {
    let (grid, guard) = parse(input);
    let (_, guard_path) = is_loop(guard, &grid);
    Some(guard_path.len() as u32)
}

pub fn part_two(input: &str) -> Option<u32> {
    let (mut grid, guard) = parse(input);
    let (_, guard_path) = is_loop(guard, &grid);

    let mut result = 0;
    for tile in guard_path {
        if tile.0 == guard.0 && tile.1 == guard.1 {
            continue
        }

        grid[tile.0][tile.1] = '#';
        let (is_loop, _) = is_loop(guard, &grid);
        if is_loop {
            result += 1;
        }
        grid[tile.0][tile.1] = '.';

    }

    // 2002 too high
    Some(result)
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
