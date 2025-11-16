package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_listConfigurationCheckDefinitionsCmd = &cobra.Command{
	Use:   "list-configuration-check-definitions",
	Short: "Lists all configuration check types supported by AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_listConfigurationCheckDefinitionsCmd).Standalone()

	ssmSap_listConfigurationCheckDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ssmSap_listConfigurationCheckDefinitionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ssmSapCmd.AddCommand(ssmSap_listConfigurationCheckDefinitionsCmd)
}
