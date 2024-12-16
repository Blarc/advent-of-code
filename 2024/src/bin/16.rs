use std::collections::{HashMap, HashSet, VecDeque};

advent_of_code::solution!(16);

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

fn print_grid(grid: &HashMap<Coord, char>) {
    let height = grid
        .iter()
        .max_by_key(|&(c, _)| c.0)
        .map(|(c, _)| c.0)
        .unwrap();
    let width = grid
        .iter()
        .max_by_key(|&(c, _)| c.1)
        .map(|(c, _)| c.1)
        .unwrap();

    for y in 0..height+1 {
        for x in 0..width+1 {
            // print!("{:?}", (y as i32, x as i32));
            print!("{}", grid.get(&(y, x)).unwrap());
        }
        println!()
    }
}

pub fn part_one(input: &str) -> Option<u32> {
    let grid = create_grid(input);

    let pos = grid
        .iter()
        .find_map(|(coord, c)| {
            if *c == 'S' {
                return Some(coord);
            }
            None
        })
        .unwrap();

    // print_grid(&grid);

    let mut min = u32::MAX;
    let mut queue = VecDeque::new();
    queue.push_back((*pos, DIRS[1], 0));
    let mut visited = HashMap::new();
    visited.insert((*pos, DIRS[2]), 0);

    while let Some((pos, dir, current_cost)) = queue.pop_front() {
        if grid[&pos] == 'E' {
            min = min.min(current_cost);
            continue;
        }

        for new_dir in DIRS {
            let mut new_cost = current_cost + 1;
            if new_dir != dir {
                new_cost += 1000;
            }

            let next_pos = (pos.0 + new_dir.0, pos.1 + new_dir.1);
            if grid.contains_key(&next_pos) && grid[&next_pos] != '#' && (!visited.contains_key(&(next_pos, new_dir)) || new_cost < visited[&(next_pos, new_dir)]) {
                visited.insert((next_pos, new_dir), new_cost);
                queue.push_back((next_pos, new_dir, new_cost));
            }
        }
    }

    // 156504 not correct
    Some(min)
}

pub fn part_two(input: &str) -> Option<u32> {
    let grid = create_grid(input);

    let pos = grid
        .iter()
        .find_map(|(coord, c)| {
            if *c == 'S' {
                return Some(coord);
            }
            None
        })
        .unwrap();

    // print_grid(&grid);

    let mut min = u32::MAX;
    let mut best_tiles = HashSet::new();
    let mut queue = VecDeque::new();

    let mut path = Vec::new();
    path.push(*pos);

    queue.push_back((*pos, DIRS[1], 0, path));
    let mut visited = HashMap::new();
    visited.insert((*pos, DIRS[2]), 0);

    while let Some((pos, dir, current_cost, path)) = queue.pop_front() {
        if grid[&pos] == 'E' {
            if current_cost < min {
                min = current_cost;
                best_tiles = HashSet::from_iter(path);
            } else if current_cost == min {
                best_tiles.extend(path);
            }
            continue;
        }

        for new_dir in DIRS {
            let mut new_cost = current_cost + 1;
            if new_dir != dir {
                new_cost += 1000;
            }

            let next_pos = (pos.0 + new_dir.0, pos.1 + new_dir.1);
            if grid.contains_key(&next_pos) && grid[&next_pos] != '#' && (!visited.contains_key(&(next_pos, new_dir)) || new_cost <= visited[&(next_pos, new_dir)]) {
                visited.insert((next_pos, new_dir), new_cost);

                let mut new_path = path.clone();
                new_path.push(next_pos);
                queue.push_back((next_pos, new_dir, new_cost, new_path));
            }
        }
    }

    Some(best_tiles.len() as u32)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(11048));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(64));
    }
}
