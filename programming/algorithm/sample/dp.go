package sample

// ClimbingStairs problem
// You are climbing a stair case. It takes n steps to reach to the top.  
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?  
func ClimbingStairs(n int) int {
	// fibnacci
	a, b := 1, 2
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
