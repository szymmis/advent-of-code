import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

export function getSingleMostCalories(input: string) {
  return input
    .split("\n\n")
    .map((arr) =>
      arr.split("\n").reduce((sum, value) => sum + Number(value), 0)
    )
    .sort((a, b) => b - a)[0];
}

export function getTopThreeMostCalories(input: string) {
  return input
    .split("\n\n")
    .map((arr) =>
      arr.split("\n").reduce((sum, value) => sum + Number(value), 0)
    )
    .sort((a, b) => b - a)
    .slice(0, 3)
    .reduce((sum, value) => sum + value);
}

console.log(getSingleMostCalories(input), getTopThreeMostCalories(input));
