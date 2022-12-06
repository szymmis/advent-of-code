import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

export function getAmountOfCharactersBeforeSignal(
  input: string,
  uniqueCharsNeeded: number
) {
  return input.split("\n").map((data) => {
    const characters = data.split("");
    const occurencesMap: Record<string, number> = {};
    for (let i = 0; i < characters.length; i++) {
      const char = characters[i];
      occurencesMap[char] = (occurencesMap[char] ?? 0) + 1;
      if (i < uniqueCharsNeeded - 1) {
        continue;
      }

      for (let j = 0; j < uniqueCharsNeeded; j++) {
        if (occurencesMap[characters[i - j]] > 1) break;
        if (j === uniqueCharsNeeded - 1) return i + 1;
      }

      occurencesMap[characters[i - (uniqueCharsNeeded - 1)]] -= 1;
    }
  });
}

console.log(
  getAmountOfCharactersBeforeSignal(input, 4),
  getAmountOfCharactersBeforeSignal(input, 14)
);
