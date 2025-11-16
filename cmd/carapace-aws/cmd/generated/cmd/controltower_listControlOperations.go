package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_listControlOperationsCmd = &cobra.Command{
	Use:   "list-control-operations",
	Short: "Provides a list of operations in progress or queued.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_listControlOperationsCmd).Standalone()

	controltower_listControlOperationsCmd.Flags().String("filter", "", "An input filter for the `ListControlOperations` API that lets you select the types of control operations to view.")
	controltower_listControlOperationsCmd.Flags().String("max-results", "", "The maximum number of results to be shown.")
	controltower_listControlOperationsCmd.Flags().String("next-token", "", "A pagination token.")
	controltowerCmd.AddCommand(controltower_listControlOperationsCmd)
}
