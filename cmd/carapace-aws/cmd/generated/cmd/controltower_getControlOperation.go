package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_getControlOperationCmd = &cobra.Command{
	Use:   "get-control-operation",
	Short: "Returns the status of a particular `EnableControl` or `DisableControl` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_getControlOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_getControlOperationCmd).Standalone()

		controltower_getControlOperationCmd.Flags().String("operation-identifier", "", "The ID of the asynchronous operation, which is used to track status.")
		controltower_getControlOperationCmd.MarkFlagRequired("operation-identifier")
	})
	controltowerCmd.AddCommand(controltower_getControlOperationCmd)
}
