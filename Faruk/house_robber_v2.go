func max (a int , b int){
	if a>b{
		return a
	}
	return b
} 

func houseRobber(house []int) int{
n:= len(house)
	if n==0 {
		return 0
}


prev:=0
prev_prev:=0


for i:=0 ; i<n ; i++{
		curr := max(prev, prev_prev+house[i])
		prev_prev = prev
		prev = curr
}
return curr

}
//complexity time O(n) space O(1)

