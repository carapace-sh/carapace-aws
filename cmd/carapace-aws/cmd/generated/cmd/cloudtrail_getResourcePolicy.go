package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Retrieves the JSON text of the resource-based policy document attached to the CloudTrail event data store, dashboard, or channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_getResourcePolicyCmd).Standalone()

		cloudtrail_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudTrail event data store, dashboard, or channel attached to the resource-based policy.")
		cloudtrail_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	})
	cloudtrailCmd.AddCommand(cloudtrail_getResourcePolicyCmd)
}
