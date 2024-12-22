use std::collections::HashMap;

advent_of_code::solution!(22);

const PRUNE: u128 = 16777216;

fn next_secret(mut x: u128) -> u128 {
    let mut new_x;

    new_x = x * 64;
    x = new_x ^ x;
    x = x % PRUNE;

    new_x = x / 32;
    x = new_x ^ x;
    x = x % PRUNE;

    new_x = x * 2048;
    x = new_x ^ x;
    x = x % PRUNE;

    x
}

pub fn part_one(input: &str) -> Option<u128> {

    let mut sum = 0;
    for line in input.lines() {
        let mut x: u128 = line.parse::<u128>().unwrap();
        for _ in 0..2000 {
            x = next_secret(x);
        }
        // println!("{}: {}", line, x);
        sum += x;
    }
    Some(sum)
}

pub fn part_two(input: &str) -> Option<u128> {

    let mut best_sequences: HashMap<Vec<i32>, Vec<i32>> = HashMap::new();
    for (line_index, line) in input.lines().enumerate() {
        let mut x: u128 = line.parse::<u128>().unwrap();
        let mut previous_digit = 0;
        let mut last_four = vec![];
        for i in 0..1999 {
            let current_digit: i32 = (x % 10) as i32;

            if i > 0 {
                let diff = current_digit - previous_digit;
                last_four.push(diff);

                if i > 3  {
                    if !best_sequences.contains_key(&last_four) || best_sequences[&last_four][line_index] == 0 {
                        if let Some(existing_vec) = best_sequences.get(&last_four) {
                            let mut tmp = existing_vec.clone();
                            tmp[line_index] = current_digit;
                            best_sequences.insert(last_four.clone(), tmp);
                        } else {
                            let mut tmp = vec![0; input.lines().count()];
                            tmp[line_index] = current_digit;
                            best_sequences.insert(last_four.clone(), tmp);
                        }
                    }

                    last_four.remove(0);
                }
            }

            previous_digit = current_digit;
            x = next_secret(x);
        }
    }

    let mut max = 0;
    for (_, value) in best_sequences.clone() {
        let tmp = value.iter().sum::<i32>() as u128;
        if tmp > max {
            max = tmp;
        }
    }

    Some(max)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(37327623));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(23));
    }
}
