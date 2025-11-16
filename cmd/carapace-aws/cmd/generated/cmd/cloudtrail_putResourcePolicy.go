package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based permission policy to a CloudTrail event data store, dashboard, or channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_putResourcePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_putResourcePolicyCmd).Standalone()

		cloudtrail_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudTrail event data store, dashboard, or channel attached to the resource-based policy.")
		cloudtrail_putResourcePolicyCmd.Flags().String("resource-policy", "", "A JSON-formatted string for an Amazon Web Services resource-based policy.")
		cloudtrail_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
		cloudtrail_putResourcePolicyCmd.MarkFlagRequired("resource-policy")
	})
	cloudtrailCmd.AddCommand(cloudtrail_putResourcePolicyCmd)
}
