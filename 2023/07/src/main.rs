use crate::hand::{Card, Hand};

mod hand;

fn calculate_hands(input: &str, special_jokers: bool) -> i32 {
    let mut hands: Vec<Hand> = input
        .lines()
        .map(|line| {
            let [cards, bid] = line.split(' ').collect::<Vec<_>>()[..] else {
                panic!()
            };
            let cards: Vec<Card> = cards
                .chars()
                .map(|c| Card::new(c, special_jokers))
                .collect::<Vec<_>>();
            let bid: i32 = bid.parse().unwrap();

            Hand::new(cards, bid, special_jokers)
        })
        .collect();

    hands.sort();

    hands
        .iter()
        .enumerate()
        .fold(0, |acc, (i, hand)| acc + (i + 1) as i32 * hand.bid)
}

fn main() {
    let input = std::fs::read_to_string("input.txt").unwrap();
    println!("Part one: {}", calculate_hands(&input, false));
    println!("Part two: {}", calculate_hands(&input, true))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_one_example() {
        let input = "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483";
        assert_eq!(calculate_hands(input, false), 6440);
    }

    #[test]
    fn part_two_example() {
        let input = "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483";
        assert_eq!(calculate_hands(input, true), 5905);
    }
}
