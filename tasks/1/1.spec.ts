import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";

import { getSingleMostCalories, getTopThreeMostCalories } from "./1";

const input = fs.readFileSync(path.join(__dirname, "./1.sample.txt"), "utf-8");

describe("Task #1", () => {
  it("Should give total calories of elve with the most calories", () => {
    expect(getSingleMostCalories(input)).toBe(24000);
  });

  it("Should give total calories of the top three elves", () => {
    expect(getTopThreeMostCalories(input)).toBe(45000);
  });
});
