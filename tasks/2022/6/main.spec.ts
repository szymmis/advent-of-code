import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import { getAmountOfCharactersBeforeSignal } from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #6", () => {
  it("Should give the amount of characters to be processed before receiving signal", () => {
    expect(getAmountOfCharactersBeforeSignal(input, 4)).toEqual([
      7, 5, 6, 10, 11,
    ]);
  });

  it("Should give the amount of characters before receiving message", () => {
    expect(getAmountOfCharactersBeforeSignal(input, 14)).toEqual([
      19, 23, 23, 29, 26,
    ]);
  });
});
