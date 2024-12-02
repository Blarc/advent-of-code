use std::str::FromStr;

advent_of_code::solution!(2);

pub fn is_safe(n: &Vec<u32>) -> bool {
    let mut increasing = None;
    for i in 0..n.len()-1 {
        let first = n[i];
        let second = n[i+1];

        if first < second && first.abs_diff(second) <= 3 {
            if let Some(false) = increasing {
                return false;
            }
            increasing = Some(true);
        } else if first > second && first.abs_diff(second) <= 3 {
            if let Some(true) = increasing {
                return false;
            }
            increasing = Some(false);
        } else {
            return false;
        }
    }
    true
}

pub fn copy_and_remove_index(n: &Vec<u32>, index: usize) -> Vec<u32> {
    let mut new_vec = n.to_vec();
    new_vec.remove(index);
    new_vec
}

pub fn part_one(input: &str) -> Option<u32> {

    let mut sum: u32 = 0;
    for line in input.lines() {
        let numbers: Vec<u32> = line
            .split(" ")
            .filter_map(|s| u32::from_str(s).ok())
            .collect();

        let safe = is_safe(&numbers);
        if safe {
            sum += 1
        }
    }

    Some(sum)
}

pub fn part_two(input: &str) -> Option<u32> {
    let mut sum: u32 = 0;
    for line in input.lines() {
        let numbers: Vec<u32> = line
            .split(" ")
            .filter_map(|s| u32::from_str(s).ok())
            .collect();

        let mut safe = is_safe(&numbers);
        if !safe {
            for i in 0..numbers.len() {
                safe = safe || is_safe(&copy_and_remove_index(&numbers, i))
            }
        }
        if safe {
            sum += 1
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
        assert_eq!(result, Some(3));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(11));
    }
}
