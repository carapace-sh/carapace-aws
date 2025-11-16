package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_deleteListenerCmd = &cobra.Command{
	Use:   "delete-listener",
	Short: "Deletes the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_deleteListenerCmd).Standalone()

	elbv2_deleteListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
	elbv2_deleteListenerCmd.MarkFlagRequired("listener-arn")
	elbv2Cmd.AddCommand(elbv2_deleteListenerCmd)
}
