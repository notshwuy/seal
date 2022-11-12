package seal

import (
	"fmt"
	"os"
	"os/exec"
)

type Process struct {
	Executable   string
	Arguments    []string
	Environment  map[string]string
	Restart      bool
	RestartLimit int
}

func (process *Process) Start() {
	command := exec.Command(process.Executable, process.Arguments...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	environment_variables := []string{}

	for key, value := range process.Environment {
		variable := fmt.Sprintf("%s=%s", key, value)
		environment_variables = append(environment_variables, variable)
	}

	command.Env = append(os.Environ(), environment_variables...)

	error := command.Start()
	must(error)

	go func() {
		state, error := command.Process.Wait()
		must(error)

		if state.Exited() {
			fmt.Printf("Process exited with %d\n", state.ExitCode())

			if process.Restart && process.RestartLimit != 0 {
				fmt.Printf("Restarting...\n")

				process.RestartLimit--
				process.Start()
			} else {
				os.Exit(state.ExitCode())
			}
		}
	}()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
