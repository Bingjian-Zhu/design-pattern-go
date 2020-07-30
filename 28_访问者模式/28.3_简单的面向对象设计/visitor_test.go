package visitor

func ExampleVisitor() {
	persons := make([]Person, 0)

	man1 := &Man{}
	man1.Action = "成功"
	persons = append(persons, man1)

	woman1 := &Woman{}
	woman1.Action = "成功"
	persons = append(persons, woman1)

	man2 := &Man{}
	man2.Action = "失败"
	persons = append(persons, man2)

	woman2 := &Woman{}
	woman2.Action = "失败"
	persons = append(persons, woman2)

	man3 := &Man{}
	man3.Action = "恋爱"
	persons = append(persons, man3)

	woman3 := &Woman{}
	woman3.Action = "恋爱"
	persons = append(persons, woman3)

	for _, item := range persons {
		item.GetConclusion()
	}

	// OutPut:
	// Man成功时，背后多半有一个伟大的女人。
	// Woman成功时，背后大多有一个不成功的男人。
	// Man失败时，闷头喝酒，谁也不用劝。
	// Woman失败时，眼泪汪汪，谁也劝不了。
	// Man恋爱时，凡事不懂也要装懂。
	// Woman恋爱时，遇事懂也装作不懂。
}
