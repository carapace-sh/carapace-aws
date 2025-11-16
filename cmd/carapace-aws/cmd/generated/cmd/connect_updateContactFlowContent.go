package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactFlowContentCmd = &cobra.Command{
	Use:   "update-contact-flow-content",
	Short: "Updates the specified flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactFlowContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateContactFlowContentCmd).Standalone()

		connect_updateContactFlowContentCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_updateContactFlowContentCmd.Flags().String("content", "", "The JSON string that represents the content of the flow.")
		connect_updateContactFlowContentCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateContactFlowContentCmd.MarkFlagRequired("contact-flow-id")
		connect_updateContactFlowContentCmd.MarkFlagRequired("content")
		connect_updateContactFlowContentCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_updateContactFlowContentCmd)
}
