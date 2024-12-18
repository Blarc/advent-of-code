use std::collections::{HashMap, HashSet, VecDeque};
use regex::Regex;

advent_of_code::solution!(18);

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
            print!("{}", grid.get(&(y, x)).unwrap_or(&'?'));
        }
        println!()
    }
}

pub fn part_one(input: &str) -> Option<u32> {
    let program_regex = Regex::new(r"-?[0-9]+").unwrap();

    let mut grid: HashMap<Coord, char> = HashMap::new();

    let bytes = 1024;
    for line in input.lines().take(bytes) {
        let program: Vec<i32> = program_regex
            .find_iter(line)
            .filter_map(|mat| mat.as_str().parse::<i32>().ok())
            .collect();

        grid.insert((program[1], program[0]), '#');
    }

    let size = 70;
    for y in 0..size+1{
        for x in 0..size+1 {
            if !grid.contains_key(&(y, x)) {
                grid.insert((y, x), '.');
            }
        }
    }

    // print_grid(&grid);
    let min = bfs(&mut grid, size);
    Some(min)
}

fn bfs(grid: &mut HashMap<Coord, char>, size: i32) -> u32 {
    let pos = (0, 0);
    let mut min = u32::MAX;
    let mut queue = VecDeque::new();
    queue.push_back((pos, 0));
    let mut visited = HashSet::new();
    visited.insert(pos);

    while let Some((pos, current_cost)) = queue.pop_front() {
        if pos == (size, size) {
            min = min.min(current_cost);
            continue;
        }

        for new_dir in DIRS {
            let mut new_cost = current_cost + 1;

            let next_pos = (pos.0 + new_dir.0, pos.1 + new_dir.1);
            if grid.contains_key(&next_pos) && grid[&next_pos] != '#' && !visited.contains(&(next_pos)) {
                visited.insert(next_pos);
                queue.push_back((next_pos, new_cost));
            }
        }
    }
    min
}

pub fn part_two(input: &str) -> Option<u32> {
    let program_regex = Regex::new(r"-?[0-9]+").unwrap();

    let mut grid: HashMap<Coord, char> = HashMap::new();


    let mut bytes = Vec::new();
    for line in input.lines() {
        let program: Vec<i32> = program_regex
            .find_iter(line)
            .filter_map(|mat| mat.as_str().parse::<i32>().ok())
            .collect();

        bytes.push((program[1], program[0]));

    }

    let num_of_bytes = 1024;
    for byte in bytes.iter().take(num_of_bytes) {
        grid.insert(*byte, '#');
    }

    let size = 70;
    for y in 0..size+1{
        for x in 0..size+1 {
            if !grid.contains_key(&(y, x)) {
                grid.insert((y, x), '.');
            }
        }
    }

    print_grid(&grid);

    let mut index = num_of_bytes + 1;
    loop {
        grid.insert(bytes[index], '#');
        let min = bfs(&mut grid, size);
        if min == u32::MAX {
            println!("{:?}", bytes[index]);
            break
        }
        index += 1;
    }

    Some(1)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(22));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
