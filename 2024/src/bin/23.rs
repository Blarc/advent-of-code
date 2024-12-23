extern crate core;

use std::collections::{HashMap, HashSet};

advent_of_code::solution!(23);

fn build_graph(input: &str) -> HashMap<String, HashSet<String>> {
    let mut graph = HashMap::new();

    for line in input.lines() {
        let nodes: Vec<&str> = line.split('-').collect();
        let (node_a, node_b) = (nodes[0].to_string(), nodes[1].to_string());

        // Add both directions since the graph is undirected
        graph
            .entry(node_a.clone())
            .or_insert_with(HashSet::new)
            .insert(node_b.clone());
        graph
            .entry(node_b)
            .or_insert_with(HashSet::new)
            .insert(node_a);
    }

    graph
}

fn find_triangles(graph: &HashMap<String, HashSet<String>>) -> Vec<Vec<String>> {
    let mut triangles = Vec::new();

    for (u, neighbors_u) in graph {
        for v in neighbors_u {
            // Ensure we only process each pair once
            if u < v {
                if let Some(neighbors_v) = graph.get(v) {
                    // Find the intersection of neighbors of u and v
                    for w in neighbors_u.intersection(neighbors_v) {
                        if u < w && v < w {
                            // Add triangle (avoid duplicates)
                            triangles.push(vec![u.clone(), v.clone(), w.clone()]);
                        }
                    }
                }
            }
        }
    }
    triangles
}

fn bron_kerbosch(
    r: HashSet<String>,
    mut p: HashSet<String>,
    mut x: HashSet<String>,
    graph: &HashMap<String, HashSet<String>>,
    cliques: &mut Vec<HashSet<String>>,
) {
    if p.is_empty() && x.is_empty() {
        cliques.push(r);
        return;
    }

    for v in p.clone() {
        let neighbors_v = graph.get(&v).unwrap();

        let mut r_new = r.clone();
        r_new.insert(v.clone());

        let p_new: HashSet<String> = p.intersection(neighbors_v).cloned().collect();
        let x_new: HashSet<String> = x.intersection(neighbors_v).cloned().collect();

        bron_kerbosch(r_new, p_new, x_new, graph, cliques);

        p.remove(&v);
        x.insert(v);
    }
}

pub fn part_one(input: &str) -> Option<u32> {
    let graph = build_graph(input);
    let triangles: Vec<_> = find_triangles(&graph)
        .into_iter()
        .filter(|triangle| triangle.iter().any(|element| element.starts_with("t")))
        .collect();

    Some(triangles.len() as u32)
}

pub fn part_two(input: &str) -> Option<u32> {
    let graph = build_graph(input);
    let mut cliques: Vec<HashSet<String>> = Vec::new();
    bron_kerbosch(
        HashSet::new(),
        graph.keys().cloned().collect(),
        HashSet::new(),
        &graph,
        &mut cliques,
    );

    let mut max_clique: Vec<String> = cliques
        .iter()
        .max_by_key(|clique| clique.len())
        .unwrap()
        .iter()
        .cloned()
        .collect();

    max_clique.sort();

    println!("{:?}", max_clique.join(","));

    Some(0)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(7));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, None);
    }
}
