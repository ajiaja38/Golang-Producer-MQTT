package exception

import "log"

func ErrorOnFail(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", err, msg)
	}
}
