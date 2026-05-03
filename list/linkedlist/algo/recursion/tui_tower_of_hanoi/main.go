// Tower of Hanoi — terminal demonstration in Go
//
// Author : Ikeh Chukwuka Favour (0xull)
// GitHub : https://github.com/0xull
// Run    : go run hanoi.go [--n N] [--delay MS] [--no-anim]
//
// Note: `go run hanoi.go` compiles and executes in one step,
// identical workflow to a scripted language.
//
// Currently, I'm taking the below classes in preparation for the internship (as well
// as just gaining the knowledge):
// 1. Introduction to RISC-V (LFD110)
// 2. Porting Software to RISC-V (LFD114)

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Rod holds a stack of disks. Index 0 is the bottom, last index is the top.
// Disk size is represented as an integer: 1 = smallest, n = largest.
type Rod struct {
	label string
	disks []int
}

func (r *Rod) push(disk int) {
	r.disks = append(r.disks, disk)
}

func (r *Rod) pop() int {
	top := r.disks[len(r.disks)-1]
	r.disks = r.disks[:len(r.disks)-1]
	return top
}

const colWidth = 25

// renderRow returns one horizontal line for a rod at a given display row.
// size == 0 means empty (only the rod glyph is drawn).
func renderRow(size int) string {
	if size == 0 {
		pad := colWidth / 2
		return strings.Repeat(" ", pad) + "│" + strings.Repeat(" ", colWidth-pad-1)
	}
	w := 2*size + 1
	disk := "▐" + strings.Repeat("█", w-2) + "▌"
	pad := (colWidth - w) / 2
	return strings.Repeat(" ", pad) + disk + strings.Repeat(" ", colWidth-pad-w)
}

// draw renders the current puzzle state to the terminal in place (no flicker).
func draw(rods [3]*Rod, n, step int, move string) {
	var sb strings.Builder
	sb.WriteString("\033[H")
	sb.WriteString(fmt.Sprintf("  Tower of Hanoi — %d disks — Step %d / %d\n", n, step, (1<<n)-1))
	sb.WriteString(fmt.Sprintf("  %-36s\n\n", move))

	for row := n - 1; row >= 0; row-- {
		for _, rod := range rods {
			var size int
			if row < len(rod.disks) {
				size = rod.disks[row]
			}
			sb.WriteString(renderRow(size))
		}
		sb.WriteByte('\n')
	}

	base := strings.Repeat("─", colWidth)
	sb.WriteString(base + base + base + "\n")

	for _, rod := range rods {
		pad := colWidth/2 - len(rod.label)/2
		sb.WriteString(strings.Repeat(" ", pad))
		sb.WriteString(rod.label)
		sb.WriteString(strings.Repeat(" ", colWidth-pad-len(rod.label)))
	}
	sb.WriteByte('\n')

	fmt.Print(sb.String())
}

// solve moves n disks from src to dst using aux as the spare rod.
//
// RECURSION: the classic divide-and-conquer approach.
//   - Base case (n == 1): move the single disk directly. No further calls.
//   - Recursive case: move top (n-1) disks out of the way onto aux,
//     move the exposed nth disk to dst, then move the (n-1) disks from
//     aux onto dst. Each recursive call reduces n by 1.
//
// Total moves: T(n) = 2T(n-1) + 1 = 2^n - 1 (provably optimal).
func solve(n int, src, dst, aux *Rod, rods [3]*Rod, step *int, total int, delay time.Duration, anim bool) {
	if n == 1 {
		disk := src.pop()
		dst.push(disk)
		*step++
		desc := fmt.Sprintf("disk %d: %s → %s", disk, src.label, dst.label)
		if anim {
			draw(rods, total, *step, desc)
			time.Sleep(delay)
		} else {
			fmt.Printf("  step %3d: %s\n", *step, desc)
		}
		return
	}

	// Move the top (n-1) disks out of the way, exposing the largest disk.
	solve(n-1, src, aux, dst, rods, step, total, delay, anim)

	// Move the exposed disk to its final position.
	disk := src.pop()
	dst.push(disk)
	*step++
	desc := fmt.Sprintf("disk %d: %s → %s", disk, src.label, dst.label)
	if anim {
		draw(rods, total, *step, desc)
		time.Sleep(delay)
	} else {
		fmt.Printf("  step %3d: %s\n", *step, desc)
	}

	// Stack the (n-1) disks on top of the disk we just placed.
	solve(n-1, aux, dst, src, rods, step, total, delay, anim)
}

func main() {
	n       := flag.Int("n", 5, "number of disks (1-12)")
	delayMs := flag.Int("delay", 280, "milliseconds between frames")
	noAnim  := flag.Bool("no-anim", false, "print step list instead of animation")
	flag.Parse()

	if *n < 1 || *n > 12 {
		fmt.Fprintln(os.Stderr, "error: n must be between 1 and 12")
		os.Exit(1)
	}

	src := &Rod{label: "Source  (A)"}
	aux := &Rod{label: "Spare   (B)"}
	dst := &Rod{label: "Target  (C)"}

	// Load disks onto the source rod, largest at the bottom.
	for i := *n; i >= 1; i-- {
		src.push(i)
	}

	rods  := [3]*Rod{src, aux, dst}
	delay := time.Duration(*delayMs) * time.Millisecond
	anim  := !*noAnim
	total := (1 << *n) - 1

	if anim {
		fmt.Print("\033[2J\033[H")
		draw(rods, *n, 0, fmt.Sprintf("initial state — %d moves total", total))
		time.Sleep(delay)
	} else {
		fmt.Printf("Tower of Hanoi — %d disks — %d moves\n\n", *n, total)
	}

	step := 0
	solve(*n, src, dst, aux, rods, &step, *n, delay, anim)

	if anim {
		draw(rods, *n, step, fmt.Sprintf("solved in %d moves (optimal: 2^%d - 1)", step, *n))
		fmt.Printf("\n  All %d disks moved from A to C.\n\n", *n)
	} else {
		fmt.Printf("\n  Solved: %d steps (optimal: 2^%d - 1 = %d).\n", step, *n, total)
	}
}
