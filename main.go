package main

import (
	"fmt"
)

// School struct represents an international school
type School struct {
	Name         string
	Curriculum   string
	Reputation   int // Rating out of 5
	FeeStructure int // Annual fee in INR
	Facilities   string
}

// AdmissionResult struct to hold admission process outcome
type AdmissionResult struct {
	SchoolName      string
	TestPassed      bool
	InterviewPassed bool
	FinalFee        int
	Admitted        bool
}

// Shortlist schools based on reputation and curriculum
func shortlistSchools() []School {
	return []School{
		{"Global International School", "IB", 5, 300000, "Sports, Labs, Music"},
		{"Elite World Academy", "IGCSE", 4, 250000, "Swimming, Robotics, Arts"},
		{"Cambridge Global", "CBSE International", 4, 200000, "Library, Theatre, Labs"},
	}
}

// Simulate the admission process
func processAdmission(school School, budget int) AdmissionResult {
	testPassed := school.Reputation >= 4 // Assume schools with good reputation conduct strict tests
	interviewPassed := testPassed        // Assume interview success follows test clearance
	finalFee := school.FeeStructure      // No discount applied
	admitted := interviewPassed && finalFee <= budget

	return AdmissionResult{
		SchoolName:      school.Name,
		TestPassed:      testPassed,
		InterviewPassed: interviewPassed,
		FinalFee:        finalFee,
		Admitted:        admitted,
	}
}

func main() {
	budget := 250000 // Set budget for admission
	schools := shortlistSchools()
	var finalDecision AdmissionResult

	fmt.Println("\n🔹 Shortlisted Schools for Admission:")
	for _, school := range schools {
		fmt.Printf("- %s (%s), Reputation: %d, Fee: ₹%d, Facilities: %s\n",
			school.Name, school.Curriculum, school.Reputation, school.FeeStructure, school.Facilities)
	}

	fmt.Println("\n🔹 Processing Admissions:")
	for _, school := range schools {
		result := processAdmission(school, budget)
		fmt.Printf("\nProcessing %s:\n", school.Name)
		fmt.Printf("  Entrance Test Passed: %t\n  Interview Passed: %t\n  Final Fee: ₹%d\n  Admission Granted: %t\n",
			result.TestPassed, result.InterviewPassed, result.FinalFee, result.Admitted)
		if result.Admitted {
			finalDecision = result
			break // Stop at first successful admission
		}
	}

	if finalDecision.Admitted {
		fmt.Printf("\n✅ Final Decision: Joining %s with Fee ₹%d\n", finalDecision.SchoolName, finalDecision.FinalFee)
	} else {
		fmt.Println("\n❌ No school selected due to budget constraints or failed admission process.")
	}
}
