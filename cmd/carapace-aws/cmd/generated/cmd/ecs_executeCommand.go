package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_executeCommandCmd = &cobra.Command{
	Use:   "execute-command",
	Short: "Runs a command remotely on a container within a task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_executeCommandCmd).Standalone()

	ecs_executeCommandCmd.Flags().String("cluster", "", "The Amazon Resource Name (ARN) or short name of the cluster the task is running in.")
	ecs_executeCommandCmd.Flags().String("command", "", "The command to run on the container.")
	ecs_executeCommandCmd.Flags().String("container", "", "The name of the container to execute the command on.")
	ecs_executeCommandCmd.Flags().Bool("interactive", false, "Use this flag to run your command in interactive mode.")
	ecs_executeCommandCmd.Flags().Bool("no-interactive", false, "Use this flag to run your command in interactive mode.")
	ecs_executeCommandCmd.Flags().String("task", "", "The Amazon Resource Name (ARN) or ID of the task the container is part of.")
	ecs_executeCommandCmd.MarkFlagRequired("command")
	ecs_executeCommandCmd.MarkFlagRequired("interactive")
	ecs_executeCommandCmd.Flag("no-interactive").Hidden = true
	ecs_executeCommandCmd.MarkFlagRequired("no-interactive")
	ecs_executeCommandCmd.MarkFlagRequired("task")
	ecsCmd.AddCommand(ecs_executeCommandCmd)
}
