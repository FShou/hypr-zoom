package main

import (
	"flag"
	"fmt"
	"github.com/fogleman/ease"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type EasingFunction func(t float64) float64

// Linear interpolation function
func lerp(start, end, t float64) float64 {
	return start + t*(end-start)
}

// Main animation loop
func zoom(duration int, steps int, startValue, endValue float64, easingFunc EasingFunction) {
	interval := float64(duration) / float64(steps) / 1000

	for i := 0; i <= steps; i++ {
		t := float64(i) / float64(steps)
		easedT := easingFunc(t)
		interpolatedValue := lerp(startValue, endValue, easedT)

		_ = exec.Command("hyprctl", "keyword", "cursor:zoom_factor", fmt.Sprintf("%f", interpolatedValue)).Run()
		time.Sleep(time.Duration(interval * float64(time.Second)))
	}
}

// Execute a shell command and return its output as a string
func executeCommand(command string) (string, error) {
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {
	output, err := executeCommand("hyprctl getoption cursor:zoom_factor | grep 'float:' | awk '{print $2}' | tr -d '[:space:]'")
	if err != nil {
		fmt.Println("Error getting initial zoom factor:", err)
		return
	}

	duration := flag.Int("duration", 500, "Duration of the animation in milliseconds")
	steps := flag.Int("steps", 100, "Number of steps in the animation")
	easing := flag.String("easing", "InOutExpo", "Easing function to use (easeInOutCubic)")
	easingOut := flag.String("easingOut", "", "Easing function to use (easeInOutCubic)")
	targetZoom := flag.Float64("target", 2.0, "Zoom Target")
	flag.Parse()

	initialZoom, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
	if err != nil {
		fmt.Println("Error parsing initial zoom factor:", err)
		return
	}

	easingFunctions := map[string]EasingFunction{
		"Linear":     ease.Linear,
		"InQuad":     ease.InQuad,
		"OutQuad":    ease.OutQuad,
		"InOutQuad":  ease.InOutQuad,
		"InCubic":    ease.InCubic,
		"OutCubic":   ease.OutCubic,
		"InOutCubic": ease.InOutCubic,
		"InQuart":    ease.InQuart,
		"OutQuart":   ease.OutQuart,
		"InOutQuart": ease.InOutQuart,
		"InQuint":    ease.InQuint,
		"OutQuint":   ease.OutQuint,
		"InOutQuint": ease.InOutQuint,
		"InSine":     ease.InSine,
		"OutSine":    ease.OutSine,
		"InOutSine":  ease.InOutSine,
		"InExpo":     ease.InExpo,
		"OutExpo":    ease.OutExpo,
		"InOutExpo":  ease.InOutExpo,
		"InCirc":     ease.InCirc,
		"OutCirc":    ease.OutCirc,
		"InOutCirc":  ease.InOutCirc,
		// "InElastic": ease.InElastic,
		// "OutElastic": ease.OutElastic,
		// "InOutElastic": ease.InOutElastic,
		"InBack":      ease.InBack,
		"OutBack":     ease.OutBack,
		"InOutBack":   ease.InOutBack,
		"InBounce":    ease.InBounce,
		"OutBounce":   ease.OutBounce,
		"InOutBounce": ease.InOutBounce,
		"InSquare":    ease.InSquare,
		"OutSquare":   ease.OutSquare,
		"InOutSquare": ease.InOutSquare,
	}
	easingFunction, exists := easingFunctions[*easing]
	if !exists {
		fmt.Println("Unknown easing function:", *easing)
		return
	}

	if initialZoom > 1 {
		*targetZoom = 1.0
		_, exists := easingFunctions[*easingOut]
		if exists {
			easingFunction = easingFunctions[*easingOut]
		}
	}

	zoom(*duration, *steps, initialZoom, *targetZoom, easingFunction)
}
