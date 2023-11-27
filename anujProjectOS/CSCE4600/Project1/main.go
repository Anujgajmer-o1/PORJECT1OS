package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Process struct represents a process
type Process struct {
	ID            int
	BurstDuration int
	ArrivalTime   int
	Priority      int
}

// ByBurstDuration implements sort.Interface for []Process based on BurstDuration
type ByBurstDuration []Process

func (a ByBurstDuration) Len() int           { return len(a) }
func (a ByBurstDuration) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBurstDuration) Less(i, j int) bool { return a[i].BurstDuration < a[j].BurstDuration }

func main() {
	// Parse input file
	processes := readProcessesFromFile(os.Args[1])

	// Sort processes by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	// Implement SJF scheduling
	sjfScheduling(processes)

	// Output results
	// You can calculate and print average turnaround time, average waiting time, and average throughput here
}

func sjfScheduling(processes []Process) {
	// Implement SJF (preemptive) scheduling logic here
	// Keep track of time, execute processes, update waiting times, etc.
}

func readProcessesFromFile(filename string) []Process {
	// Implement logic to read processes from the file and return a slice of Process
	// Each line in the file contains a record with comma-separated fields
	// ProcessID, BurstDuration, ArrivalTime, Priority
	// Hint: You can use the strconv.Atoi function to convert string to int
	// You may want to create a Process struct and return a slice of Process
}
