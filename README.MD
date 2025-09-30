# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

# Files
- 'grade_calculator.go' - Main implementation containing the 'GradeCalculator' struct and methods for adding grades and calculating final letter grades
- 'grade_calculator_test.go' - Unit tests that verify the grade calculator works correctly
- 'go.mod' - Go module file that defines the Go version requirement

## How Grades Are Calculated

1. Calculate the average of all assignments
2. Calculate the average of all exams  
3. Calculate the average of all essays
4. Apply weights: '(assignments × 0.50) + (exams × 0.35) + (essays × 0.15)'
5. Convert the numerical result to a letter grade

'Example:'
- Assignments average: 85
- Exams average: 90
- Essays average: 88
- Final: (85 × 0.50) + (90 × 0.35) + (88 × 0.15) = 87.2 → 'B'