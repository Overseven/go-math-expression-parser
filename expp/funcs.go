package expp

import "math"

func mult(vars ...float64) float64        { return vars[0] * vars[1] }
func div(vars ...float64) float64         { return vars[0] / vars[1] }
func pow(vars ...float64) float64         { return math.Pow(vars[0], vars[1]) }
func divReminder(vars ...float64) float64 { return float64(int(vars[0]) % int(vars[1])) }

func sum(vars ...float64) float64 { return vars[0] + vars[1] }
func sub(vars ...float64) float64 { return vars[0] - vars[1] }
