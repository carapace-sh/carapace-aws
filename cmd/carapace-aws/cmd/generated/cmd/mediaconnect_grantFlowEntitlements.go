package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_grantFlowEntitlementsCmd = &cobra.Command{
	Use:   "grant-flow-entitlements",
	Short: "Grants entitlements to an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_grantFlowEntitlementsCmd).Standalone()

	mediaconnect_grantFlowEntitlementsCmd.Flags().String("entitlements", "", "The list of entitlements that you want to grant.")
	mediaconnect_grantFlowEntitlementsCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to grant entitlements on.")
	mediaconnect_grantFlowEntitlementsCmd.MarkFlagRequired("entitlements")
	mediaconnect_grantFlowEntitlementsCmd.MarkFlagRequired("flow-arn")
	mediaconnectCmd.AddCommand(mediaconnect_grantFlowEntitlementsCmd)
}
