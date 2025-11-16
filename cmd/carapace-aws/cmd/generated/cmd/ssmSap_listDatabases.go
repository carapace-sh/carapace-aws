package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listDatabasesCmd = &cobra.Command{
	Use:   "list-databases",
	Short: "Lists the SAP HANA databases of an application registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listDatabasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_listDatabasesCmd).Standalone()

		ssmSap_listDatabasesCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_listDatabasesCmd.Flags().String("component-id", "", "The ID of the component.")
		ssmSap_listDatabasesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ssmSap_listDatabasesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	ssmSapCmd.AddCommand(ssmSap_listDatabasesCmd)
}
