import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import { A, B } from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #n", () => {
  it("Should pass", () => {
    expect(A(input)).toBe(null);
  });

  it("Should pass", () => {
    expect(B(input)).toBe(null);
  });
});
