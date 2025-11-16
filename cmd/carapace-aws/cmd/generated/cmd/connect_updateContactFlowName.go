package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateContactFlowNameCmd = &cobra.Command{
	Use:   "update-contact-flow-name",
	Short: "The name of the flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateContactFlowNameCmd).Standalone()

	connect_updateContactFlowNameCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
	connect_updateContactFlowNameCmd.Flags().String("description", "", "The description of the flow.")
	connect_updateContactFlowNameCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateContactFlowNameCmd.Flags().String("name", "", "The name of the flow.")
	connect_updateContactFlowNameCmd.MarkFlagRequired("contact-flow-id")
	connect_updateContactFlowNameCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_updateContactFlowNameCmd)
}
