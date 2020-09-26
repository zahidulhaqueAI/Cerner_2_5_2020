 // cerner_2^5_2020
 
 func binarySearch(ranked []int32, low int32, high int32, score int32) int32 {
     mid := low + (high - low)/2
     if ranked[mid] == score {   return mid
     } else if score > ranked[mid] && score < ranked[mid-1] {
         return mid
     } else if score > ranked[mid] && score < ranked[mid-1] {
         return mid-1
     }
     if score > ranked[mid]{     return binarySearch(ranked,low, mid-1,score)
     } else {    return binarySearch(ranked,mid+1,high,score)    }
}
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    n := len(player)
    m := len(ranked)
    result := make([]int32,n)
    rank := make([]int32,m)
    rank[0] = 1
    for i := 1; i < m; i++{
        if ranked[i] == ranked[i-1]{   rank[i] = rank[i-1]
        } else {    rank[i] = rank[i-1]+1
        }
    }
    for i := range player{   score := player[i]
        if score >= ranked[0] { result[i] = 1
        } else if score < ranked[m-1]{
            result[i] = rank[m-1]+1
        }else {   index := binarySearch(ranked, 0, int32(m-1), score)
            result[i] = rank[index]
        }
    }
    return result
}
