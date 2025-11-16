package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_getAssignmentCmd = &cobra.Command{
	Use:   "get-assignment",
	Short: "The `GetAssignment` operation retrieves the details of the specified Assignment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_getAssignmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_getAssignmentCmd).Standalone()

		mturk_getAssignmentCmd.Flags().String("assignment-id", "", "The ID of the Assignment to be retrieved.")
		mturk_getAssignmentCmd.MarkFlagRequired("assignment-id")
	})
	mturkCmd.AddCommand(mturk_getAssignmentCmd)
}
