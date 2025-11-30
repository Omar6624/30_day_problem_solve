func max (a int , b int){
	if a>b{
		return a
}
return b

} 
func houseRobber(house []int) int {
	n:= len(house)
	if n==0 {
		return 0
}
//var dp [n+2]int //cant be created at runtime cause it uses a (n) which is  a runtime variable. it //needs constants

	dp:=make([]int , n+2) //need slice
	
	for i:=0 ; i<n ; i++{
		dp[i+2]  = max(dp[i-1], dp[i-2]+house[i])
}

	return dp[len(dp)-1]
}

//complexity time: O(n) space O (n)
