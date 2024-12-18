use regex::Regex;
use std::collections::HashMap;

advent_of_code::solution!(17);

fn combo(operand: u128, registers: &HashMap<char, u128>) -> u128 {
    if operand == 4 {
        return registers[&'A'];
    } else if operand == 5 {
        return registers[&'B'];
    } else if operand == 6 {
        return registers[&'C'];
    }
    operand
}

fn run_program(registers: &HashMap<char, u128>, program: &Vec<u128>) -> Vec<u128> {
    let mut registers = registers.clone();

    let mut result = Vec::new();
    let mut i = 0;
    while i < program.len() - 1 {
        let instruction = program[i];
        let operand = program[i + 1];

        match instruction {
            // adv
            0 => {
                let numerator = registers[&'A'];
                let denominator = 2_u128.pow(combo(operand, &registers) as u32);
                registers.insert('A', numerator / denominator);
            }
            // bxl
            1 => {
                let tmp = registers[&'B'] ^ operand;
                registers.insert('B', tmp);
            }
            // bst
            2 => {
                let tmp = combo(operand, &registers) % 8;
                registers.insert('B', tmp);
            }
            // jnz
            3 => {
                if registers[&'A'] != 0 {
                    i = operand as usize;
                    continue;
                }
            }
            // bxc
            4 => {
                let tmp = registers[&'B'] ^ registers[&'C'];
                registers.insert('B', tmp);
            }
            // out
            5 => {
                let tmp = combo(operand, &registers) % 8;
                result.push(tmp);
            }
            // bdv
            6 => {
                let numerator = registers[&'A'];
                let denominator = 2_u128.pow(combo(operand, &registers) as u32);
                registers.insert('B', numerator / denominator);
            }
            7 => {
                let numerator = registers[&'A'];
                let denominator = 2_u32.pow(combo(operand, &registers) as u32) as u128;
                registers.insert('C', numerator / denominator);
            }
            _ => {
                panic!("Invalid instruction!")
            }
        }

        i += 2;
    }

    result
}

pub fn part_one(input: &str) -> Option<u128> {
    let mut registers = HashMap::new();
    let split: Vec<_> = input.split("\n\n").collect();
    let register_regex = Regex::new(r"([A-Z]): (-?[0-9]+)").unwrap();
    for line in split[0].lines() {
        let cap = register_regex.captures(line).unwrap();
        let register: char = cap[1].parse().expect("Not a char.");
        let number: u128 = cap[2].parse().expect("Not an unsigned integer.");
        registers.insert(register, number);
    }

    let program_regex = Regex::new(r"-?[0-9]+").unwrap();
    let program: Vec<u128> = program_regex
        .find_iter(split[1])
        .filter_map(|mat| mat.as_str().parse::<u128>().ok())
        .collect();

    println!("{:?}", registers);
    println!("{:?}", program);

    let result = run_program(&registers, &program);

    let result_string = result
        .iter()
        .map(|num| num.to_string())
        .collect::<Vec<_>>()
        .join(",");

    println!("{}", result_string);
    Some(1)
}

fn find_register(index: usize, program: &Vec<u128>, registers: HashMap<char, u128>) -> (bool, u128) {
    println!("{} {}", index, registers[&'A']);

    for j in 0..8 {
        let mut new_registers = registers.clone();
        new_registers.entry('A').and_modify(|x| *x += j);
        let result = run_program(&new_registers, &program);
        println!(
            "j {} A {} {:?} {:?}",
            j,
            new_registers[&'A'],
            result,
            &program[(program.len() - index - 1)..program.len()]
        );

        if result.eq(&program[(program.len() - index - 1)..program.len()]) {
            if index == program.len() - 1 {
                return (true, new_registers[&'A']);
            }

            new_registers.entry('A').and_modify(|x| *x *= 8);
            let (found, result) = find_register(index + 1, program, new_registers);
            if found {
                return (found, result);
            }
        }
    }

    (false, 0)
}

pub fn part_two(input: &str) -> Option<u128> {
    let mut registers = HashMap::new();
    let split: Vec<_> = input.split("\n\n").collect();
    let register_regex = Regex::new(r"([A-Z]): (-?[0-9]+)").unwrap();
    for line in split[0].lines() {
        let cap = register_regex.captures(line).unwrap();
        let register: char = cap[1].parse().expect("Not a char.");
        let number: u128 = cap[2].parse().expect("Not an unsigned integer.");
        registers.insert(register, number);
    }

    let program_regex = Regex::new(r"-?[0-9]+").unwrap();
    let program: Vec<u128> = program_regex
        .find_iter(split[1])
        .filter_map(|mat| mat.as_str().parse::<u128>().ok())
        .collect();

    registers.insert('A', 0);
    let (_, result) = find_register(0, &program, registers);
    Some(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    // #[test]
    // fn test_part_one() {
    //     let result = part_one(&advent_of_code::template::read_file("examples", DAY));
    //     assert_eq!(result, None);
    // }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(117440));
    }
}
