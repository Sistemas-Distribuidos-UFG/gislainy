package main

// StudentNotes is n1, n2, n3 of student
type StudentNotes struct {
	N1 float64 `json:"n1"`
	N2 float64 `json:"n2"`
	N3 float64 `json:"n3"`
}

// CalculateFinalResult return if student is aprroved or disapproved
func CalculateFinalResult(sn StudentNotes) Response {

	var response Response;

	averageN1AndN2 := (sn.N1 + sn.N2) / 2;

	response.Message = "Disapproved"

	if (averageN1AndN2 >= 7) {
		response.Message = "Approved"
	}

	if (averageN1AndN2 >= 3) {
		averageFinal := (averageN1AndN2 + sn.N3) / 2

		if (averageFinal >= 5) {
			response.Message = "Approved"
		}
	}

	return response
}