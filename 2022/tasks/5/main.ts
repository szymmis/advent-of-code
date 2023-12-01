import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

export function getNamesOfCratesOnTop(
  input: string,
  crateMover: "9000" | "9001"
) {
  const stacks: string[][] = [];
  const moves: { quantity: number; from: number; to: number }[] = [];

  input.split("\n").forEach((line) => {
    if (line.includes("[")) {
      Array.from(line.matchAll(/[\w]/g)).forEach((match) => {
        const { 0: char, index = -1 } = match;
        const stackIndex = Math.floor((index - 1) / 4);
        const stack = stacks[stackIndex] ?? [];
        stacks[stackIndex] = [char, ...stack];
      });
    } else if (line.includes("move")) {
      const match = line.match(/move (\d+) from (\d+) to (\d+)/);
      if (match) {
        const quantity = Number(match[1]);
        const from = Number(match[2]) - 1;
        const to = Number(match[3]) - 1;
        moves.push({ quantity, from, to });
      }
    }
  });

  moves.forEach(({ quantity, from, to }) => {
    const takenCrates = stacks[from].splice(
      stacks[from].length - quantity,
      quantity
    );
    if (crateMover === "9000") takenCrates.reverse();
    stacks[to].push(...takenCrates);
  });

  return stacks.map((stack) => stack[stack.length - 1]).join("");
}

console.log(
  getNamesOfCratesOnTop(input, "9000"),
  getNamesOfCratesOnTop(input, "9001")
);
