package generate_parentheses

var ret = []string{}

func GenerateParentheses(n int) []string {
	ret = []string{}
	dfs(n, 0, 0, "")
	return ret
}

func dfs(n, leftCount, rightCount int, val string) {
	if leftCount == n && rightCount == n {
		ret = append(ret, val)
	}
	if leftCount < n {
		dfs(n, leftCount+1, rightCount, val+"(")
	}
	if leftCount > rightCount && rightCount < n {
		dfs(n, leftCount, rightCount+1, val+")")
	}
}
