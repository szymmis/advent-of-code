import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import { getNamesOfCratesOnTop } from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #5", () => {
  it("Should give names of crates on top moved using CrateMover 9000", () => {
    expect(getNamesOfCratesOnTop(input, "9000")).toBe("CMZ");
  });

  it("Should give names of crates on top moved using CrateMover 9001", () => {
    expect(getNamesOfCratesOnTop(input, "9001")).toBe("MCD");
  });
});
