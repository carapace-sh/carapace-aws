package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeContactFlowCmd = &cobra.Command{
	Use:   "describe-contact-flow",
	Short: "Describes the specified flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeContactFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeContactFlowCmd).Standalone()

		connect_describeContactFlowCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_describeContactFlowCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeContactFlowCmd.MarkFlagRequired("contact-flow-id")
		connect_describeContactFlowCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_describeContactFlowCmd)
}
