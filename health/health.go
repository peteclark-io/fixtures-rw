package health

import "time"

// Runner runs all healthchecks
type Runner interface {
	RunChecks()
}

// CheckResult is the json rep of the check
type CheckResult struct {
	Name        string        `json:"name"`
	Severity    int           `json:"severity"`
	Healthy     bool          `json:"healthy"`
	Description string        `json:"description"`
	Impact      string        `json:"impact"`
	CheckOutput string        `json:"checkOutput"`
	Duration    time.Duration `json:"duration"`
}

// AggregateResult aggregates all health checks.
type AggregateResult struct {
	Application   string        `json:"application"`
	Time          time.Time     `json:"time"`
	OverallHealth bool          `json:"overallHealth"`
	Results       []CheckResult `json:"results"`
}

// Aggregator runs all checks
func Aggregator(checks ...func() CheckResult) func() interface{} {
	return func() interface{} {
		var results []CheckResult
		overall := true
		for _, check := range checks {
			r := check()
			results = append(results, r)
			if !r.Healthy {
				overall = false
			}
		}

		return AggregateResult{"match-rw", time.Now(), overall, results}
	}
}

// Ping does a basic connectivity check
func Ping() func() CheckResult {
	return func() CheckResult {
		return CheckResult{
			Name:        "ping",
			Severity:    1,
			Healthy:     true,
			Description: "Simple ping check.",
			Impact:      "No impact.",
			CheckOutput: "pong",
			Duration:    time.Millisecond * 0,
		}
	}
}