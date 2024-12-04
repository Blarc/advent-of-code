use std::collections::HashMap;
use std::ops::Add;

advent_of_code::solution!(4);

const XMAS: &[char] = &['X', 'M', 'A', 'S'];
const MAS: &[char] = &['M', 'A', 'S'];

pub fn part_one(input: &str) -> Option<u32> {
    let mut grid: Vec<Vec<char>> = Vec::new();
    for line in input.lines() {
        let row = line.chars().collect();
        grid.push(row);
    }

    // The input is square
    let mut result: u32 = 0;
    let size = grid.len();
    for y in 0..size {
        for x in 0..size {
            // The word starts with X
            if grid[y][x] != 'X' {
                continue
            }

            // Compute all possible directions
            for j in -1..=1 {
                'x_dir: for i in -1..=1 {

                    // Direction (0,0) is not valid
                    if j == 0 && i == 0 {
                        continue
                    }

                    let mut pos = (y as isize, x as isize);
                    for step in 1..=3 {
                        pos.0 = pos.0 + j;
                        pos.1 = pos.1 + i;
                        if 0 <= pos.0 && pos.0 < size as isize && 0 <= pos.1 && pos.1 < size as isize {
                            if grid[pos.0 as usize][pos.1 as usize] != XMAS[step] {
                                continue 'x_dir
                            }
                        } else {
                            continue 'x_dir
                        }
                    }
                    result += 1;
                }
            }
        }
    }


    Some(result)
}

pub fn part_two(input: &str) -> Option<u32> {
    let mut grid: Vec<Vec<char>> = Vec::new();
    for line in input.lines() {
        let row = line.chars().collect();
        grid.push(row);
    }

    // The input is square
    let mut cross = HashMap::new();

    let size = grid.len();
    for y in 0..size {
        for x in 0..size {
            // The word starts with X
            if grid[y][x] != 'M' {
                continue
            }

            // Compute all possible directions
            for j in vec![-1, 1] {
                'x_dir: for i in vec![-1, 1] {
                    let mut pos = (y as isize, x as isize);
                    for step in 1..=2 {
                        pos.0 = pos.0 + j;
                        pos.1 = pos.1 + i;
                        if 0 <= pos.0 && pos.0 < size as isize && 0 <= pos.1 && pos.1 < size as isize {
                            if grid[pos.0 as usize][pos.1 as usize] != MAS[step] {
                                continue 'x_dir
                            }
                        } else {
                            continue 'x_dir
                        }
                    }

                    // One step back
                    pos.0 = pos.0 - j;
                    pos.1 = pos.1 - i;

                    cross.entry(pos).and_modify(|x| *x += 1).or_insert(1);
                }
            }
        }
    }

    let result: u32 = cross.values().filter(|&&x| x > 1).count() as u32;

    Some(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(18));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(9));
    }
}
