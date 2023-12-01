import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import { getFullyContainedPairsCount, getOverlappingPairsCount } from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #4", () => {
  it("Should give amount of pairs where one range fully contain the other", () => {
    expect(getFullyContainedPairsCount(input)).toBe(2);
  });

  it("Should give amount of pairs that overlap with each other", () => {
    expect(getOverlappingPairsCount(input)).toBe(4);
  });
});
