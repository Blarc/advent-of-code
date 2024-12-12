use std::collections::{HashMap, HashSet};

advent_of_code::solution!(12);

const DIRS: [(i32, i32); 4] = [(-1, 0), (0, 1), (1, 0), (0, -1)];

type Coord = (i32, i32);
fn create_grid(input: &str) -> HashMap<Coord, char> {
    input
        .lines()
        .enumerate()
        .flat_map(|(y, line)| {
            line.chars().enumerate().map(move |(x, c)| {
                (
                    (y as i32, x as i32),
                    c
                )
            })
        })
        .collect()
}

pub fn dfs(pos: Coord, grid: &HashMap<Coord, char>, visited: &mut HashSet<Coord>) -> (u32, u32) {
    visited.insert(pos);
    let mut area = 1;
    let mut perimeter = 4;

    for dir in DIRS {
        let neighbour = (pos.0 + dir.0, pos.1 + dir.1);
        if grid.get(&pos) == grid.get(&neighbour) {
            perimeter -= 1;
            if !visited.contains(&neighbour) {
                let (neighbour_area, neighbour_perimeter) = dfs(neighbour, grid, visited);
                area += neighbour_area;
                perimeter += neighbour_perimeter;
            }
        }
    }
    (area, perimeter)
}

fn get_sides(pos: Coord, grid: &HashMap<Coord, char>) -> [bool; 4] {
    let mut sides = [false, false, false, false];
    for (i, dir) in DIRS.iter().enumerate() {
        let neighbour = (pos.0 + dir.0, pos.1 + dir.1);
        if grid.get(&pos) != grid.get(&neighbour) {
            sides[i] = true;
        }
    }
    sides
}

fn check(pos: Coord, grid: &HashMap<Coord, char>, visited: &mut HashSet<Coord>, dir: usize, side: usize) -> bool {
    let dy = DIRS[dir % 4].0;
    let dx = DIRS[dir % 4].1;

    let mut neighbour = (pos.0 + dy, pos.1 + dx);
    while grid.get(&pos) == grid.get(&neighbour) {
        if get_sides(neighbour, grid)[side] {
            if visited.contains(&neighbour) {
                return false
            }
            neighbour = (neighbour.0 + dy, neighbour.1 + dx);
        } else {
            return true
        }
    }

    true
}

pub fn dfs2(pos: Coord, grid: &HashMap<Coord, char>, visited: &mut HashSet<Coord>) -> (u64, u64) {
    visited.insert(pos);

    let mut n_sides = 0;
    let sides = get_sides(pos, grid);
    for (i, side) in sides.iter().enumerate() {
        if *side {
            if check(pos, grid, visited, i+1, i) && check(pos, grid, visited, i+3, i) {
                n_sides += 1;
            }
        }
    }

    let mut area = 1;
    for dir in DIRS {
        let neighbour = (pos.0 + dir.0, pos.1 + dir.1);
        if grid.get(&pos) == grid.get(&neighbour) {
            if !visited.contains(&neighbour) {
                let (neighbour_area, neighbour_sides) = dfs2(neighbour, grid, visited);
                area += neighbour_area;
                n_sides += neighbour_sides;
            }
        }
    }

    (area, n_sides)
}

pub fn part_one(input: &str) -> Option<u32> {
    let grid = create_grid(input);
    let visited = &mut HashSet::new();

    let mut result = 0;
    for coord in grid.keys().clone() {
        if !visited.contains(coord) {
            let (area, perimeter) = dfs(*coord, &grid, visited);
            result += area * perimeter;
        }
    }
    Some(result)
}

pub fn part_two(input: &str) -> Option<u64> {
    let grid = create_grid(input);
    let visited = &mut HashSet::new();

    let mut result = 0;
    for coord in grid.keys().clone() {
        if !visited.contains(coord) {
            let (area, sides) = dfs2(*coord, &grid, visited);
            result += area * sides;
        }
    }

    // 910474 too low
    Some(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(1930));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(1206));
    }
}
