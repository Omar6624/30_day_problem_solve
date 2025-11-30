func max (a int , b int){
	if a>b{
		return a
	}
	return b
} 

func rob(house []int , start int, end int) int{
prev:=0
prev_prev := 0

for i:=start ; i<end ;i++{
	curr:=max(prev , prev_prev+house[i])
  	prev_prev = prev
	prev = curr
}
return curr
}

func houseRobber2(house []int) int {
	n:=len(house)

	a:=rob(house,0,n-1)
	b:=rob(house,1,n)

	return max(a,b)
	
}
