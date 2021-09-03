var regexGroup singleflight.Group

// RegexMatch reports whether the string s contains any match of the regular expression pattern.
func RegexMatch(pat, s string) (bool, error) {
	val, err, _ := regexGroup.Do(pat, func() (interface{}, error) {
		return regexp.Compile(pat)
	})
	if err != nil {
		return false, err
	}
	return val.(*regexp.Regexp).MatchString(s), nil
}
