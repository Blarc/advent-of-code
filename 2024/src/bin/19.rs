use std::collections::{HashMap, HashSet};
use std::fs::read_to_string;

advent_of_code::solution!(19);

fn is_possible(c: &str, goal: &str, patterns: &Vec<&str>, visited: &mut HashSet<String>) -> bool {
    // println!("{}", c);
    if c == goal {
        return true;
    }

    if c.len() >= goal.len() {
        return false
    }

    if visited.contains(c) {
        return false;
    }

    visited.insert(c.to_string());

    for pattern in patterns {
        let new = format!("{}{}", c, pattern);
        if goal.starts_with(&new) && is_possible(&new, goal, patterns, visited) {
            return true;
        }
    }

    false
}

fn num_of_options(c: &str, goal: &str, patterns: &Vec<&str>, visited: &mut HashMap<String, u64>) -> u64 {
    if c == goal {
        return 1;
    }

    if c.len() >= goal.len() {
        return 0
    }

    if visited.contains_key(c) {
        return visited[c];
    }

    let mut result = 0;
    for pattern in patterns {
        let new = format!("{}{}", c, pattern);
        if goal.starts_with(&new) {
            result += num_of_options(&new, goal, patterns, visited);
        }
    }

    visited.insert(c.to_string(), result);
    result
}

pub fn part_one(input: &str) -> Option<u32> {
    let input_split: Vec<&str> = input.split("\n\n").collect();
    let patterns: Vec<&str> = input_split[0].split(", ").collect();

    let mut sum = 0;
    for line in input_split[1].trim().lines() {
        if is_possible("", line.trim(), &patterns, &mut HashSet::new()) {
            sum += 1;
        }
    }
    Some(sum)
}

pub fn part_two(input: &str) -> Option<u64> {
    let input_split: Vec<&str> = input.split("\n\n").collect();
    let patterns: Vec<&str> = input_split[0].split(", ").collect();

    let mut sum = 0;
    for line in input_split[1].trim().lines() {
        sum += num_of_options("", line.trim(), &patterns, &mut HashMap::new())
    }
    Some(sum)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(6));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(16));
    }
}
