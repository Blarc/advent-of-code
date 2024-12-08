use std::collections::{HashMap, HashSet};

advent_of_code::solution!(8);

fn create_grid(input: &str) -> HashMap<char, Vec<(usize, usize)>> {

    let mut m: HashMap<char, Vec<(usize, usize)>> = HashMap::new();
    for (y, l) in input.lines().enumerate() {
        for (x, c) in l.chars().enumerate() {
            if c != '.' {
                m.entry(c)
                    .or_insert_with(Vec::new)
                    .push((y, x));
            }
        }
    }
    m
}

fn anti1(a: (usize, usize), b: (usize, usize), size: usize) -> Option<(usize, usize)> {
    let x= a.0 as i32 + (a.0 as i32 - b.0 as i32);
    let y = a.1 as i32 + (a.1 as i32 - b.1 as i32);

    if 0 <= x && x < size as i32 && 0 <= y && y < size as i32 {
        return Some((x as usize, y as usize))
    }
    None
}

fn anti2(a: (usize, usize), b: (usize, usize), size: usize) -> HashSet<(usize, usize)> {
    let mut y = a.0 as i32;
    let mut x = a.1 as i32;

    let mut nodes = HashSet::new();
    while 0 <= x && x < size as i32 && 0 <= y && y < size as i32 {
        nodes.insert((y as usize, x as usize));
        y += a.0 as i32 - b.0 as i32;
        x += a.1 as i32 - b.1 as i32;
    }
    nodes
}


pub fn part_one(input: &str) -> Option<u32> {
    let size = input.lines().count();
    let m = create_grid(input);

    let mut result = HashSet::new();
    for (_, antennas) in m {
        for i in 0..antennas.len() {
            for j in (i+1)..antennas.len() {
                let a = antennas[i];
                let b = antennas[j];

                let op1 = anti1(a, b, size);
                let op2 = anti1(b, a, size);
                if op1.is_some() {
                    result.insert(op1.unwrap());
                }
                if op2.is_some() {
                    result.insert(op2.unwrap());
                }
            }
        }
    }

    Some(result.len() as u32)
}

pub fn part_two(input: &str) -> Option<u32> {
    let size = input.lines().count();
    let m = create_grid(input);

    let mut result: HashSet<(usize, usize)> = HashSet::new();
    for (_, antennas) in m {
        for i in 0..antennas.len() {
            for j in (i+1)..antennas.len() {
                let a = antennas[i];
                let b = antennas[j];

                let op1 = anti2(a, b, size);
                let op2 = anti2(b, a, size);

                result.extend(&op1);
                result.extend(&op2);
            }
        }
    }


    Some(result.len() as u32)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(14));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(34));
    }
}
