package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Deletes the resource-based policy attached to the CloudTrail event data store, dashboard, or channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_deleteResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_deleteResourcePolicyCmd).Standalone()

		cloudtrail_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudTrail event data store, dashboard, or channel you're deleting the resource-based policy from.")
		cloudtrail_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	cloudtrailCmd.AddCommand(cloudtrail_deleteResourcePolicyCmd)
}
