package usecases

import (
	"crud-tasks/domain/models"
	"crud-tasks/domain/models/gateways"
	"fmt"
)

type taskUseCase struct {
	gateway gateways.TaskGateway
}

func FirstTest(task models.Task) {
	fmt.Println("Funciono, la data que llego fue -> ")
	task.Print()
}
