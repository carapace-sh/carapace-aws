package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createContactFlowVersionCmd = &cobra.Command{
	Use:   "create-contact-flow-version",
	Short: "Publishes a new version of the flow provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createContactFlowVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createContactFlowVersionCmd).Standalone()

		connect_createContactFlowVersionCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_createContactFlowVersionCmd.Flags().String("contact-flow-version", "", "The identifier of the flow version.")
		connect_createContactFlowVersionCmd.Flags().String("description", "", "The description of the flow version.")
		connect_createContactFlowVersionCmd.Flags().String("flow-content-sha256", "", "Indicates the checksum value of the flow content.")
		connect_createContactFlowVersionCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createContactFlowVersionCmd.Flags().String("last-modified-region", "", "The Amazon Web Services Region where this resource was last modified.")
		connect_createContactFlowVersionCmd.Flags().String("last-modified-time", "", "The Amazon Web Services Region where this resource was last modified.")
		connect_createContactFlowVersionCmd.MarkFlagRequired("contact-flow-id")
		connect_createContactFlowVersionCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_createContactFlowVersionCmd)
}
