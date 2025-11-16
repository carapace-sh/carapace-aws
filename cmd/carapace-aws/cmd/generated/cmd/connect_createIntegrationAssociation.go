package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createIntegrationAssociationCmd = &cobra.Command{
	Use:   "create-integration-association",
	Short: "Creates an Amazon Web Services resource association with an Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createIntegrationAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createIntegrationAssociationCmd).Standalone()

		connect_createIntegrationAssociationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createIntegrationAssociationCmd.Flags().String("integration-arn", "", "The Amazon Resource Name (ARN) of the integration.")
		connect_createIntegrationAssociationCmd.Flags().String("integration-type", "", "The type of information to be ingested.")
		connect_createIntegrationAssociationCmd.Flags().String("source-application-name", "", "The name of the external application.")
		connect_createIntegrationAssociationCmd.Flags().String("source-application-url", "", "The URL for the external application.")
		connect_createIntegrationAssociationCmd.Flags().String("source-type", "", "The type of the data source.")
		connect_createIntegrationAssociationCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createIntegrationAssociationCmd.MarkFlagRequired("instance-id")
		connect_createIntegrationAssociationCmd.MarkFlagRequired("integration-arn")
		connect_createIntegrationAssociationCmd.MarkFlagRequired("integration-type")
	})
	connectCmd.AddCommand(connect_createIntegrationAssociationCmd)
}
