package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    out := map[string]int{}
    for points, letters := range in {
        for _, ltr := range letters {
            out[strings.ToLower(ltr)] = points
        }
    }

    return out
}
