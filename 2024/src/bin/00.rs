advent_of_code::solution!(1);

pub fn part_one(input: &str) -> Option<u32> {
    let mut max: u32 = 0;
    let mut sum: u32 = 0;
    for line in input.lines() {
        if line.is_empty() {
            if sum > max {
                max = sum
            }
            sum = 0;
            continue;
        }
        let current: u32 = line.trim().parse().unwrap();
        sum += current;
    }
    Some(max)
}

pub fn part_two(input: &str) -> Option<u32> {
    let mut sum: u32 = 0;

    let mut maxs: [u32; 3] = [0, 0, 0];
    for line in input.lines() {
        if line.is_empty() {
            for i in 0..maxs.len() {
                if sum > maxs[i] {
                    let mut j = maxs.len() - 1;
                    while j > i {
                        maxs[j] = maxs[j-1];
                        j = j - 1;
                    }

                    maxs[i] = sum;
                    break;
                }
            }
            sum = 0;
            continue;
        }

        let current: u32 = line.trim().parse().unwrap();
        sum += current;
    }

    let result: u32 = maxs.iter().sum();
    Some(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(24000));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(45000));
    }
}
