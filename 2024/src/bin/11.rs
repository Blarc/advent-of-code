use std::collections::HashMap;

advent_of_code::solution!(11);

fn parse(input: &str) -> Vec<u64> {
    input
        .trim()
        .split(" ")
        .map(|x| x.parse().expect("Not an integer"))
        .collect()
}

fn blink(stone: u64, i: u64, t: u64, memo: &mut HashMap<(u64, u64), u64>) -> u64 {
    if let Some(value) = memo.get(&(stone, i)) {
        return *value;
    }

    if i == t {
        return 1;
    }

    let result;
    if stone == 0 {
        result = blink(1, i + 1, t, memo);
    } else if stone.ilog10() % 2 == 1 {
        let log = (stone.ilog10() + 1) / 2;
        let first = stone / 10_u64.pow(log);
        let second = stone % 10_u64.pow(log);
        result = blink(first, i + 1, t, memo) + blink(second, i + 1, t, memo);
    } else {
        result = blink(stone * 2024, i + 1, t, memo);
    }

    memo.insert((stone, i), result);
    result
}

pub fn part_one(input: &str) -> Option<u64> {
    let stones = parse(input);
    let memo = &mut HashMap::new();
    let mut sum = 0;
    for stone in stones {
        sum += blink(stone, 0, 25, memo);
    }
    Some(sum)
}

pub fn part_two(input: &str) -> Option<u64> {
    let stones = parse(input);
    let memo = &mut HashMap::new();
    let mut sum = 0;
    for stone in stones {
        sum += blink(stone, 0, 75, memo);
    }
    Some(sum)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(55312));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(65601038650482));
    }
}
