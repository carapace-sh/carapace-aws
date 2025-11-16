package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_getReservedNodeExchangeConfigurationOptionsCmd = &cobra.Command{
	Use:   "get-reserved-node-exchange-configuration-options",
	Short: "Gets the configuration options for the reserved-node exchange.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_getReservedNodeExchangeConfigurationOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_getReservedNodeExchangeConfigurationOptionsCmd).Standalone()

		redshift_getReservedNodeExchangeConfigurationOptionsCmd.Flags().String("action-type", "", "The action type of the reserved-node configuration.")
		redshift_getReservedNodeExchangeConfigurationOptionsCmd.Flags().String("cluster-identifier", "", "The identifier for the cluster that is the source for a reserved-node exchange.")
		redshift_getReservedNodeExchangeConfigurationOptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `GetReservedNodeExchangeConfigurationOptions` request.")
		redshift_getReservedNodeExchangeConfigurationOptionsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
		redshift_getReservedNodeExchangeConfigurationOptionsCmd.Flags().String("snapshot-identifier", "", "The identifier for the snapshot that is the source for the reserved-node exchange.")
		redshift_getReservedNodeExchangeConfigurationOptionsCmd.MarkFlagRequired("action-type")
	})
	redshiftCmd.AddCommand(redshift_getReservedNodeExchangeConfigurationOptionsCmd)
}
