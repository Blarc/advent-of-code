use regex::Regex;
use std::collections::{HashMap, VecDeque};

advent_of_code::solution!(21);

type Coord = (i32, i32);
const DIRS: [(i32, i32, char); 4] = [(-1, 0, '^'), (0, 1, '>'), (1, 0, 'v'), (0, -1, '<')];

fn create_num_keypad() -> HashMap<Coord, char> {
    let keypad_data = [
        ((0, 0), '7'),
        ((0, 1), '8'),
        ((0, 2), '9'),
        ((1, 0), '4'),
        ((1, 1), '5'),
        ((1, 2), '6'),
        ((2, 0), '1'),
        ((2, 1), '2'),
        ((2, 2), '3'),
        ((3, 1), '0'),
        ((3, 2), 'A'),
    ];
    keypad_data.into_iter().collect()
}

fn create_dir_keypad() -> HashMap<Coord, char> {
    let keypad_data = [
        ((0, 1), '^'),
        ((0, 2), 'A'),
        ((1, 0), '<'),
        ((1, 1), 'v'),
        ((1, 2), '>'),
    ];
    keypad_data.into_iter().collect()
}

// Function to compute all shortest paths with the least direction changes
fn all_shortest_paths_with_least_turns(
    start: Coord,
    map: &HashMap<Coord, char>,
) -> HashMap<char, Vec<String>> {
    let mut paths: HashMap<char, Vec<(String, usize)>> = HashMap::new();
    let mut queue = VecDeque::new();

    // Initialize BFS
    paths.insert(map[&start], vec![(String::new(), 0)]); // Start with an empty path and zero turns
    queue.push_back((start, String::new(), None, 0)); // (curr_pos, path_so_far, last_dir, direction_changes)

    while let Some((current, current_path, last_dir, direction_changes)) = queue.pop_front() {
        for dir in DIRS {
            let neighbor = (current.0 + dir.0, current.1 + dir.1);
            let new_direction = dir.2;

            // If the neighbor is in the map
            if map.contains_key(&neighbor) {
                let neighbor_char = map[&neighbor];

                // Calculate direction changes
                let mut new_direction_changes = direction_changes;
                if let Some(last) = last_dir {
                    if last != new_direction {
                        new_direction_changes += 1;
                    }
                }

                // Build the new path
                let new_path = format!("{}{}", current_path, new_direction);

                if !paths.contains_key(&neighbor_char) {
                    // First time visiting this neighbor
                    paths.insert(
                        neighbor_char,
                        vec![(new_path.clone(), new_direction_changes)],
                    );
                    queue.push_back((
                        neighbor,
                        new_path,
                        Some(new_direction),
                        new_direction_changes,
                    ));
                } else {
                    // Compare with existing paths for this neighbor
                    let current_paths = &paths[&neighbor_char];
                    let current_best_length = current_paths[0].0.len();
                    let current_best_turns = current_paths[0].1;

                    if new_path.len() < current_best_length
                        || (new_path.len() == current_best_length
                            && new_direction_changes < current_best_turns)
                    {
                        // Found a better path, replace all existing paths
                        paths.insert(
                            neighbor_char,
                            vec![(new_path.clone(), new_direction_changes)],
                        );
                        queue.push_back((
                            neighbor,
                            new_path,
                            Some(new_direction),
                            new_direction_changes,
                        ));
                    } else if new_path.len() == current_best_length
                        && new_direction_changes == current_best_turns
                    {
                        // Found an equally good path, add it to the list
                        paths
                            .get_mut(&neighbor_char)
                            .unwrap()
                            .push((new_path.clone(), new_direction_changes));
                        queue.push_back((
                            neighbor,
                            new_path,
                            Some(new_direction),
                            new_direction_changes,
                        ));
                    }
                }
            }
        }
    }

    // Convert to a HashMap<char, Vec<String>> by stripping the usize
    paths
        .into_iter()
        .map(|(key, path_list)| (key, path_list.into_iter().map(|(path, _)| path).collect()))
        .collect()
}

