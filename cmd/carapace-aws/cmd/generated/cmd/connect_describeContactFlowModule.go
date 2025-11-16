package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeContactFlowModuleCmd = &cobra.Command{
	Use:   "describe-contact-flow-module",
	Short: "Describes the specified flow module.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeContactFlowModuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeContactFlowModuleCmd).Standalone()

		connect_describeContactFlowModuleCmd.Flags().String("contact-flow-module-id", "", "The identifier of the flow module.")
		connect_describeContactFlowModuleCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeContactFlowModuleCmd.MarkFlagRequired("contact-flow-module-id")
		connect_describeContactFlowModuleCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_describeContactFlowModuleCmd)
}
