package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_startConfigurationChecksCmd = &cobra.Command{
	Use:   "start-configuration-checks",
	Short: "Initiates configuration check operations against a specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_startConfigurationChecksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_startConfigurationChecksCmd).Standalone()

		ssmSap_startConfigurationChecksCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_startConfigurationChecksCmd.Flags().String("configuration-check-ids", "", "The list of configuration checks to perform.")
		ssmSap_startConfigurationChecksCmd.MarkFlagRequired("application-id")
	})
	ssmSapCmd.AddCommand(ssmSap_startConfigurationChecksCmd)
}
