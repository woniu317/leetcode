package main

func main() {
	numCourses := 2
	prerequisites := [][]int{{}, {1, 0}}
	println(canFinish(numCourses, prerequisites))
}
func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]int, numCourses)
	visitedCode := 1
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if dfs(&prerequisites, &visited, i, visitedCode) == false {
				return false
			}
		}
	}
	return true
}

func dfs(prerequisites *[][]int, visited *[]int, visitedNode, visitedCode int) bool {
	//有环false
	for i := 0; i < len((*prerequisites)[visitedNode]); i++ {
		if (*visited)[(*prerequisites)[visitedNode][i]] == 0 {
			(*visited)[(*prerequisites)[visitedNode][i]] = visitedCode
			return dfs(prerequisites, visited, (*prerequisites)[visitedNode][i], visitedCode)
		} else if (*visited)[(*prerequisites)[visitedNode][i]] == visitedCode {
			return false
		}
	}
	return true
}
