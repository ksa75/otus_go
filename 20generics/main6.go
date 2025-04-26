package main

// Невозможная комбинация "истинных" интерфейсов и интерфейсов-констрейнтов

type interfaceNormal interface {
	Method()
}

// невозможный интерфейс!
type interfaceConstraint interface {
	int | string | interfaceNormal
}
