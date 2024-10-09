package main

func findLatestTime(s string) string {
	t := []byte(s)
	if t[0] == '?' {
		if t[1] == '?' || t[1] <= '1' {
			t[0] = '1'
		} else {
			t[0] = '0'
		}
	}
	if t[1] == '?' {
		if t[0] == '1' {
			t[1] = '1'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}
