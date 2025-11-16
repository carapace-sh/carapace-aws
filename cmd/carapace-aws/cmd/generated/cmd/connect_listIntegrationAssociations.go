package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listIntegrationAssociationsCmd = &cobra.Command{
	Use:   "list-integration-associations",
	Short: "Provides summary information about the Amazon Web Services resource associations for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listIntegrationAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listIntegrationAssociationsCmd).Standalone()

		connect_listIntegrationAssociationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listIntegrationAssociationsCmd.Flags().String("integration-arn", "", "The Amazon Resource Name (ARN) of the integration.")
		connect_listIntegrationAssociationsCmd.Flags().String("integration-type", "", "The integration type.")
		connect_listIntegrationAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listIntegrationAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listIntegrationAssociationsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listIntegrationAssociationsCmd)
}
