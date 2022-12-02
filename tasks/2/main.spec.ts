import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import {
  getTotalScoreUsingBasicStrategy,
  getTotalScoreUsingSecretStrategy,
} from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #2", () => {
  it("Should give total score of all games", () => {
    expect(getTotalScoreUsingBasicStrategy(input)).toBe(15);
  });

  it("Should give total score of all games using elf strategy", () => {
    expect(getTotalScoreUsingSecretStrategy(input)).toBe(12);
  });
});
