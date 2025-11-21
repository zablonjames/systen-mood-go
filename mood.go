package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type SystemMood struct {
	CPUUsage    int
	MemoryUsage int
	Mood        string
	Message     string
}

func main() {
	fmt.Println("🖥️  System Mood Monitor Starting...")
	fmt.Println("=====================================\n")

	// Monitor the system every 3 seconds
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	count := 0
	for range ticker.C {
		mood := checkSystemMood()
		displayMood(mood)
		
		count++
		if count >= 5 {
			fmt.Println("\n✅ Monitoring complete! Your system says thanks!")
			break
		}
		
		fmt.Println("\n---")
	}
}

func checkSystemMood() SystemMood {
	// Simulate CPU and memory usage (in real app, you'd use actual system metrics)
	cpuUsage := rand.Intn(100)
	memUsage := rand.Intn(100)
	
	var mood string
	var message string
	
	// Determine mood based on usage
	if cpuUsage < 30 && memUsage < 40 {
		mood = "😊 Relaxed"
		message = "System is chilling! Time to push some code!"
	} else if cpuUsage < 60 && memUsage < 70 {
		mood = "💪 Productive"
		message = "Good balance! Keep the momentum going!"
	} else if cpuUsage < 85 && memUsage < 85 {
		mood = "😅 Busy"
		message = "Working hard! Maybe close a few tabs?"
	} else {
		mood = "🔥 Stressed"
		message = "Whoa! Your system needs a break. Save your work!"
	}
	
	return SystemMood{
		CPUUsage:    cpuUsage,
		MemoryUsage: memUsage,
		Mood:        mood,
		Message:     message,
	}
}

func displayMood(sm SystemMood) {
	fmt.Printf("⏰ Time: %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("🎯 CPU Usage: %d%%\n", sm.CPUUsage)
	fmt.Printf("💾 Memory Usage: %d%%\n", sm.MemoryUsage)
	fmt.Printf("😎 System Mood: %s\n", sm.Mood)
	fmt.Printf("💬 Message: %s\n", sm.Message)
	fmt.Printf("🖥️  OS: %s | Arch: %s\n", runtime.GOOS, runtime.GOARCH)
}
