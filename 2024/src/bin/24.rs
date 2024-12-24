use std::cmp::Ordering;
use std::collections::{HashMap, VecDeque};
use std::hash::{Hash, Hasher};
use tinyjson::format;

advent_of_code::solution!(24);

pub fn part_one(input: &str) -> Option<u64> {
    let input_split: Vec<&str> = input.split("\n\n").collect();

    let mut wires: HashMap<&str, u64> = HashMap::new();
    for line in input_split[0].lines() {
        let wire: Vec<&str> = line.split(": ").collect();
        wires.insert(wire[0], wire[1].parse::<u64>().unwrap());
    }

    let mut lines: VecDeque<&str> = input_split[1].lines().collect();

    while !lines.is_empty() {
        let line = lines.pop_front().unwrap();
        let gate: Vec<&str> = line.split(" ").collect();

        if !wires.contains_key(gate[0]) || !wires.contains_key(gate[2]) {
            lines.push_back(line);
            continue;
        }

        let a = wires[gate[0]];
        let b = wires[gate[2]];

        match gate[1] {
            "AND" => {
                wires.insert(gate[4], a & b);
            }
            "XOR" => {
                wires.insert(gate[4], a ^ b);
            }
            "OR" => {
                wires.insert(gate[4], a | b);
            }
            _ => {}
        }
    }

    // Filter and sort
    let mut results: Vec<(&&str, &u64)> =
        wires.iter().filter(|(k, _v)| k.starts_with("z")).collect();

    // Sort alphabetically by key
    results.sort_by(|a, b| a.0.cmp(b.0));

    // Collect the binary values into a vector in sorted order
    let binary_values: Vec<u64> = results.iter().rev().map(|(_, &v)| v).collect();

    let decimal_value = binary_values.iter().fold(0, |acc, &bit| acc * 2 + bit);

    Some(decimal_value)
}

fn compute_gate(a: u8, b: u8, gate: &String) -> u8 {
    match gate.as_str() {
        "AND" => a & b,
        "XOR" => a ^ b,
        "OR" => a | b,
        _ => panic!("Invalid gate!"),
    }
}

#[derive(Clone, Debug)]
struct Gate {
    output: String,
    a: String,
    b: String,
    function: String,
}

impl PartialEq for Gate {
    fn eq(&self, other: &Self) -> bool {
        self.output == other.output
    }
}

// Eq indicates that `PartialEq` is reflexive, symmetric, and transitive
impl Eq for Gate {}

impl PartialOrd for Gate {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        self.output.partial_cmp(&other.output)
    }
}

impl Ord for Gate {
    fn cmp(&self, other: &Self) -> Ordering {
        self.output.cmp(&other.output)
    }
}

impl Hash for Gate {
    fn hash<H: Hasher>(&self, state: &mut H) {
        self.output.hash(state); // Only hash the `output` field
    }
}

fn parse_wires(input: &str) -> HashMap<String, u8> {
    let mut wires: HashMap<String, u8> = HashMap::new();
    for line in input.lines() {
        let wire: Vec<&str> = line.split(": ").collect();
        let value = wire[1].parse::<u8>().unwrap();
        wires.insert(wire[0].to_string(), value);
    }
    wires
}

fn parse_gates(input: &str) -> (HashMap<String, Gate>, HashMap<String, Vec<Gate>>) {
    let mut gates: HashMap<String, Gate> = HashMap::new();
    let mut gates_map: HashMap<String, Vec<Gate>> = HashMap::new();
    for line in input.lines() {
        // Split the line into components
        let gate: Vec<&str> = line.split(" ").collect();

        // Extract inputs and output
        let a = gate[0].to_string();
        let op = gate[1];
        let b = gate[2].to_string();
        let output = gate[4].to_string();
        let new_gate = Gate {
            output,
            a: a.clone(),
            b: b.clone(),
            function: op.to_string(),
        };

        gates.insert(gate[4].to_string(), new_gate.clone());
        gates_map
            .entry(a)
            .or_insert(Vec::new())
            .push(new_gate.clone());
        gates_map
            .entry(b)
            .or_insert(Vec::new())
            .push(new_gate.clone());
    }
    (gates, gates_map)
}

fn compute(mut wires: HashMap<String, u8>, gates_map: &HashMap<String, Gate>) -> u64 {
    // Convert all gates into a queue for processing
    let mut gates: VecDeque<&Gate> = VecDeque::from(gates_map.values().collect::<Vec<&Gate>>());

    // Process the gates
    while let Some(gate) = gates.pop_front() {
        // Check if both input wires are ready
        let a_ready = wires.get(&gate.a);
        let b_ready = wires.get(&gate.b);

        if a_ready.is_none() || b_ready.is_none() {
            gates.push_back(gate); // Requeue the gate for later processing
            continue;
        }

        // Retrieve the input values
        let a = *a_ready.unwrap();
        let b = *b_ready.unwrap();

        // Compute the result for the gate and store it in `wires`
        let result = compute_gate(a, b, &gate.function);
        wires.insert(gate.output.clone(), result);
    }

    // Convert binary to decimal
    binary_to_decimal('z', &wires)
}

fn binary_to_decimal(name: char, wires: &HashMap<String, u8>) -> u64 {
    // Filter for wires with keys starting with "z" and sort them
    let mut results: Vec<_> = wires
        .iter()
        .filter(|(k, _)| k.starts_with(name)) // Only consider keys starting with "z"
        .collect();
    results.sort_by(|a, b| a.0.cmp(b.0)); // Sort by key (wire name)

    // Convert sorted binary values into decimal
    results
        .iter()
        .rev() // Reverse for most-significant-to-least-significant order
        .map(|(_, &v)| v as u64)
        .fold(0, |acc, bit| acc * 2 + bit) // Convert binary to decimal
}

