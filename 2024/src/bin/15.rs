use std::collections::{HashMap, HashSet};

advent_of_code::solution!(15);

type Coord = (i32, i32);
const DIRS: [Coord; 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];

fn create_grid(input: &str) -> HashMap<Coord, char> {
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

fn create_grid_2(input: &str) -> HashMap<Coord, char> {
    let mut grid: HashMap<Coord, char> = HashMap::new();
    for (y, line) in input.lines().enumerate() {
        for (x, char) in line.chars().enumerate() {
            let fixed_y = y as i32;
            let fixed_x = (x * 2) as i32;

            match char {
                '#' => {
                    grid.insert((fixed_y, fixed_x), '#');
                    grid.insert((fixed_y, fixed_x + 1), '#');
                }
                'O' => {
                    grid.insert((fixed_y, fixed_x), '[');
                    grid.insert((fixed_y, fixed_x + 1), ']');
                }
                '.' => {
                    grid.insert((fixed_y, fixed_x), '.');
                    grid.insert((fixed_y, fixed_x + 1), '.');
                }
                '@' => {
                    grid.insert((fixed_y, fixed_x), '@');
                    grid.insert((fixed_y, fixed_x + 1), '.');
                }
                _ => {}
            }
        }
    }

    grid
}

fn parse_moves(input: &str) -> Vec<Coord> {
    input
        .lines()
        .flat_map(|l| {
            l.chars().map(|x| {
                if x == '^' {
                    DIRS[0]
                } else if x == '>' {
                    DIRS[1]
                } else if x == 'v' {
                    DIRS[2]
                } else if x == '<' {
                    DIRS[3]
                } else {
                    panic!("Invalid move {} !", x)
                }
            })
        })
        .collect()
}

fn print_map(grid: &HashMap<Coord, char>, height: usize, width: usize) {
    for y in 0..height {
        for x in 0..width {
            // print!("{:?}", (y as i32, x as i32));
            print!("{}", grid.get(&(y as i32, x as i32)).unwrap());
        }
        println!()
    }
}

fn push(pos: Coord, dir: Coord, grid: &mut HashMap<Coord, char>) -> Coord {
    let next = (pos.0 + dir.0, pos.1 + dir.1);
    match grid.get(&next) {
        Some('#') => pos, // Wall, no movement
        Some('.') => {
            // Move to empty space
            grid.insert(next, grid[&pos]);
            grid.insert(pos, '.');
            next
        }
        Some('O') => {
            // Attempt to push the object forward
            let pushed_to = push(next, dir, grid);
            if pushed_to != next {
                // Successful push of 'O', current `pos` moves to `next`
                grid.insert(next, grid[&pos]);
                grid.insert(pos, '.');
                next
            } else {
                // Push failed, stay in place
                pos
            }
        }
        _ => panic!("Unexpected grid state at {:?}", next),
    }
}

fn check_vertical(pos: Coord, dir: Coord, grid: &HashMap<Coord, char>) -> Option<HashSet<Coord>> {
    let next = (pos.0 + dir.0, pos.1 + dir.1);

    let mut checked_coords = HashSet::new();
    match grid.get(&next) {
        Some('.') => {
            checked_coords.insert(pos);
        }
        Some('#') => return None,
        Some('[') => {
            let coords_left = check_vertical(next, dir, grid);
            if coords_left.is_none() {
                return None;
            }
            checked_coords.extend(coords_left.unwrap());

            if grid[&pos] != '[' {
                let coords_right = check_vertical((next.0, next.1 + 1), dir, grid);
                if coords_right.is_none() {
                    return None;
                }
                checked_coords.extend(coords_right.unwrap());
            }

            checked_coords.insert(pos);
        }
        Some(']') => {
            let coords_right = check_vertical(next, dir, grid);
            if coords_right.is_none() {
                return None;
            }
            checked_coords.extend(coords_right.unwrap());

            if grid[&pos] != ']' {
                let coords_left = check_vertical((next.0, next.1 - 1), dir, grid);
                if coords_left.is_none() {
                    return None;
                }
                checked_coords.extend(coords_left.unwrap());
            }

            checked_coords.insert(pos);
        }
        _ => {}
    }

    Some(checked_coords)
}

fn push_horizontal(pos: Coord, dir: Coord, grid: &mut HashMap<Coord, char>) -> Coord {
    let next = (pos.0 + dir.0, pos.1 + dir.1);
    match grid.get(&next) {
        Some('#') => pos, // Wall, no movement
        Some('.') => {
            // Move to empty space
            grid.insert(next, grid[&pos]);
            grid.insert(pos, '.');
            next
        }
        Some(']') | Some('[') => {
            // Attempt to push the object forward
            let pushed_to = push_horizontal(next, dir, grid);
            if pushed_to != next {
                // Successful push of 'O', current `pos` moves to `next`
                grid.insert(next, grid[&pos]);
                grid.insert(pos, '.');
                next
            } else {
                // Push failed, stay in place
                pos
            }
        }
        _ => panic!("Unexpected grid state at {:?}", next),
    }
}

pub fn part_one(input: &str) -> Option<u32> {
    let input_split: Vec<&str> = input.split("\n\n").collect();
    let mut grid = create_grid(input_split[0]);
    let grid_size = input_split[0].lines().count();
    let moves = parse_moves(input_split[1]);
    let pos = grid
        .iter()
        .find_map(|(coord, c)| {
            if *c == '@' {
                return Some(coord);
            }
            None
        })
        .unwrap();

    let mut current_pos = *pos;
    for m in moves {
        current_pos = push(current_pos, m, &mut grid);
    }

    print_map(&grid, grid_size, grid_size);

    let result = grid
        .iter()
        .filter(|&(_, c)| *c == 'O')
        .map(|(coord, _)| (100 * coord.0 + coord.1) as u32)
        .sum();

    Some(result)
}

pub fn part_two(input: &str) -> Option<u32> {
    let input_split: Vec<&str> = input.split("\n\n").collect();
    let mut grid = create_grid_2(input_split[0]);
    let grid_size = input_split[0].lines().count();
    let moves = parse_moves(input_split[1]);
    let pos = grid
        .iter()
        .find_map(|(coord, c)| {
            if *c == '@' {
                return Some(coord);
            }
            None
        })
        .unwrap();

    let mut current_pos = *pos;
    for m in moves {
        if m == DIRS[0] || m == DIRS[2] {
            let coords = check_vertical(current_pos, m, &grid);
            if coords.is_some() {
                let mut coords_vec: Vec<Coord> = coords.unwrap().into_iter().collect();
                coords_vec.sort_by_key(|(y, _)| *y);

                if m == DIRS[2] {
                    coords_vec.reverse()
                }

                for coord in coords_vec {
                    let next = (coord.0 + m.0, coord.1 + m.1);
                    if grid[&coord] == '@' {
                        current_pos = next;
                    }
                    grid.insert(next, grid[&coord]);
                    grid.insert(coord, '.');
                }
            }
        } else {
            current_pos = push_horizontal(current_pos, m, &mut grid);
        }
    }

    print_map(&grid, grid_size, grid_size * 2);
    let result = grid
        .iter()
        .filter(|&(_, c)| *c == '[')
        .map(|(coord, _)| (100 * coord.0 + coord.1) as u32)
        .sum();

    // too low 1419357
    Some(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(10092));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(9021));
    }
}
