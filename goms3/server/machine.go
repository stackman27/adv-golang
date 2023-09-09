package server

import (
	"fmt"
	"io"
	"log"

	"github.com/go-microservice3/fib"
	"github.com/go-microservice3/stack"
	"github.com/go-microservice3/types"
)

type OperatorType string

const (
	PUSH OperatorType = "PUSH" // to ensure type safety we use OperatorType
	POP               = "POP"
	ADD               = "ADD"
	SUB               = "SUB"
	MUL               = "MUL"
	DIV               = "DIV"
	FIB               = "FIB"
)

type MachineServer struct{
	types.UnimplementedMachineServer
}

 
// Execute set of instructions given to the machine
// Machine_ExecuteServer is an interface that contain send, recv and other funcs
func (s *MachineServer) Execute(stream types.Machine_ExecuteServer) error {
	// We are going to put instructions in STACK (LIFO) and execute it
	var stack stack.Stack

	for {
	instruction, err := stream.Recv()
	if err == io.EOF {
		log.Println("EOF")
		return nil
	}

	if err != nil {
		return err
	}

	operand := instruction.GetOperand()
	operator := instruction.GetOperator()
	op_type := OperatorType(operator)

	fmt.Printf("Operand: %v, Operator: %v\n", operand, operator)

	switch op_type {
	case PUSH:
		stack.ExportPush(float32(operand))
	case POP:
		stack.ExportPop()

	case ADD, SUB, MUL, DIV:
		// get the last 2 value from the stack
		val1, canPop := stack.ExportPop()
		val2, canPop := stack.ExportPop()

		if !canPop {
			return fmt.Errorf("unable to pop values from stack, please check stack length")
		}

		var res float32
		switch op_type {
		case ADD:
			res = val1 + val2

		case SUB:
			res = val1 - val2

		case MUL:
			res = val1 * val2

		case DIV:
			res = val1 / val2

		default:
			return fmt.Errorf("unidentified operator")
		}

		stack.ExportPush(res)
		err := stream.Send(&types.Result{
			Output: res,
		})

		if err != nil {
			return err
		}

	case FIB:
		val1, canPop := stack.ExportPop()
		if !canPop {
			return fmt.Errorf("invalid set of instruction to pop")
		}

		// get the val1 digit fibonacci sequence
		nthFib := fib.GetnthFibSeq(int(val1))
		stack.ExportPush(float32(nthFib))

		err := stream.Send(&types.Result{
			Output: float32(nthFib),
		})

		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Operation is not implemented yet: %s", operator)
	}
}
}
