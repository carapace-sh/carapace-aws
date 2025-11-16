package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists all the applications registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listApplicationsCmd).Standalone()

	ssmSap_listApplicationsCmd.Flags().String("filters", "", "The filter of name, value, and operator.")
	ssmSap_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ssmSap_listApplicationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ssmSapCmd.AddCommand(ssmSap_listApplicationsCmd)
}
