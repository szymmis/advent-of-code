import fs from "fs";
import path from "path";

const input = fs.readFileSync(path.join(__dirname, "./input.txt"), "utf-8");

export function getTotalScoreUsingBasicStrategy(input: string) {
  const games = input
    .split("\n")
    .map((str) => {
      const [opponentToken, playerToken] = str.split(" ");
      return [opponentToken.charCodeAt(0) - 65, playerToken.charCodeAt(0) - 88];
    })
    .map(([opponent, player]) => {
      const figureWeight = player + 1;
      if (opponent === player) return figureWeight + 3;
      else {
        if (opponent % 3 === (player - 1 + 3) % 3) return figureWeight + 6;
        return figureWeight + 0;
      }
    });

  return games.reduce((sum, gamePoints) => sum + gamePoints);
}

export function getTotalScoreUsingSecretStrategy(input: string) {
  const games = input
    .split("\n")
    .map((str) => {
      const [opponentToken, playerStrategy] = str.split(" ");

      const opponentFigure = opponentToken.charCodeAt(0) - 65;
      let playerFigure;

      switch (playerStrategy) {
        case "X":
          playerFigure = (opponentFigure - 1 + 3) % 3;
          break;
        case "Y":
          playerFigure = opponentFigure;
          break;
        case "Z":
        default:
          playerFigure = (opponentFigure + 1) % 3;
      }

      return [opponentFigure, playerFigure];
    })
    .map(([opponent, player]) => {
      const figureWeight = player + 1;
      if (opponent === player) return figureWeight + 3;
      else {
        if (opponent % 3 === (player - 1 + 3) % 3) return figureWeight + 6;
        return figureWeight + 0;
      }
    });

  return games.reduce((sum, gamePoints) => sum + gamePoints);
}

console.log(getTotalScoreUsingSecretStrategy(input));
