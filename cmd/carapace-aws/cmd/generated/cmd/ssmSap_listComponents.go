package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listComponentsCmd = &cobra.Command{
	Use:   "list-components",
	Short: "Lists all the components registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listComponentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_listComponentsCmd).Standalone()

		ssmSap_listComponentsCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_listComponentsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ssmSap_listComponentsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	ssmSapCmd.AddCommand(ssmSap_listComponentsCmd)
}
