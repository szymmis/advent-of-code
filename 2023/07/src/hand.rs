use std::{collections::HashMap, fmt::Debug};

#[derive(Debug, PartialEq, Eq)]
pub struct Hand {
    pub cards: Vec<Card>,
    pub bid: i32,
    pub power: HandPower,
    pub special_jokers: bool,
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
pub enum HandPower {
    FiveOfKind = 6,
    FourOfKind = 5,
    FullHouse = 4,
    ThreeOfKind = 3,
    TwoPair = 2,
    OnePair = 1,
    HighCard = 0,
}

impl Hand {
    pub fn new(cards: Vec<Card>, bid: i32, special_jokers: bool) -> Self {
        let power = {
            let mut map: HashMap<char, i32> = HashMap::new();

            for Card(card, _) in &cards {
                *map.entry(*card).or_insert(0) += 1;
            }

            if special_jokers && map.keys().len() > 1 {
                if let Some(count) = map.remove(&'J') {
                    *map.values_mut().max().unwrap() += count;
                }
            }

            match map.keys().len() {
                1 => HandPower::FiveOfKind,
                2 => {
                    if *map.values().max().unwrap() == 4 {
                        HandPower::FourOfKind
                    } else {
                        HandPower::FullHouse
                    }
                }
                3 => {
                    if *map.values().max().unwrap() == 3 {
                        HandPower::ThreeOfKind
                    } else {
                        HandPower::TwoPair
                    }
                }
                4 => HandPower::OnePair,
                _ => HandPower::HighCard,
            }
        };

        Self {
            cards,
            bid,
            power,
            special_jokers,
        }
    }
}

impl Ord for Hand {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        match self.power.cmp(&other.power) {
            std::cmp::Ordering::Equal => {
                for (c1, c2) in self.cards.iter().zip(other.cards.iter()) {
                    match c1.cmp(c2) {
                        std::cmp::Ordering::Equal => continue,
                        ord => return ord,
                    }
                }
                panic!("Completely two equal hands!")
            }
            ord => ord,
        }
    }
}

impl PartialOrd for Hand {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

#[derive(Eq)]
pub struct Card(pub char, pub i32);

impl Card {
    pub fn new(char: char, special_jokers: bool) -> Self {
        let power = match char {
            'A' => 14,
            'K' => 13,
            'Q' => 12,
            'J' => {
                if special_jokers {
                    1
                } else {
                    11
                }
            }
            'T' => 10,
            '2'..='9' => char.to_digit(10).unwrap() as i32,
            _ => panic!("Unknown card type: {}", char),
        };

        Self(char, power)
    }

    fn value(&self) -> i32 {
        self.1
    }
}

impl Debug for Card {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.0)
    }
}

impl PartialEq for Card {
    fn eq(&self, other: &Self) -> bool {
        self.value() == other.value()
    }
}

impl Ord for Card {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.value().cmp(&other.value())
    }
}

impl PartialOrd for Card {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}
