import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

export function getFullyContainedPairsCount(input: string) {
  return input
    .split("\n")
    .map((str) => str.split(",").map((str) => str.split("-").map(Number)))
    .reduce((sum, [[x1, x2], [x3, x4]]) => {
      return (x1 >= x3 && x2 <= x4) || (x3 >= x1 && x4 <= x2) ? sum + 1 : sum;
    }, 0);
}

export function getOverlappingPairsCount(input: string) {
  return input
    .split("\n")
    .map((str) => str.split(",").map((str) => str.split("-").map(Number)))
    .reduce((sum, [[x1, x2], [x3, x4]]) => {
      return x1 <= x4 && x2 >= x3 ? sum + 1 : sum;
    }, 0);
}

console.log(
  getFullyContainedPairsCount(input),
  getOverlappingPairsCount(input)
);
