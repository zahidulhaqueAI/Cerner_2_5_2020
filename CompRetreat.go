// cerner_2^5_2020

import "math"
func solve(supervisors []int32, groups [][]int32) int32 {
    empMap := make(map[int32]int32)
    for i := int32(0); i < int32(len(supervisors)); i++ {
        empMap[i+2] = supervisors[i]
    }
    Totcost := int32(0)
    m := int32(len(supervisors) + 1)
    var allVisited int32
    for i := 0; i < len(groups); i++ { cnt := int32(0)
        isVisited := make(map[int32]int)
        for j := m; j >= 1; j-- {
            if allVisited == m { allVisited = 0
                break
            }
            cnt++
            isVisited[j] = 1 // for m
            allVisited += 1 // for m
            ind := empMap[j]
            allVisited += 1
            for k := int32(0); k < groups[i][1]-1; k++ {
                if isVisited[ind] != 1 { isVisited[ind] = 1
                    ind = empMap[ind]
                    allVisited += 1
                }
            }      
        }
        Totcost += cnt * groups[i][0]
    }
    return (Totcost % int32(math.Pow(10,9)+7))
    }
