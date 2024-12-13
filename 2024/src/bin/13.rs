use regex::Regex;

advent_of_code::solution!(13);

// 8400 = x * 94 + y * 22
// 5400 = x * 34 + y * 67


// 8400 / 22 = x * 94 / 22 + y
// 8400 / 22 - x * 94 / 22 = y
// (8400 - x * 94) / 22 = y
// y = (8400 - x * 94) / 22

// 5400 = x * 34 + y * 67
// 5400 = x * 34 + ((8400 - x * 94) / 22) * 67
// 5400 / 67 = (x * 34) / 67 + (8400 - x * 94) / 22
// 5400 * 22 / 67 = ((x * 34) * 22) / 67 + 8400 - x * 94
// 5400 * 22 / 67 * 94 = ((x * 34) * 22) / 67 * 94 + 8400 / 94 - x
// 5400 * 22 / 67 * 94 - 8400 / 94 = ((x * 34) * 22) / 67 * 94 - x

pub fn part_one(input: &str) -> Option<i128> {
    let re = Regex::new(r"[0-9]+").unwrap();

    let mut sum: i128 = 0;
    let mut nums = vec![];
    for line in input.lines() {
        if !line.is_empty() {
            let num: Vec<i128> = re
                .find_iter(line)
                .filter_map(|mat| mat.as_str().parse::<i128>().ok())
                .collect();

            nums.extend(num)
        } else {
            let ax = nums[0];
            let ay = nums[1];
            let bx = nums[2];
            let by = nums[3];
            let px = nums[4];
            let py = nums[5];

            if ((bx * py - px * by) % (bx * ay - ax * by)) == 0 {
                let x = (bx * py - px * by) / (bx * ay - ax * by);

                if ((px - x * ax) % bx) == 0 {
                    let y = (px - x * ax) / bx;
                    sum += x * 3 + y;
                }
            }

            nums = vec![];
        }
    }

    // 31612 too high
    Some(sum)
}

pub fn part_two(input: &str) -> Option<i128> {
    let re = Regex::new(r"[0-9]+").unwrap();

    let mut sum: i128 = 0;
    let mut nums = vec![];
    for line in input.lines() {
        if !line.is_empty() {
            let num: Vec<i128> = re
                .find_iter(line)
                .filter_map(|mat| mat.as_str().parse::<i128>().ok())
                .collect();

            nums.extend(num)
        } else {
            let ax = nums[0];
            let ay = nums[1];
            let bx = nums[2];
            let by = nums[3];
            let px = nums[4] + 10000000000000;
            let py = nums[5] + 10000000000000;

            if ((bx * py - px * by) % (bx * ay - ax * by)) == 0 {
                let x = (bx * py - px * by) / (bx * ay - ax * by);

                if ((px - x * ax) % bx) == 0 {
                    let y = (px - x * ax) / bx;
                    sum += x * 3 + y;
                }
            }

            nums = vec![];
        }
    }

    // 102718967795500
    // 102302974257181 too low
    Some(sum)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(480));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(875318608908));
    }
}
