package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_revokeFlowEntitlementCmd = &cobra.Command{
	Use:   "revoke-flow-entitlement",
	Short: "Revokes an entitlement from a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_revokeFlowEntitlementCmd).Standalone()

	mediaconnect_revokeFlowEntitlementCmd.Flags().String("entitlement-arn", "", "The Amazon Resource Name (ARN) of the entitlement that you want to revoke.")
	mediaconnect_revokeFlowEntitlementCmd.Flags().String("flow-arn", "", "The flow that you want to revoke an entitlement from.")
	mediaconnect_revokeFlowEntitlementCmd.MarkFlagRequired("entitlement-arn")
	mediaconnect_revokeFlowEntitlementCmd.MarkFlagRequired("flow-arn")
	mediaconnectCmd.AddCommand(mediaconnect_revokeFlowEntitlementCmd)
}
