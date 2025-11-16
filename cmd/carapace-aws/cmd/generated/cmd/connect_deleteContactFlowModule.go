package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteContactFlowModuleCmd = &cobra.Command{
	Use:   "delete-contact-flow-module",
	Short: "Deletes the specified flow module.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteContactFlowModuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteContactFlowModuleCmd).Standalone()

		connect_deleteContactFlowModuleCmd.Flags().String("contact-flow-module-id", "", "The identifier of the flow module.")
		connect_deleteContactFlowModuleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteContactFlowModuleCmd.MarkFlagRequired("contact-flow-module-id")
		connect_deleteContactFlowModuleCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_deleteContactFlowModuleCmd)
}
