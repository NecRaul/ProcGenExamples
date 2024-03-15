import java.util.Arrays;
import java.util.ArrayList;
import java.util.List;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.Files;

public class GameOfLife {
    public static void main(String[] args) {
        Path filePath = Paths.get("state_grid.txt");
        List<String> stateGrid = new ArrayList<>();
        try {
            stateGrid = Files.readAllLines(filePath);
        } catch (Exception e) {
            e.printStackTrace();
        }
        boolean[][] booleanArray = new boolean[102][102];
        boolean[][] updatedArray = new boolean[102][102];
        for (int i = 0; i < stateGrid.size(); i++) {
            String stateLine = stateGrid.get(i);
            for (int j = 0; j < stateGrid.size(); j++) {
                Character stateCell = stateLine.charAt(j);
                booleanArray[i + 1][j + 1] = stateCell == '#';
            }
        }
        updatedArray = Arrays.stream(booleanArray).map(boolean[]::clone).toArray(boolean[][]::new);
        for (int iter = 0; iter < 100; iter++) {
            for (int i = 1; i < booleanArray.length - 1; i++) {
                for (int j = 1; j < booleanArray.length - 1; j++) {
                    int count = 0;
                    for (int k = -1; k <= 1; k++) {
                        for (int l = -1; l <= 1; l++) {
                            if ((k != 0 || l != 0) && booleanArray[i + k][j + l]) {
                                count++;
                            }
                        }
                    }
                    if (booleanArray[i][j] && !(count == 2 || count == 3)) {
                        updatedArray[i][j] = false;
                    } else if (!booleanArray[i][j] && count == 3) {
                        updatedArray[i][j] = true;
                    }
                }
            }
            booleanArray = Arrays.stream(updatedArray).map(boolean[]::clone).toArray(boolean[][]::new);
            if ((iter + 1) % 10 == 0) {
                int count = countAliveCells(booleanArray);
                System.out.println("Iteration " + (iter + 1) + ": " + count);
            }
        }
    }

    private static int countAliveCells(boolean[][] booleanArray) {
        int count = 0;
        for (int i = 0; i < booleanArray.length; i++) {
            for (int j = 0; j < booleanArray.length; j++) {
                if (booleanArray[i][j]) {
                    count++;
                }
            }
        }
        return count;
    }
}
