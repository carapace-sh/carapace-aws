package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_getConfigurationCheckOperationCmd = &cobra.Command{
	Use:   "get-configuration-check-operation",
	Short: "Gets the details of a configuration check operation by specifying the operation ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_getConfigurationCheckOperationCmd).Standalone()

	ssmSap_getConfigurationCheckOperationCmd.Flags().String("operation-id", "", "The ID of the configuration check operation.")
	ssmSap_getConfigurationCheckOperationCmd.MarkFlagRequired("operation-id")
	ssmSapCmd.AddCommand(ssmSap_getConfigurationCheckOperationCmd)
}
