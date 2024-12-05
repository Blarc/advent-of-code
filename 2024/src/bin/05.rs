use std::cmp::Ordering;
use std::collections::{HashMap, HashSet};

advent_of_code::solution!(5);

pub fn part_one(input: &str) -> Option<u32> {
    let mut rules: HashMap<u32, HashSet<u32>> = HashMap::new();

    let mut read_updates = false;
    let mut sum = 0;
    for line in input.lines() {
        if line.is_empty() {
            read_updates = true;
            continue;
        }

        if !read_updates {
            let n: Vec<u32> = line
                .split("|")
                .map(|x| x.parse().expect("Not an integer"))
                .collect();

            rules
                .entry(n[0])
                .and_modify(|s| {
                    s.insert(n[1]);
                })
                .or_insert_with(|| {
                    let mut set: HashSet<u32> = HashSet::new();
                    set.insert(n[1]);
                    set
                });
        } else {
            let n: Vec<u32> = line
                .split(",")
                .map(|x| x.parse().expect("Not an integer"))
                .collect();

            let correct = n.is_sorted_by(|x1, x2| {
                if let Some(rule1) = rules.get(x1) {
                    if rule1.contains(x2) {
                        return true;
                    }
                }

                if let Some(rule2) = rules.get(x2) {
                    if rule2.contains(x1) {
                        return false;
                    }
                }

                return true;
            });


            if correct {
                let middle_index = n.len() / 2;
                sum += n[middle_index];
                // println!("{:?}: {}", n, n[middle_index])
            }
        }
    }
    Some(sum)
}

pub fn part_two(input: &str) -> Option<u32> {
    let mut rules: HashMap<u32, HashSet<u32>> = HashMap::new();

    let mut read_updates = false;
    let mut sum = 0;
    for line in input.lines() {
        if line.is_empty() {
            read_updates = true;
            continue;
        }

        if !read_updates {
            let n: Vec<u32> = line
                .split("|")
                .map(|x| x.parse().expect("Not an integer"))
                .collect();

            rules
                .entry(n[0])
                .and_modify(|s| {
                    s.insert(n[1]);
                })
                .or_insert_with(|| {
                    let mut set: HashSet<u32> = HashSet::new();
                    set.insert(n[1]);
                    set
                });
        } else {
            let mut n: Vec<u32> = line
                .split(",")
                .map(|x| x.parse().expect("Not an integer"))
                .collect();

            let correct = n.is_sorted_by(|x1, x2| {
                if let Some(rule1) = rules.get(x1) {
                    if rule1.contains(x2) {
                        return true;
                    }
                }

                if let Some(rule2) = rules.get(x2) {
                    if rule2.contains(x1) {
                        return false;
                    }
                }

                return true;
            });


            if !correct {

                n.sort_by(|x3, x4| {
                    if let Some(rule1) = rules.get(x3) {
                        if rule1.contains(x4) {
                            return Ordering::Less;
                        }
                    }

                    if let Some(rule2) = rules.get(x4) {
                        if rule2.contains(x3) {
                            return Ordering::Greater;
                        }
                    }

                    return Ordering::Equal;
                });

                let middle_index = n.len() / 2;
                sum += n[middle_index];
                // println!("{:?}: {}", n, n[middle_index])
            }
        }
    }
    Some(sum)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(143));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(123));
    }
}
