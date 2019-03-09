package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Arafatk/glot"
)

func random() int64 {
	rand.Seed(time.Now().Unix())
	rand.Seed(rand.Int63n(100000000))
	return rand.Int63n(2) + 1
}

func main() {
	var (
		minimum float64 = 10 * NUM_OF_CELLS
		maximum float64 = 0
	)
	dimensions := 2
	var isRand bool = false
	var isA bool = false
	plot, _ := glot.NewPlot(dimensions, false, false)
	for k := 1; k <= 10; k += 9 {
		cells := make([]*cell, 1001)
		for i := range cells {
			cells[i] = &cell{
				-1,
				false,
				false,
				NUM_OF_CELLS / 100 * 75,
				NUM_OF_CELLS / 10,
				2,
				5,
				false,
				1,
				false,
			}
		}
		if isA {
			for i := 900; i <= 990; i++ {
				cells[i].delta = 0
			}
		}
		cells[500/k] = Kill()
		one := make([]float64, 0)
		two := make([]float64, 0)
		var period int64 = NUM_OF_CELLS
		var timer int64
		var iterator = 1
		for ; timer <= 10*NUM_OF_CELLS; timer++ {
			iterator = 1
			if timer%period == 0 {
				cells[0].GivePotential(timer)
				cells[NUM_OF_CELLS].GivePotential(timer)
				if isRand {
					period = random()
				}
			}
			for i := 0; i <= int(NUM_OF_CELLS) && i >= 0; i += iterator {
				iterator = 1
				if cells[i].dead {
					if iterator < 0 {
						break
					} else {
						i = int(NUM_OF_CELLS)
						iterator = -1
					}
				}
				if timer == cells[i].compressedTime+cells[i].uptime+cells[i].refractorTime {
					cells[i].Decompress()
					continue
				} else if cells[i].uptime >= 0 && timer == cells[i].compressedTime+cells[i].uptime {
					cells[i].SetRefractor()
					continue
				}
				if cells[i].uptime >= 0 && timer == cells[i].delta+cells[i].uptime && cells[i].potential {
					if !cells[i].refractor {
						cells[i+iterator].GivePotential(timer)
					}
					cells[i].potential = false
				}
			}
			var s float64
			for _, c := range cells {
				if c.compressed && !c.refractor {
					s += float64(c.compressedSize)
				} else {
					s += float64(c.decompressedSize)
				}
			}
			s = (s * s / (4 * math.Pi)) / 1000000
			log.Println(timer, s)
			one = append(one, float64(timer)/float64(period)-2)
			two = append(two, s)
			s = 0
		}
		locMinimum := min(two[5*NUM_OF_CELLS : 6*NUM_OF_CELLS])
		locMaximum := max(two[5*NUM_OF_CELLS : 6*NUM_OF_CELLS])
		log.Println(locMaximum / locMinimum)
		if locMaximum > maximum {
			maximum = locMaximum
		}
		if locMinimum < minimum {
			minimum = locMinimum
		}
		pointGroupName := ""
		if k == 1 {
			pointGroupName = "Здоровое псевдосердце"
		} else {
			pointGroupName = fmt.Sprint("Повреждена клетка №", 50)
		}
		style := "lines"
		points := [][]float64{one, two}
		_ = plot.AddPointGroup(pointGroupName, style, points)
		if isA {
			break
		} else {
			isA = true
			k--
		}
	}

	_ = plot.SetTitle("Одномерная модель Винера-Розенблюта")
	_ = plot.SetXLabel("Время, с")
	_ = plot.SetYLabel("Объём")
	_ = plot.SetXrange(0, 4)
	d := maximum - minimum
	_ = plot.SetYrange(0, int(maximum+d))
	// Optional: Setting axis ranges
	f, _ := os.Open("src/last.txt")
	var str string
	_, _ = fmt.Fscan(f, &str)
	err := plot.SavePlot("output/" + str + ".png")
	log.Println(err)
	num, err := strconv.Atoi(strings.Trim(str, " \n\t"))
	log.Println(err)
	num++
	log.Println(num)
	_ = os.Remove("src/last.txt")
	f, err = os.Create("src/last.txt")
	_, _ = fmt.Fprint(f, num)
}
