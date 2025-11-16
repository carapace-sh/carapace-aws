package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_updateLinkModuleFlowCmd = &cobra.Command{
	Use:   "update-link-module-flow",
	Short: "Updates a link module flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_updateLinkModuleFlowCmd).Standalone()

	rtbfabric_updateLinkModuleFlowCmd.Flags().String("client-token", "", "The unique client token.")
	rtbfabric_updateLinkModuleFlowCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_updateLinkModuleFlowCmd.Flags().String("link-id", "", "The unique identifier of the link.")
	rtbfabric_updateLinkModuleFlowCmd.Flags().String("modules", "", "The configuration of a module.")
	rtbfabric_updateLinkModuleFlowCmd.MarkFlagRequired("client-token")
	rtbfabric_updateLinkModuleFlowCmd.MarkFlagRequired("gateway-id")
	rtbfabric_updateLinkModuleFlowCmd.MarkFlagRequired("link-id")
	rtbfabric_updateLinkModuleFlowCmd.MarkFlagRequired("modules")
	rtbfabricCmd.AddCommand(rtbfabric_updateLinkModuleFlowCmd)
}
