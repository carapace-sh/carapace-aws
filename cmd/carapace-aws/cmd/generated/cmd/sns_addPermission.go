package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_addPermissionCmd = &cobra.Command{
	Use:   "add-permission",
	Short: "Adds a statement to a topic's access control policy, granting access for the specified Amazon Web Services accounts to the specified actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_addPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_addPermissionCmd).Standalone()

		sns_addPermissionCmd.Flags().String("action-name", "", "The action you want to allow for the specified principal(s).")
		sns_addPermissionCmd.Flags().String("awsaccount-id", "", "The Amazon Web Services account IDs of the users (principals) who will be given access to the specified actions.")
		sns_addPermissionCmd.Flags().String("label", "", "A unique identifier for the new policy statement.")
		sns_addPermissionCmd.Flags().String("topic-arn", "", "The ARN of the topic whose access control policy you wish to modify.")
		sns_addPermissionCmd.MarkFlagRequired("action-name")
		sns_addPermissionCmd.MarkFlagRequired("awsaccount-id")
		sns_addPermissionCmd.MarkFlagRequired("label")
		sns_addPermissionCmd.MarkFlagRequired("topic-arn")
	})
	snsCmd.AddCommand(sns_addPermissionCmd)
}
