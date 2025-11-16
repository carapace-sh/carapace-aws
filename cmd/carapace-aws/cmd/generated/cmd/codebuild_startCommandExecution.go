package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_startCommandExecutionCmd = &cobra.Command{
	Use:   "start-command-execution",
	Short: "Starts a command execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_startCommandExecutionCmd).Standalone()

	codebuild_startCommandExecutionCmd.Flags().String("command", "", "The command that needs to be executed.")
	codebuild_startCommandExecutionCmd.Flags().String("sandbox-id", "", "A `sandboxId` or `sandboxArn`.")
	codebuild_startCommandExecutionCmd.Flags().String("type", "", "The command type.")
	codebuild_startCommandExecutionCmd.MarkFlagRequired("command")
	codebuild_startCommandExecutionCmd.MarkFlagRequired("sandbox-id")
	codebuildCmd.AddCommand(codebuild_startCommandExecutionCmd)
}
