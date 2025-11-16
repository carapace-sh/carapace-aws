package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteContactFlowCmd = &cobra.Command{
	Use:   "delete-contact-flow",
	Short: "Deletes a flow for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteContactFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteContactFlowCmd).Standalone()

		connect_deleteContactFlowCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_deleteContactFlowCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_deleteContactFlowCmd.MarkFlagRequired("contact-flow-id")
		connect_deleteContactFlowCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_deleteContactFlowCmd)
}
