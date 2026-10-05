func largestRectangleArea(heights [] int ) int{
	stack :=[]int{}
	maxArea :=0
	n:=len(heights)

	for i :=0;i<=n;i++ {
		for len(stack)>0 && (i==n) || heights[stack[len(stack)-1]]>=heights[i]{
			h:=heights[stack[len(stack)-1]]
			stack =stack[:len(stack)-1]

			left :=-1
			if len(stack)>0{
				left = stack[len(stack)-1]
			}
			width :=i-left-1
			area :=h*width
			if area >maxArea{
				maxArea=area
			}
		}
		stack =append(stack,i)

	}
	return maxArea
}