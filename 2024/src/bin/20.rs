use std::collections::{HashMap, HashSet, VecDeque};

advent_of_code::solution!(20);

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

    for y in 0..height + 1 {
        for x in 0..width + 1 {
            // print!("{:?}", (y as i32, x as i32));
            print!("{}", grid.get(&(y, x)).unwrap_or(&'?'));
        }
        println!()
    }
}

fn grid_find(target: char, grid: &HashMap<Coord, char>) -> &Coord {
    grid.iter()
        .find_map(|(coord, c)| {
            if *c == target {
                return Some(coord);
            }
            None
        })
        .unwrap()
}

fn bfs(
    grid: &HashMap<Coord, char>,
    pos: &Coord,
    target: &Coord,
    starting_cheat: Option<(Coord, Coord)>,
    base_cost: Option<i32>,
) -> HashMap<Option<((i32, i32), (i32, i32))>, i32> {
    let mut queue = VecDeque::new();
    queue.push_back((*pos, 0, starting_cheat));

    let mut visited = HashSet::new();
    visited.insert((*pos, starting_cheat));

    let mut cheats = HashMap::new();

    while let Some((pos, current_cost, cheat)) = queue.pop_front() {
        if grid[&pos] == 'E' {
            cheats.insert(cheat, current_cost);
            continue;
        }

        if visited.contains(&(*target, cheat)) {
            continue;
        }

        if base_cost.is_some() && current_cost > base_cost.unwrap() {
            continue;
        }

        // Check all possible directions
        for new_dir in DIRS {
            let next_pos = (pos.0 + new_dir.0, pos.1 + new_dir.1);

            if let Some(&next_char) = grid.get(&next_pos) {
                // If we encounter a '#', attempt to use the cheat if available
                if next_char == '#' && cheat.is_none() {
                    let next_pos_2 = (next_pos.0 + new_dir.0, next_pos.1 + new_dir.1);

                    if let Some(&next_char_2) = grid.get(&next_pos_2) {
                        if next_char_2 != '#'
                            && !visited.contains(&(next_pos_2, Some((next_pos, next_pos_2))))
                        {
                            visited.insert((next_pos_2, Some((next_pos, next_pos_2))));
                            queue.push_back((
                                next_pos_2,
                                current_cost + 2,
                                Some((next_pos, next_pos_2)),
                            ));
                        }
                    }
                }
                // If it's not a wall ('#'), proceed as usual
                else if next_char != '#' {
                    if !visited.contains(&(next_pos, cheat)) {
                        visited.insert((next_pos, cheat));
                        queue.push_back((next_pos, current_cost + 1, cheat));
                    }
                }
            }
        }
    }
    cheats
}

fn find_path(
    start: &Coord,
    end: &Coord,
    grid: &HashMap<Coord, char>,
) -> (Vec<Coord>, HashMap<Coord, usize>) {
    let mut visited = HashMap::new();
    let mut path = Vec::new();

    let mut pos = start.clone();
    visited.insert(pos, 0);
    path.push(pos);

    let mut index = 1;
    while pos != *end {
        for new_dir in DIRS {
            let next_pos = (pos.0 + new_dir.0, pos.1 + new_dir.1);
            if let Some(&next_char) = grid.get(&next_pos) {
                if !visited.contains_key(&next_pos) && next_char != '#' {
                    pos = next_pos;
                    visited.insert(pos, index);
                    path.push(pos);
                    index += 1;
                }
            }
        }
    }

    (path, visited)
}

pub fn part_one(input: &str) -> Option<u32> {
    let grid = create_grid(input);
    // print_grid(&grid);

    let pos = grid_find('S', &grid);
    let end = grid_find('E', &grid);
    println!("start {:?}\nend {:?}", pos, end);

    let mut result = 0;
    let mut result_m = HashMap::new();
    let (path, cost) = find_path(pos, end, &grid);
    // println!("{:?}", path);
    // println!("{:?}", path.len());
    // println!("{:?}", cost);
    // println!("{:?}", cost.len());
    println!("Base path cost: {:?}", cost[&end]);

    for pos in path.iter() {
        for dir in DIRS {
            let next_pos = (pos.0 + dir.0, pos.1 + dir.1);
            if let Some(&next_char) = grid.get(&next_pos) {
                if next_char == '#' {
                    let next_pos_2 = (next_pos.0 + dir.0, next_pos.1 + dir.1);
                    if let Some(&next_char_2) = grid.get(&next_pos_2) {
                        if next_char_2 != '#' {
                            let pos_cost = cost[&pos];
                            let cheat_to_end = cost[&end] - cost[&next_pos_2];

                            let full_cost = pos_cost + cheat_to_end + 2;
                            if full_cost < cost[&end] {
                                let saved = cost[&end] - full_cost;
                                if saved >= 100 {
                                    *result_m.entry(saved).or_insert(0) += 1;
                                    result += 1;
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    // Collect into a Vec of (key, value) pairs
    let mut entries: Vec<(usize, i32)> = result_m.iter()
        .map(|(&key, &value)| (key, value))
        .collect();

    // Sort by keys
    entries.sort_by(|a, b| a.0.cmp(&b.0));

    // for x in entries.iter() {
    //     println!("{:?}", x);
    // }

    // Too high 7068
    Some(result)
}


fn manhattan_distance(a: (i32, i32), b: (i32, i32)) -> usize {
    ((a.0 - b.0).abs() + (a.1 - b.1).abs()) as usize
}

pub fn part_two(input: &str) -> Option<u32> {
    let grid = create_grid(input);
    // print_grid(&grid);
    let size = 20;

    let pos = grid_find('S', &grid);
    let end = grid_find('E', &grid);
    println!("start {:?}\nend {:?}", pos, end);

    let mut result = 0;
    let mut result_m = HashMap::new();
    let mut result_ma = HashSet::new();
    let (path, cost) = find_path(pos, end, &grid);
    println!("Base path cost: {:?}", cost[&end]);

    for pos in path.iter() {
        for y in -size-1..size+1 {
            for x in -size-1..size+1 {
                let next_pos = (pos.0 + y, pos.1 + x);
                if manhattan_distance(*pos, next_pos) <= size as usize {
                    if let Some(&next_char) = grid.get(&next_pos) {
                        if next_char != '#' {
                            let pos_cost = cost[&pos];
                            let cheat_to_end = cost[&end] - cost[&next_pos];

                            let full_cost = pos_cost + cheat_to_end + manhattan_distance(*pos, next_pos);
                            if full_cost < cost[&end] {
                                let saved = cost[&end] - full_cost;
                                if saved >= 100 {
                                    *result_m.entry(saved).or_insert(0) += 1;
                                    result_ma.insert((pos, next_pos));
                                    result += 1;
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    // Collect into a Vec of (key, value) pairs
    let mut entries: Vec<(usize, i32)> = result_m.iter()
        .map(|(&key, &value)| (key, value))
        .collect();

    // Sort by keys
    entries.sort_by(|a, b| a.0.cmp(&b.0));

    // for x in entries.iter() {
    //     println!("{:?}", x);
    // }

    // Too high 1066443
    Some(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(44));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(43));
    }
}
