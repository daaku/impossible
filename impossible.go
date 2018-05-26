// Package impossible converts errors into panics. It is meant to be used when
// the error is impossible.
package impossible

// Error panics with the error if it is not nil.
func Error(err error) {
	if err != nil {
		panic(err)
	}
}
