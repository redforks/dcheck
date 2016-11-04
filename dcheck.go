// Package dcheck allow checks delayed before enter Starting life phase.
// Many checks is done in init() funcs, if the checks depends on init() execution order,
// such as some object registered in other init() funcs, use dcheck delays the check
// before Starting phase.
package dcheck

import (
	"log"

	"github.com/redforks/life"

	"github.com/redforks/hal"
	"github.com/redforks/testing/reset"
)

// Checker is a function do the check job, return non nil error if the check failed.
type Checker func() error

var checkers []Checker

// Add a checker, the checker will run before life package Starting phase.
// If any checker failed, dcheck abort the application after dump the error
func Add(checker Checker) {
	checkers = append(checkers, checker)
}

func check() {
	log.Print("[spork/dcheck] Starting dcheck")
	for _, checker := range checkers {
		if err := checker(); err != nil {
			log.Print(err)
			hal.Exit(13)
		}
	}
	log.Print("[spork/dcheck] Complete without error")
}

func init() {
	reset.Register(func() {
		checkers = nil
	}, func() {
		life.RegisterHook("dcheck", 1, life.BeforeStarting, check)
	})
}
