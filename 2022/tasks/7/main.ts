import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

const getDirectorySizes = (input: string) => {
  const sizes: Record<string, number> = {};
  const directoryStack: string[] = [];

  input.split("\n").forEach((line) => {
    if (line.includes("$")) {
      const directory = line.match(/\$ cd (.+)/)?.[1];
      if (directory) {
        if (directory === "..") {
          const currentDir = directoryStack.join(".")!;
          directoryStack.pop();
          const previousDir = directoryStack.join(".")!;
          sizes[previousDir] =
            (sizes[previousDir] ?? 0) + (sizes[currentDir] ?? 0);
        } else {
          directoryStack.push(directory);
        }
      }
    } else {
      if (line.includes("dir")) return;
      const currentDir = directoryStack.join(".")!;
      const fileSize = Number(line.match(/\d+/)?.[0] ?? 0);
      sizes[currentDir] = (sizes[currentDir] ?? 0) + fileSize;
    }
  });

  while (directoryStack.length > 1) {
    const currentDir = directoryStack.join(".")!;
    directoryStack.pop();
    const previousDir = directoryStack.join(".")!;
    sizes[previousDir] = (sizes[previousDir] ?? 0) + (sizes[currentDir] ?? 0);
  }

  return sizes;
};

export function getSumOfDirectorySizes(input: string) {
  return Object.values(getDirectorySizes(input)).reduce(
    (sum, size) => (size <= 100000 ? sum + size : sum),
    0
  );
}

export function getSizeOfDirectoryToDelete(input: string) {
  const TOTAL_DISK_SPACE = 70000000;
  const UPDATE_DISK_SPACE = 30000000;
  const sizes = getDirectorySizes(input);
  const unusedSpace = TOTAL_DISK_SPACE - sizes["/"];
  const spaceNeeded = UPDATE_DISK_SPACE - unusedSpace;

  return Object.values(sizes)
    .filter((size) => size >= spaceNeeded)
    .sort((a, b) => a - b)[0];
}

console.log(getSumOfDirectorySizes(input), getSizeOfDirectoryToDelete(input));
