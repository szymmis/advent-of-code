import fs from "fs";
import path from "path";
import { expect } from "@jest/globals";
import { getSumOfDirectorySizes, getSizeOfDirectoryToDelete } from "./main";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

describe("Task #7", () => {
  it("Should give total size sum of directories of size less than 100000", () => {
    expect(getSumOfDirectorySizes(input)).toBe(95437);
  });

  it("Should give the total size of directory to be deleted", () => {
    expect(getSizeOfDirectoryToDelete(input)).toBe(24933642);
  });
});
