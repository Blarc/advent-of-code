use regex::Regex;

advent_of_code::solution!(3);

pub fn part_one(input: &str) -> Option<u32> {
    let mut sum: u32 = 0;
    for line in input.lines() {
        let re = Regex::new(r"mul\(([0-9]+),([0-9]+)\)").unwrap();
        for cap in re.captures_iter(line) {
            let first: u32 = cap[1].parse().expect("Not an integer");
            let second: u32 = cap[2].parse().expect("Not an integer");
            sum += first * second;
        }
    }
    Some(sum)
}

pub fn part_two(input: &str) -> Option<u32> {
    let mut sum: u32 = 0;
    let mut do_use = true;

    for line in input.lines() {
        let re = Regex::new(r"(mul)\(([0-9]+),([0-9]+)\)|(don't\(\))|(do\(\))").unwrap();

        for cap in re.captures_iter(line) {
            if cap.get(1).is_some() && do_use {
                let first: u32 = cap[2].parse().expect("Not an integer");
                let second: u32 = cap[3].parse().expect("Not an integer");
                sum += first * second;
            }
            else if cap.get(4).is_some() {
                do_use = false
            }
            else if cap.get(5).is_some() {
                do_use = true
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
        assert_eq!(result, Some(161));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(49));
    }
}
