package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteIntegrationAssociationCmd = &cobra.Command{
	Use:   "delete-integration-association",
	Short: "Deletes an Amazon Web Services resource association from an Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteIntegrationAssociationCmd).Standalone()

	connect_deleteIntegrationAssociationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteIntegrationAssociationCmd.Flags().String("integration-association-id", "", "The identifier for the integration association.")
	connect_deleteIntegrationAssociationCmd.MarkFlagRequired("instance-id")
	connect_deleteIntegrationAssociationCmd.MarkFlagRequired("integration-association-id")
	connectCmd.AddCommand(connect_deleteIntegrationAssociationCmd)
}
