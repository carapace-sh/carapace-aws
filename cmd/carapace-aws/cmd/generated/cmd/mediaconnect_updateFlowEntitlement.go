package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateFlowEntitlementCmd = &cobra.Command{
	Use:   "update-flow-entitlement",
	Short: "Updates an entitlement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateFlowEntitlementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_updateFlowEntitlementCmd).Standalone()

		mediaconnect_updateFlowEntitlementCmd.Flags().String("description", "", "A description of the entitlement.")
		mediaconnect_updateFlowEntitlementCmd.Flags().String("encryption", "", "The type of encryption that will be used on the output associated with this entitlement.")
		mediaconnect_updateFlowEntitlementCmd.Flags().String("entitlement-arn", "", "The Amazon Resource Name (ARN) of the entitlement that you want to update.")
		mediaconnect_updateFlowEntitlementCmd.Flags().String("entitlement-status", "", "An indication of whether you want to enable the entitlement to allow access, or disable it to stop streaming content to the subscriber’s flow temporarily.")
		mediaconnect_updateFlowEntitlementCmd.Flags().String("flow-arn", "", "The ARN of the flow that is associated with the entitlement that you want to update.")
		mediaconnect_updateFlowEntitlementCmd.Flags().String("subscribers", "", "The Amazon Web Services account IDs that you want to share your content with.")
		mediaconnect_updateFlowEntitlementCmd.MarkFlagRequired("entitlement-arn")
		mediaconnect_updateFlowEntitlementCmd.MarkFlagRequired("flow-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_updateFlowEntitlementCmd)
}
