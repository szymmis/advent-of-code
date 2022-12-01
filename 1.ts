import fs from "fs";

const input = fs.readFileSync("1.input.txt", "utf-8");

console.log(
  input
    .split("\n\n")
    .map((arr) =>
      arr.split("\n").reduce((sum, value) => sum + Number(value), 0)
    )
    .sort((a, b) => b - a)
    .slice(0, 3)
    .reduce((sum, value) => sum + value)
);
