import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./sample.txt"), "utf-8");

export function A(input: string) {}
export function B(input: string) {}

console.log(A(input), B(input));
