package funcs

// FuncType - internal type of functions
type FuncType func(args ...float64) (float64, error)

// count of operator priorities
const LevelsOfPriorities = 3
