use regex::Regex;

advent_of_code::solution!(7);

type Operation = fn(&u64, &u64) -> u64;

const OPS1: [fn(&u64, &u64) -> u64; 2] = [|a, b| a + b, |a, b| a * b];

const OPS2: [Operation; 3] = [
    |a, b| a + b,
    |a, b| a * b,
    |a, b| a * 10_u64.pow(power(*b)) + b,
];

fn power(number: u64) -> u32 {
    let mut power = 0;
    let mut result = 1;

    while result * 10 <= number {
        result *= 10;
        power += 1;
    }
    power + 1
}

fn is_solvable(
    target: u64,
    result: u64,
    numbers: &[u64],
    mut step: usize,
    ops: &[Operation],
) -> bool {
    if step.eq(&(numbers.len() - 1)) {
        return target == result;
    }

    if result > target {
        return false;
    }

    for op in ops {
        let r = op(&result, &numbers[step + 1]);
        step += 1;
        if is_solvable(target, r, numbers, step, ops) {
            return true;
        }
        step -= 1;
    }
    false
}

pub fn solve(input: &str, ops: &[Operation]) -> u64 {
    let mut sum = 0;
    for line in input.lines() {
        let re = Regex::new(r"[0-9]+").unwrap();
        let num: Vec<u64> = re
            .find_iter(line)
            .filter_map(|mat| mat.as_str().parse::<u64>().ok())
            .collect();

        let solvable = is_solvable(num[0], num[1], &num[1..num.len()], 0, ops);
        if solvable {
            sum += num[0];
        }
    }
    sum
}

pub fn part_one(input: &str) -> Option<u64> {
    Some(solve(input, &OPS1))
}

pub fn part_two(input: &str) -> Option<u64> {
    Some(solve(input, &OPS2))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(3749));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(11387));
    }
}
