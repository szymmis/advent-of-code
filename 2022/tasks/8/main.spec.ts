import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import { getVisibleTreesQuantity, getHighestScenicScore } from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #8", () => {
  it("Should give the number of visible trees", () => {
    expect(getVisibleTreesQuantity(input)).toBe(21);
  });

  it("Should give the highest scenic score of all the trees", () => {
    expect(getHighestScenicScore(input)).toBe(8);
  });
});
