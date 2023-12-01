import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

export function getTotalProritiesOfCommonItems(input: string) {
  return input
    .split("\n")
    .map((str) => [
      str.substring(0, str.length / 2),
      str.substring(str.length / 2),
    ])
    .map(([left, right]) => {
      return (
        left
          .split("")
          .find((char) => right.includes(char))
          ?.charCodeAt(0) ?? -1
      );
    })
    .map((char) => {
      if (char === -1) throw new Error("");
      if (char >= 97) return char - 96;
      else return char - 38;
    })
    .reduce((sum, value) => sum + value);
}

export function getTotalProritiesOfBadgeItems(input: string) {
  return input
    .split("\n")
    .reduce((array, line, index) => {
      const subArray = array[Math.floor(index / 3)];
      array[Math.floor(index / 3)] = subArray ? [...subArray, line] : [line];
      return array;
    }, [] as string[][])
    .map((group) => {
      return (
        group
          .sort((a, b) => a.length - b.length)[0]
          .split("")
          .find((char) => group[1].includes(char) && group[2].includes(char))
          ?.charCodeAt(0) ?? -1
      );
    })
    .map((char) => {
      if (char === -1) throw new Error("");
      if (char >= 97) return char - 96;
      else return char - 38;
    })
    .reduce((sum, value) => sum + value);
}

console.log(
  getTotalProritiesOfCommonItems(input),
  getTotalProritiesOfBadgeItems(input)
);
