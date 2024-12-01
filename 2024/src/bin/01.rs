use std::collections::HashMap;

advent_of_code::solution!(1);

pub fn part_one(input: &str) -> Option<u32> {

    let mut firsts = vec![0; input.lines().count()];
    let mut seconds = vec![0; input.lines().count()];

    let mut index = 0;
    for line in input.lines() {
        let split_line: Vec<&str> = line.split(" ").collect();
        let first: u32 = split_line.first()?.parse().expect("Not an integer");
        let second: u32 = split_line.last()?.parse().expect("Not an integer");
        firsts[index] = first;
        seconds[index] = second;
        index += 1;
    }

    firsts.sort();
    seconds.sort();

    let mut sum: u32 = 0;
    for i in 0..firsts.len() {
        sum += firsts[i].abs_diff(seconds[i]);
    }

    Some(sum)
}

pub fn part_two(input: &str) -> Option<u32> {

    let mut firsts = vec![0; input.lines().count()];
    let mut count: HashMap<u32, u32> = HashMap::new();

    let mut index = 0;
    for line in input.lines() {
        let split_line: Vec<&str> = line.split(" ").collect();
        let first: u32 = split_line.first()?.parse().expect("Not an integer");
        firsts[index] = first;
        count.insert(first, 0);
        index += 1;
    }

    for line in input.lines() {
        let split_line: Vec<&str> = line.split(" ").collect();
        let second: u32 = split_line.last()?.parse().expect("Not an integer");

        if count.contains_key(&second) {
            let c = count.entry(second).or_default();
            *c += 1;
        }
    }

    let mut sum: u32 = 0;
    for first in firsts {

        if count.contains_key(&first) {
            let tmp = first * count.get(&first).unwrap();
            sum += tmp
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
        assert_eq!(result, Some(11));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(31));
    }
}
