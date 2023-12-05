package day5

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type mapping struct {
	a, b, c int
}

func (m *mapping) String() string {
	return fmt.Sprintf(`{dest:%d, y:%d, offset:%d}`, m.a, m.b, m.c)
}

func sortLocations(widthByLocation map[int]int) []int {
	sortedLocations := make([]int, len(widthByLocation))
	i := 0
	for location, _ := range widthByLocation {
		sortedLocations[i] = location
		i++
	}
	sort.Ints(sortedLocations)
	return sortedLocations
}

func mapLocations(widthByLocation map[int]int, groupedMappings [][]*mapping) {
	for groupIndex := 0; groupIndex < len(groupedMappings); groupIndex++ {

		sortedLocations := sortLocations(widthByLocation)

		for _, x1 := range sortedLocations {
			x2 := x1 + widthByLocation[x1]

			for _, m := range groupedMappings[groupIndex] {
				if x1 == x2 {
					break
				}

				y1 := m.b
				y2 := m.b + m.c

				overlap := min(x2, y2) - max(x1, y1)
				if overlap < 0 {
					continue
				}

				if y1 > x1 {
					gap := y1 - x2
					widthByLocation[x1] = gap
					x1 = x1 + gap
				}

				offset := m.a - m.b

				widthByLocation[x1+offset] = overlap
				delete(widthByLocation, x1)
				x1 += overlap
			}
		}
	}
}

func getSectionsWithBlocks(scanner *bufio.Scanner) [][]*mapping {
	sections := make([][]*mapping, 0)

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		if idx := strings.Index(text, "map:"); idx == -1 {
			continue
		}

		blocks := make([]*mapping, 0)

		for scanner.Scan() {
			next := scanner.Text()

			if len(next) == 0 {
				break
			}

			split := strings.Split(next, " ")

			a, _ := strconv.Atoi(split[0])
			b, _ := strconv.Atoi(split[1])
			c, _ := strconv.Atoi(split[2])

			blocks = append(blocks, &mapping{a, b, c})
		}

		sort.Slice(blocks, func(i, j int) bool {
			return blocks[i].b < blocks[j].b
		})

		sections = append(sections, blocks)
	}

	return sections
}
