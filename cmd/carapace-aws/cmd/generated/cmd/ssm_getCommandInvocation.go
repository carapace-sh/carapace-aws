package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getCommandInvocationCmd = &cobra.Command{
	Use:   "get-command-invocation",
	Short: "Returns detailed information about command execution for an invocation or plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getCommandInvocationCmd).Standalone()

	ssm_getCommandInvocationCmd.Flags().String("command-id", "", "(Required) The parent command ID of the invocation plugin.")
	ssm_getCommandInvocationCmd.Flags().String("instance-id", "", "(Required) The ID of the managed node targeted by the command.")
	ssm_getCommandInvocationCmd.Flags().String("plugin-name", "", "The name of the step for which you want detailed results.")
	ssm_getCommandInvocationCmd.MarkFlagRequired("command-id")
	ssm_getCommandInvocationCmd.MarkFlagRequired("instance-id")
	ssmCmd.AddCommand(ssm_getCommandInvocationCmd)
}
