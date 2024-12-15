use regex::Regex;
use std::collections::HashSet;

advent_of_code::solution!(14);

const WIDTH: u32 = 101; // 11 / 101
const HALF_WIDTH: u32 = WIDTH / 2;
const HEIGHT: u32 = 103; // 7 / 103
const HALF_HEIGHT: u32 = HEIGHT / 2;
const TIME: usize = 100;

fn pos_at_time(robot: Vec<i32>, time: usize) -> (u32, u32) {
    let x = ((robot[0] + robot[2] * time as i32) % WIDTH as i32 + WIDTH as i32) as u32 % WIDTH;
    let y = ((robot[1] + robot[3] * time as i32) % HEIGHT as i32 + HEIGHT as i32) as u32 % HEIGHT;
    (x, y)
}

pub fn part_one(input: &str) -> Option<u32> {
    let re = Regex::new(r"-?[0-9]+").unwrap();

    let mut quadrants = [0, 0, 0, 0];
    for line in input.lines() {
        if !line.is_empty() {
            let num: Vec<i32> = re
                .find_iter(line)
                .filter_map(|mat| mat.as_str().parse::<i32>().ok())
                .collect();

            let (x, y) = pos_at_time(num, TIME);
            if x < HALF_WIDTH && y < HALF_HEIGHT {
                quadrants[0] += 1;
            } else if x < HALF_WIDTH && y > HALF_HEIGHT {
                quadrants[1] += 1;
            } else if x > HALF_WIDTH && y < HALF_HEIGHT {
                quadrants[2] += 1;
            } else if x > HALF_WIDTH && y > HALF_HEIGHT {
                quadrants[3] += 1;
            }
        }
    }

    Some(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])
}

pub fn part_two(input: &str) -> Option<u32> {
    let re = Regex::new(r"-?[0-9]+").unwrap();

    let mut robots = Vec::new();

    for line in input.lines() {
        if !line.is_empty() {
            let num: Vec<i32> = re
                .find_iter(line)
                .filter_map(|mat| mat.as_str().parse::<i32>().ok())
                .collect();

            robots.push(num);
        }
    }

    let mut min_count = 999999;
    let mut min_time = 0;
    let mut time = 0;
    loop {
        if time == 10500 {
            break
        }

        let mut count = 0;
        for robot in robots.clone() {
            let pos = pos_at_time(robot, time);
            count += HALF_WIDTH.abs_diff(pos.0) + HALF_HEIGHT.abs_diff(pos.1);
            if count > min_count {
                break
            }
        }

        if count < min_count {
            min_count = count;
            min_time = time;
        }

        time += 1;
    }


    let mut positions = HashSet::new();
    for robot in robots.clone() {
        let pos = pos_at_time(robot, min_time);
        positions.insert(pos);
    }
    print_positions(&positions);

    Some(min_time as u32)
}

fn print_positions(positions: &HashSet<(u32, u32)>) {
    let mut index = 0;
    for y in 0..HEIGHT {
        for x in 0..WIDTH {
            if positions.contains(&(x, y)) {
                print!("x");
                index += 1;
            } else {
                print!(".")
            }
        }
        println!();
    }
    println!();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(12));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
