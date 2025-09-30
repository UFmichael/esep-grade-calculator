package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 71, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 75, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 61, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 65, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 67, Assignment)
	gradeCalculator.AddGrade("exam 1", 21, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 36, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeCutoffs(t *testing.T) {
	testCases := []struct {
		name        string
		assignments []struct {
			name      string
			score     int
			gradeType GradeType
		}
		expectedGrade string
	}{
		{
			name: "A grade at 90",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 90, Assignment},
				{"exam 1", 90, Exam},
				{"essay 1", 90, Essay},
			},
			expectedGrade: "A",
		},
		{
			name: "B grade at 89",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 89, Assignment},
				{"exam 1", 90, Exam},
				{"essay 1", 90, Essay},
			},
			expectedGrade: "B",
		},
		{
			name: "B grade at 80",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 80, Assignment},
				{"exam 1", 80, Exam},
				{"essay 1", 80, Essay},
			},
			expectedGrade: "B",
		},
		{
			name: "C grade at 79",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 79, Assignment},
				{"exam 1", 80, Exam},
				{"essay 1", 80, Essay},
			},
			expectedGrade: "C",
		},
		{
			name: "C grade at 70",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 70, Assignment},
				{"exam 1", 70, Exam},
				{"essay 1", 70, Essay},
			},
			expectedGrade: "C",
		},
		{
			name: "D grade at 69",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 69, Assignment},
				{"exam 1", 70, Exam},
				{"essay 1", 70, Essay},
			},
			expectedGrade: "D",
		},
		{
			name: "D grade at 60",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 60, Assignment},
				{"exam 1", 60, Exam},
				{"essay 1", 60, Essay},
			},
			expectedGrade: "D",
		},
		{
			name: "F grade at 59",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 59, Assignment},
				{"exam 1", 60, Exam},
				{"essay 1", 60, Essay},
			},
			expectedGrade: "F",
		},
		{
			name: "F grade at 50",
			assignments: []struct {
				name      string
				score     int
				gradeType GradeType
			}{
				{"assignment 1", 50, Assignment},
				{"exam 1", 50, Exam},
				{"essay 1", 50, Essay},
			},
			expectedGrade: "F",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gradeCalculator := NewGradeCalculator()
			for _, assignment := range tc.assignments {
				gradeCalculator.AddGrade(assignment.name, assignment.score, assignment.gradeType)
			}
			actual_value := gradeCalculator.GetFinalGrade()
			if tc.expectedGrade != actual_value {
				t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", tc.expectedGrade, actual_value)
			}
		})
	}

}

func TestStrings(t *testing.T) {
	expected_value1 := "assignment"
	expected_value2 := "exam"
	expected_value3 := "essay"

	actual_value1 := Assignment.String()
	actual_value2 := Exam.String()
	actual_value3 := Essay.String()

	if expected_value1 != actual_value1 {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value1, actual_value1)
	}

	if expected_value2 != actual_value2 {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value2, actual_value2)
	}

	if expected_value3 != actual_value3 {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected_value3, actual_value3)
	}
}

func TestGetGradePass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 87, Assignment)
	gradeCalculator.AddGrade("exam 1", 91, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 96, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeFail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 67, Assignment)
	gradeCalculator.AddGrade("exam 1", 21, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 36, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
