package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactFlowModuleContentCmd = &cobra.Command{
	Use:   "update-contact-flow-module-content",
	Short: "Updates specified flow module for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactFlowModuleContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateContactFlowModuleContentCmd).Standalone()

		connect_updateContactFlowModuleContentCmd.Flags().String("contact-flow-module-id", "", "The identifier of the flow module.")
		connect_updateContactFlowModuleContentCmd.Flags().String("content", "", "The JSON string that represents the content of the flow.")
		connect_updateContactFlowModuleContentCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateContactFlowModuleContentCmd.MarkFlagRequired("contact-flow-module-id")
		connect_updateContactFlowModuleContentCmd.MarkFlagRequired("content")
		connect_updateContactFlowModuleContentCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_updateContactFlowModuleContentCmd)
}
