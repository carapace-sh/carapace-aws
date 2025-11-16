package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactFlowMetadataCmd = &cobra.Command{
	Use:   "update-contact-flow-metadata",
	Short: "Updates metadata about specified flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactFlowMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateContactFlowMetadataCmd).Standalone()

		connect_updateContactFlowMetadataCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_updateContactFlowMetadataCmd.Flags().String("contact-flow-state", "", "The state of flow.")
		connect_updateContactFlowMetadataCmd.Flags().String("description", "", "The description of the flow.")
		connect_updateContactFlowMetadataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateContactFlowMetadataCmd.Flags().String("name", "", "The name of the flow.")
		connect_updateContactFlowMetadataCmd.MarkFlagRequired("contact-flow-id")
		connect_updateContactFlowMetadataCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_updateContactFlowMetadataCmd)
}
