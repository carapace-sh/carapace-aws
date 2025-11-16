package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listConnectorsV2Cmd = &cobra.Command{
	Use:   "list-connectors-v2",
	Short: "Grants permission to retrieve a list of connectorsV2 and their metadata for the calling account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listConnectorsV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_listConnectorsV2Cmd).Standalone()

		securityhub_listConnectorsV2Cmd.Flags().String("connector-status", "", "The status for the connectorV2.")
		securityhub_listConnectorsV2Cmd.Flags().String("max-results", "", "The maximum number of results to be returned.")
		securityhub_listConnectorsV2Cmd.Flags().String("next-token", "", "The pagination token per the Amazon Web Services Pagination standard")
		securityhub_listConnectorsV2Cmd.Flags().String("provider-name", "", "The name of the third-party provider.")
	})
	securityhubCmd.AddCommand(securityhub_listConnectorsV2Cmd)
}
