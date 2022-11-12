package main

import (
	"fmt"
	"os"

	"time"

	"github.com/spf13/cobra"
	Seal "github.com/sxhk0/seal"
	Shared "github.com/sxhk0/seal/shared"
)

func main() {
	work_path := Shared.WorkingDirectory()

	environment := map[string]string{}

	loadEnvironmentFile := true
	restart := true

	seal := &cobra.Command{
		Use:   "seal",
		Short: "Seal - A tool to run commands with monitoring features.",
	}

	sealRun := &cobra.Command{
		Use:                   "run [flags] -- [command] [...args]",
		Short:                 "Runs a program.",
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		Run: func(command *cobra.Command, args []string) {
			if loadEnvironmentFile {
				environment = Shared.ReadEnvironmentFile(work_path)
			}

			executable := args[0]
			arguments := args[1:]

			println(executable, arguments)

			process := Seal.Process{
				Executable:   executable,
				Arguments:    arguments,
				Environment:  environment,
				Restart:      restart,
				RestartLimit: 5,
			}

			process.Start()

			for {
				// do not close
				time.Sleep(5 * time.Second)
			}
		},
	}

	sealRun.Flags().BoolP("auto-restart", "r", restart, "restarts when process goes down")
	sealRun.Flags().BoolP("load-env", "e", loadEnvironmentFile, "loads .env file")

	seal.AddCommand(sealRun)

	if error := seal.Execute(); error != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", error)
		os.Exit(1)
	}
}
