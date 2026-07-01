package multiverb

const name = "multiverb"

// PlanCommand is a verb constructor; a package exposing several verbs
// satisfies the entry-point standard without a bare Command().
func PlanCommand() string { return name + " plan" }

// ApplyCommand is the second verb constructor.
func ApplyCommand() string { return name + " apply" }

// helperCommand is unexported and must not satisfy the entry point.
func helperCommand() string { return name }
