package n1688

func numberOfMatches(n int) int {

	t := n

	var res int

	for {

		if t <= 1 {
			break
		}

		d := t / 2
		r := t % 2

		t = d + r

		res += d
	}

	return res
}
