package main

import "fmt"

type Command interface{
	Execute()
}

type ConsoleOutput struct{
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

type CommandQueue struct {
	queue []Command
}
  
//executes command at the burst of 3
func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)
	if len(p.queue) == 3 {
	  for _, command := range p.queue {
		command.Execute()
	  }
	  p.queue = make([]Command, 3)
	}
}

func CreateCommand(str string) Command {
	return &ConsoleOutput{message:str}
}

func main() {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First message"))
	queue.AddCommand(CreateCommand("Second message"))
	queue.AddCommand(CreateCommand("Third message"))
	queue.AddCommand(CreateCommand("Fourth message"))
	queue.AddCommand(CreateCommand("Fifth message"))
}