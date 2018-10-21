package algo

// Sieve is an implementation of Sieve of Eratosthenes.
// Sieve returns slice of all primes numbers that are lower than given number.
func Sieve(n int) []int {
	if n < 3 {
		return []int{}
	}

	numbers := make([]bool, n)
	numbers[0] = true
	numbers[1] = true

	for i := 2; i*i <= n; i++ {
		if !numbers[i] {
			for j := i * i; j < n; j += i {
				numbers[j] = true
			}
		}
	}

	var primes []int
	for i := 0; i < n; i++ {
		if !numbers[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
