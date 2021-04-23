package basic

var Basic int

func CountSalary(basicSalary float32) (clean float32) {
	var salary float32 = basicSalary - (2.5*basicSalary)/100
	return salary
}
