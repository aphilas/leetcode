func backspaceCompare(s string, t string) bool {
	ss := make([]byte, 0, len(s))
	ts := make([]byte, 0, len(t))

	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
            if (len(ss) > 0) {
                ss = ss[:len(ss)-1]  
            }
		} else {
			ss = append(ss, s[i])
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
            if (len(ts) > 0) {
                ts = ts[:len(ts)-1] 
            }
		} else {
			ts = append(ts, t[i])
		}
	}
    
    if len(ss) != len(ts) {
		return false
	}

	for i := 0; i < len(ss) && i < len(ts); i++ {
		if ss[i] != ts[i] {
			return false
		}
	}

	return true
}