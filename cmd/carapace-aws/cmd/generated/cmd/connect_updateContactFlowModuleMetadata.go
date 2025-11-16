package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactFlowModuleMetadataCmd = &cobra.Command{
	Use:   "update-contact-flow-module-metadata",
	Short: "Updates metadata about specified flow module.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactFlowModuleMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateContactFlowModuleMetadataCmd).Standalone()

		connect_updateContactFlowModuleMetadataCmd.Flags().String("contact-flow-module-id", "", "The identifier of the flow module.")
		connect_updateContactFlowModuleMetadataCmd.Flags().String("description", "", "The description of the flow module.")
		connect_updateContactFlowModuleMetadataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateContactFlowModuleMetadataCmd.Flags().String("name", "", "The name of the flow module.")
		connect_updateContactFlowModuleMetadataCmd.Flags().String("state", "", "The state of flow module.")
		connect_updateContactFlowModuleMetadataCmd.MarkFlagRequired("contact-flow-module-id")
		connect_updateContactFlowModuleMetadataCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_updateContactFlowModuleMetadataCmd)
}
