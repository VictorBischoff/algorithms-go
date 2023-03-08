package recursion

// Rec_fractorial returns the fractorial of the given argument (n) and does so recursivly 
func Rec_fractorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Rec_fractorial(n-1)
}
