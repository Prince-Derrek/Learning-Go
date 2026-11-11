# CLI-Fin-Calc

CLI-Fin-Calc is an interactive, terminal-based calculator written in Go. It supports basic arithmetic and a few calculations over multiple integers, and keeps calculation history and statistics for the current run.

## Requirements

- Go 1.27.3 or later (the module declares Go 1.27.3)

## Run

From the project directory:

```powershell
go run .
```

To build and run an executable on Windows:

```powershell
go build -o fin-calc.exe .
.\fin-calc.exe
```

The program displays a main menu. Enter the number for the action you want:

1. **Perform Calculation** — open the calculation menu. After each calculation, enter `y` to continue calculating or `n` to return to the main menu.
2. **View Calculation History** — choose an operation to view its records, or choose **10** to view all records.
3. **View Statistics** — choose a statistic to display.
4. **Clear History** — remove all records from the current run.
5. **Exit**

### Calculations

The calculation menu accepts these choices:

| Choice | Operation | Inputs |
| --- | --- | --- |
| 1 | Addition | Two integers |
| 2 | Subtraction | Two integers |
| 3 | Product (multiplication) | Two integers |
| 4 | Division | Two integers |
| 5 | Modulus | Two integers |
| 6 | Percentage | Two integers |
| 7 | Average | A count followed by that many integers |
| 8 | Minimum | A count followed by that many integers |
| 9 | Maximum | A count followed by that many integers |

Division, modulus, and percentage are displayed as numeric results; all entered values are integers. The percentage calculation treats the first integer as the 100% value and calculates what percentage the second integer is of it.

### Statistics

The statistics menu provides:

- Total calculations
- Total additions, subtractions, multiplications, and divisions (each separately)
- Largest result
- Smallest result
- Average of all recorded results

## History and calculation behavior

- History is held in memory and is lost when the program exits. Clearing history removes all records for the current run.
- Each record contains the input operands and the result.
- Division currently divides the larger of the two entered values by the smaller, rather than preserving the order in which the values were entered. Avoid zero as an input to division or as the first value in a percentage calculation.
- Percentage uses integer arithmetic, so fractional percentage values are truncated.
- Enter at least one integer for average, minimum, and maximum. An empty list returns `0` for average; minimum and maximum require a non-empty list.

## Development

Run the package checks from the project directory:

```powershell
go test ./...
```

There are currently no test files. At the time this README was written, `go test ./...` reports a `go vet` warning in `services/input/userInput.go` for a redundant newline in a `fmt.Println` call. The CLI itself was run successfully with `go run .`.
