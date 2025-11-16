package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_removePermissionCmd = &cobra.Command{
	Use:   "remove-permission",
	Short: "Removes a statement from a topic's access control policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_removePermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_removePermissionCmd).Standalone()

		sns_removePermissionCmd.Flags().String("label", "", "The unique label of the statement you want to remove.")
		sns_removePermissionCmd.Flags().String("topic-arn", "", "The ARN of the topic whose access control policy you wish to modify.")
		sns_removePermissionCmd.MarkFlagRequired("label")
		sns_removePermissionCmd.MarkFlagRequired("topic-arn")
	})
	snsCmd.AddCommand(sns_removePermissionCmd)
}
