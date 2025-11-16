package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteContactFlowVersionCmd = &cobra.Command{
	Use:   "delete-contact-flow-version",
	Short: "Deletes the particular version specified in flow version identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteContactFlowVersionCmd).Standalone()

	connect_deleteContactFlowVersionCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
	connect_deleteContactFlowVersionCmd.Flags().String("contact-flow-version", "", "The identifier of the flow version.")
	connect_deleteContactFlowVersionCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteContactFlowVersionCmd.MarkFlagRequired("contact-flow-id")
	connect_deleteContactFlowVersionCmd.MarkFlagRequired("contact-flow-version")
	connect_deleteContactFlowVersionCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_deleteContactFlowVersionCmd)
}
