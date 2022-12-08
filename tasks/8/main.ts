import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

export function getVisibleTreesQuantity(input: string) {
  const isTreeVisible = (grid: number[][], x: number, y: number) => {
    const treeHeight = grid[y][x];

    let invisibleSides = 0;

    for (let k = x - 1; k >= 0; k--) {
      if (treeHeight <= grid[y][k]) {
        invisibleSides += 1;
        break;
      }
    }
    for (let k = x + 1; k < grid[y].length; k++) {
      if (treeHeight <= grid[y][k]) {
        invisibleSides += 1;
        break;
      }
    }
    for (let k = y - 1; k >= 0; k--) {
      if (treeHeight <= grid[k][x]) {
        invisibleSides += 1;
        break;
      }
    }
    for (let k = y + 1; k < grid.length; k++) {
      if (treeHeight <= grid[k][x]) {
        invisibleSides += 1;
        break;
      }
    }

    return invisibleSides < 4;
  };

  const grid = input.split("\n").map((line) => line.split("").map(Number));
  const edgeTreesCount = grid.length * 2 + (grid[0].length - 2) * 2;

  let visibleTreesCount = 0;
  for (let i = 1; i < grid.length - 1; i++) {
    for (let j = 1; j < grid[i].length - 1; j++) {
      if (isTreeVisible(grid, i, j)) {
        visibleTreesCount++;
      }
    }
  }

  return edgeTreesCount + visibleTreesCount;
}

export function getHighestScenicScore(input: string) {
  const getScenicScore = (grid: number[][], x: number, y: number) => {
    const treeHeight = grid[y][x];

    let score = { left: 0, right: 0, top: 0, bottom: 0 };

    for (let k = x - 1; k >= 0; k--) {
      score.left += 1;
      if (treeHeight <= grid[y][k]) break;
    }
    for (let k = x + 1; k < grid[y].length; k++) {
      score.right += 1;
      if (treeHeight <= grid[y][k]) break;
    }
    for (let k = y - 1; k >= 0; k--) {
      score.top += 1;
      if (treeHeight <= grid[k][x]) break;
    }
    for (let k = y + 1; k < grid.length; k++) {
      score.bottom += 1;
      if (treeHeight <= grid[k][x]) break;
    }

    const { left, right, top, bottom } = score;
    return left * right * top * bottom;
  };

  const grid = input.split("\n").map((line) => line.split("").map(Number));
  let highestScenicScore = 0;

  for (let i = 1; i < grid.length - 1; i++) {
    for (let j = 1; j < grid[i].length - 1; j++) {
      const score = getScenicScore(grid, i, j);
      if (score > highestScenicScore) {
        highestScenicScore = score;
      }
    }
  }

  return highestScenicScore;
}

console.log(getVisibleTreesQuantity(input), getHighestScenicScore(input));
