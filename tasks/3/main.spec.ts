import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import {
  getTotalProritiesOfBadgeItems,
  getTotalProritiesOfCommonItems,
} from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #3", () => {
  it("Should give total priorities of common item types in each rucksack", () => {
    expect(getTotalProritiesOfCommonItems(input)).toBe(157);
  });

  it("Should give total priorities of item types that define groups", () => {
    expect(getTotalProritiesOfBadgeItems(input)).toBe(70);
  });
});
