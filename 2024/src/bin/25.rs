use regex::Regex;

advent_of_code::solution!(25);

fn is_overlap(a: &Vec<usize>, b: &Vec<usize>, max_height: usize) -> bool {
    for i in 0..a.len() {
        if a[i] + b[i] > max_height {
            return false
        }
    }
    true
}

pub fn part_one(input: &str) -> Option<u32> {
    let keys_and_locks=input.split("\n\n");
    let regex = Regex::new(r"^#+$").unwrap();

    let mut locks: Vec<Vec<usize>> = Vec::new();
    let mut keys: Vec<Vec<usize>> = Vec::new();

    let mut max_height = 0;
    for x in keys_and_locks {
        let lines: Vec<&str> = x.lines().collect();
        let first_line = lines.first().unwrap();

        let mut heights: Vec<usize> = vec![0; first_line.len()];
        if regex.is_match(first_line) {
            for (y, line) in lines.iter().enumerate() {
                for (x, char) in line.chars().enumerate() {
                    if char == '#' && heights[x] < y {
                        heights[x] = y;
                        if y > max_height {
                            max_height = y;
                        }
                    }
                }
            }
            locks.push(heights);
        } else if regex.is_match(lines.last().unwrap()) {
            for (y, line) in lines.iter().rev().enumerate() {
                for (x, char) in line.chars().enumerate() {
                    if char == '#' && heights[x] < y {
                        heights[x] = y;
                        if y > max_height {
                            max_height = y;
                        }
                    }
                }
            }
            keys.push(heights);
        } else {
            panic!("Invalid.")
        }
    }

    // println!("{}", max_height);
    let mut sum = 0;
    for key in keys.iter() {
        for lock in locks.iter() {
            let is_overlap = is_overlap(key, lock, max_height);
            if is_overlap {
                sum += 1;
            }
            // println!("{:?} {:?} {}", lock, key, is_overlap);
        }
    }


    Some(sum)

}

pub fn part_two(input: &str) -> Option<u32> {
    None
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
        assert_eq!(result, None);
    }
}
