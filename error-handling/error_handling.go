// Package  provides ...
package erratum

const testVersion = 2

func Use(o ResourceOpener, s string) (err error) {
	r, err := o()
	if err != nil {
		_, ok := err.(TransientError)
		if ok {
			return Use(o, s)
		}
		return
	}
	/**
	A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns A deferred function's arguments are evaluated when the defer statement is evaluated.
	*/
	defer func() {
		/**
		Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
		*/
		if res := recover(); res != nil {
			if f, ok := res.(FrobError); ok {
				r.Defrob(f.defrobTag)
			}
			err = res.(error)
		}
		r.Close()
	}()
	r.Frob(s)
	return
}