fn find_output(gate: Gate, gates_map: &HashMap<String, Vec<Gate>>) -> Vec<Gate> {
    if gate.output.starts_with("z") {
        return vec![gate];
    }

    let mut result = Vec::new();
    let next_gates = gates_map.get(&gate.output).unwrap();
    for next_gate in next_gates {
        let tmp = find_output(next_gate.clone(), gates_map);
        result.extend(tmp);
    }
    result
}

fn decrement_z_gate_output(output: &String) -> Option<String> {
    if output.starts_with('z') {
        if let Ok(num) = output[1..].parse::<u32>() {
            return Some(format!("z{:02}", num.saturating_sub(1)));
        }
    }
    None
}

pub fn part_two(input: &str) -> Option<u64> {
    let input_split: Vec<&str> = input.split("\n\n").collect();

    let wires = parse_wires(input_split[0]);
    let (mut gates, gates_map) = parse_gates(input_split[1]);

    let mut first = Vec::new();
    let mut second = Vec::new();

    for gate in gates.values() {
        if gate.function != "XOR" && gate.output.starts_with("z") && gate.output != "z45" {
            first.push(gate.clone());
            println!("not XOR: {:?}", gate);
        } else if gate.function == "XOR"
            && !gate.a.starts_with("x")
            && !gate.a.starts_with("y")
            && !gate.b.starts_with("x")
            && !gate.b.starts_with("y")
            && !gate.output.starts_with("z")
        {
            second.push(gate.clone());
            println!("yes XOR: {:?}", gate);
        }
    }

    let mut swaps: Vec<(String, String)> = Vec::new();
    let mut new_input = input.clone().to_string();
    for a_gate in second {
        let mut z_outputs = find_output(a_gate.clone(), &gates_map);
        z_outputs.sort();
        let z_outputs_strings: Vec<String> = z_outputs.iter().map(|x| x.output.clone()).collect();
        println!("{:?} {:?}", a_gate, z_outputs_strings);

        let z_gate = z_outputs.first().unwrap();
        let z_gate_output_decremented = decrement_z_gate_output(&z_gate.output).unwrap();

        swaps.push((a_gate.output.clone(), z_gate_output_decremented.clone()));

        println!("Swapping --> {} with --> {}", a_gate.output, "placeholder");
        new_input = new_input.replace(
            format!("-> {}", a_gate.output).as_str(),
            format!("-> {}", "placeholder").as_str(),
        );

        println!("Swapping --> {} with --> {}", z_gate_output_decremented, a_gate.output);
        new_input = new_input.replace(
            format!("-> {}", z_gate_output_decremented).as_str(),
            format!("-> {}", a_gate.output).as_str(),
        );

        println!("Swapping --> {} with --> {}", "placeholder", z_gate_output_decremented);
        new_input = new_input.replace(
            format!("-> {}", "placeholder").as_str(),
            format!("-> {}", z_gate_output_decremented).as_str(),
        );


        // Not sure why this does not work...
        // gates.remove(&a_gate.output);
        // gates.remove(&z_gate_output_decremented);
        //
        // gates.insert(
        //     z_gate_output_decremented.clone(),
        //     Gate {
        //         output: z_gate_output_decremented.clone(),
        //         a: a_gate.a.clone(),
        //         b: a_gate.b.clone(),
        //         function: a_gate.function.clone(),
        //     },
        // );
        //
        // gates.insert(
        //     a_gate.output.clone(),
        //     Gate {
        //         output: a_gate.output.clone(),
        //         a: z_gate.a.clone(),
        //         b: z_gate.b.clone(),
        //         function: z_gate.function.clone(),
        //     },
        // );
    }






    let leading_zeros = 33;

    let last_gate = gates_map["x33"].clone();


    swaps.push((last_gate[0].output.clone(), last_gate[1].output.clone()));
    println!("Swapping --> {} with --> {}", last_gate[0].output, "placeholder");
    new_input = new_input.replace(
        format!("-> {}", last_gate[0].output).as_str(),
        format!("-> {}", "placeholder").as_str(),
    );

    println!("Swapping --> {} with --> {}", last_gate[1].output, last_gate[0].output);
    new_input = new_input.replace(
        format!("-> {}", last_gate[1].output).as_str(),
        format!("-> {}", last_gate[0].output).as_str(),
    );

    println!("Swapping --> {} with --> {}", "placeholder", last_gate[1].output);
    new_input = new_input.replace(
        format!("-> {}", "placeholder").as_str(),
        format!("-> {}", last_gate[1].output).as_str(),
    );

    let result = part_one(new_input.as_str()).unwrap();
    let x = binary_to_decimal('x', &wires);
    let y = binary_to_decimal('y', &wires);
    let sum = x + y;

    println!("{:b}", sum);
    println!("{:b}", result);
    // 1111000000000000000000000000000000000
    println!("{:b}", result ^ sum);


    // println!("{}", compute(wires, &gates));



    println!("{:?}", swaps);

    let mut last_vec = Vec::new();
    for (a, b) in swaps {
        last_vec.push(a);
        last_vec.push(b);
    }

    last_vec.sort();

    println!("{:?}",  last_vec.join(","));


    // fgc,z12,mtj,z29,dtv,z37,vvm,dgr
    Some(0)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(2024));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