fn create_graph(map: &HashMap<Coord, char>) -> HashMap<char, HashMap<char, Vec<String>>> {
    let mut graph = HashMap::new();
    for (coord, char) in map {
        graph.insert(*char, all_shortest_paths_with_least_turns(*coord, &map));
    }
    graph
}

fn find_shortest_sequence(
    code: String,
    dir_graph: &HashMap<char, HashMap<char, Vec<String>>>,
    num_graph: &HashMap<char, HashMap<char, Vec<String>>>,
    depth: usize,
    max_depth: usize,
    memo: &mut HashMap<String, (String, usize, usize)>
) -> (String, usize) {

    if let Some((cached_code, cached_length, cached_depth)) = memo.get(&code.clone()) {
        if depth > 2 {
            return (cached_code.clone(), cached_length * max_depth)
        }
    }

    if depth == max_depth {
        return (code.clone(), code.len());
    }

    let graph = if depth == 0 { num_graph } else { dir_graph };

    let mut result = 0;
    let mut result_code = String::new();
    let mut pos = 'A';
    for c in code.chars() {
        // Initialize the minimum value to the largest possible value
        let mut min = usize::MAX;
        let mut min_code = String::new();

        // Traverse all paths for the current character
        if let Some(paths) = graph.get(&pos).and_then(|m| m.get(&c)) {
            for path in paths {
                // Calculate the cost of the current path
                let new_code = format!("{}{}", path, 'A');
                let (tmp_code, tmp) = find_shortest_sequence(new_code, dir_graph, num_graph, depth + 1, max_depth, memo);
                if tmp < min {
                    min = tmp;
                    min_code = tmp_code
                }
            }
        }

        pos = c;
        result += min;
        result_code = format!("{}{}", result_code, min_code);
    }

    // memo.insert(code.clone(), (result_code.clone(), result, depth));
    (result_code, result)
}

pub fn part_one(input: &str) -> Option<u32> {
    let dir_keypad = create_dir_keypad();
    let num_keypad = create_num_keypad();

    let dir_graph = create_graph(&dir_keypad);
    let num_graph = create_graph(&num_keypad);

    let num_regex = Regex::new(r"([1-9][0-9]*)").unwrap();
    let mut result = 0;
    for line in input.lines() {
        let (shortest_sequence, _) = find_shortest_sequence(line.to_string(), &dir_graph, &num_graph, 0, 3, &mut HashMap::new());
        let tmp = num_regex.find(line).unwrap().as_str();
        let numeric: u32 = tmp.parse::<u32>().ok().unwrap();
        println!("{}: {:<80} -> {} * {}", line, shortest_sequence, shortest_sequence, numeric);
        result += shortest_sequence.len() as u32 * numeric;
    }

    Some(result)
}

pub fn part_two(input: &str) -> Option<u32> {
    let dir_keypad = create_dir_keypad();
    let num_keypad = create_num_keypad();

    let dir_graph = create_graph(&dir_keypad);
    let num_graph = create_graph(&num_keypad);

    // let (shortest_sequence_0, _) = find_shortest_sequence(String::from('>'), &dir_graph, &num_graph, 1, 10, &mut HashMap::new());
    // println!("{}: {} {}", '>', shortest_sequence_0, shortest_sequence_0.len());


    let mut line = String::from("029A");
    for i in 0..5 {
        let (shortest_sequence_0, _) = find_shortest_sequence(line, &dir_graph, &num_graph, i, i+1, &mut HashMap::new());
        println!("{}: {} {}", i,  shortest_sequence_0, shortest_sequence_0.len());
        line = shortest_sequence_0;
    }

    let a = String::from("v<<A>>^A<A>A<AA>vA^Av<AAA^>A");
    for char in a {
    }

    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(126384));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(126384));
    }
}
