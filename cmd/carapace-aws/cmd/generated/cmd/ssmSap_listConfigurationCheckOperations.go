package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listConfigurationCheckOperationsCmd = &cobra.Command{
	Use:   "list-configuration-check-operations",
	Short: "Lists the configuration check operations performed by AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listConfigurationCheckOperationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_listConfigurationCheckOperationsCmd).Standalone()

		ssmSap_listConfigurationCheckOperationsCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_listConfigurationCheckOperationsCmd.Flags().String("filters", "", "The filters of an operation.")
		ssmSap_listConfigurationCheckOperationsCmd.Flags().String("list-mode", "", "The mode for listing configuration check operations.")
		ssmSap_listConfigurationCheckOperationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ssmSap_listConfigurationCheckOperationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ssmSap_listConfigurationCheckOperationsCmd.MarkFlagRequired("application-id")
	})
	ssmSapCmd.AddCommand(ssmSap_listConfigurationCheckOperationsCmd)
}
