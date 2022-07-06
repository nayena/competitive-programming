package main

type item struct {
	Letter rune
	Count  int
}

func removeDuplicates(s string, k int) string {
	var stack []item

	for _, letter := range s {
		if len(stack) == 0 {
			stack = append(stack, item{letter, 1})
		} else {
			if top(stack).Letter == letter {
				stack = incrementTop(stack)
			} else {
				stack = append(stack, item{letter, 1})
			}

			if top(stack).Count == k {
				stack = pop(stack)
			}
		}
	}

	result := ""
	for _, item := range stack {
		for i := 0; i < item.Count; i++ {
			result += string(item.Letter)
		}
	}
	return result
}

func incrementTop(stack []item) []item {
	stack[len(stack)-1].Count++
	return stack
}

func top(stack []item) item {
	return stack[len(stack)-1]
}

func pop(stack []item) []item {
	return stack[:len(stack)-1]
}
