advent_of_code::solution!(9);

#[derive(Clone, Debug)]
struct Node {
    id: u32,
    size: u32,
    is_empty: bool,
}

pub fn part_one(input: &str) -> Option<u64> {
    let mut disk: Vec<i32> = Vec::new();

    let mut index: i32 = 0;
    let mut is_empty = false;
    input
        .chars()
        .filter(|x| x.is_digit(10))
        .map(|x| x.to_digit(10).expect("Character is not a digit!") as i32)
        .for_each(|x| {
            for _ in 0..x {
                if is_empty {
                    disk.push(-1);
                } else {
                    disk.push(index);
                }
            }

            if !is_empty {
                index += 1;
            }
            is_empty = !is_empty
        });

    let mut empty_index: usize = 0;
    let mut end_index: usize = disk.len() - 1;

    loop {
        while disk[empty_index] != -1 {
            empty_index += 1;
        }

        if empty_index > end_index {
            break;
        }

        disk[empty_index] = disk[end_index];
        disk[end_index] = -1;
        end_index -= 1;
    }

    // for x in &disk {
    //     if *x == -1 {
    //         print!(".");
    //     } else {
    //         print!("{},", x);
    //     }
    // }
    // println!();

    // 0099811188827773336446555566
    // 0099811188827773336446555566

    let sum: u64 = disk
        .iter()
        .filter(|&x| *x != -1)
        .enumerate()
        .map(|(i, x)| i as u64 * *x as u64)
        .sum();

    // 2339331266
    Some(sum)
}

fn print_disk(disk: &Vec<Node>) {
    for i in 0..disk.len() {
        for _ in 0..disk[i].size {
            if disk[i].is_empty {
                print!(".");
            } else {
                print!("{}", disk[i].id);
            }
        }
    }
    println!();
}

pub fn part_two(input: &str) -> Option<u64> {
    let mut disk: Vec<Node> = Vec::new();

    let mut is_empty = false;

    let mut index = 0;
    input
        .chars()
        .filter(|x| x.is_digit(10))
        .map(|x| x.to_digit(10).expect("Character is not a digit!"))
        .for_each(|x| {
            if is_empty {
                disk.push(Node {
                    id: index,
                    size: x,
                    is_empty: true,
                });
            } else {
                disk.push(Node {
                    id: index,
                    size: x,
                    is_empty: false,
                });
                index += 1;
            }
            is_empty = !is_empty
        });

    let mut node_index = disk.len() - 1;
    while node_index > 0 {
        let node = &disk[node_index];
        if !node.is_empty {
            for empty_index in 0..disk.len() {
                if empty_index > node_index {
                    continue
                }

                let empty = &disk[empty_index];
                if empty.is_empty {
                    // println!("{}:{} {}", node.id, node.size, empty.size);
                    if node.size == empty.size {
                        disk.swap(empty_index, node_index);
                        // print_disk(&disk);
                        break;
                    }
                    if node.size < empty.size {
                        let new_empty_size = empty.size - node.size;
                        let new_empty_size_2 = node.size;

                        disk.swap(empty_index, node_index);
                        disk.remove(node_index);

                        disk.insert(empty_index+1, Node {
                            id: 0,
                            size: new_empty_size,
                            is_empty: true,
                        });

                        disk.insert(node_index, Node {
                            id: 0,
                            size: new_empty_size_2,
                            is_empty: true
                        });

                        // print_disk(&disk);
                        break;
                    }
                }
            }
        }
        node_index -= 1;
    }

    let mut index: u64 = 0;
    let mut sum: u64 = 0;
    for i in 0..disk.len() {
        for _ in 0..disk[i].size {
            if !disk[i].is_empty {
                sum += index * disk[i].id as u64;
            }
            index += 1;
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
        assert_eq!(result, Some(1928));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(2858));
    }
}
