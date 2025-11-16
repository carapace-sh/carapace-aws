package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_acceptReservedNodeExchangeCmd = &cobra.Command{
	Use:   "accept-reserved-node-exchange",
	Short: "Exchanges a DC1 Reserved Node for a DC2 Reserved Node with no changes to the configuration (term, payment type, or number of nodes) and no additional costs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_acceptReservedNodeExchangeCmd).Standalone()

	redshift_acceptReservedNodeExchangeCmd.Flags().String("reserved-node-id", "", "A string representing the node identifier of the DC1 Reserved Node to be exchanged.")
	redshift_acceptReservedNodeExchangeCmd.Flags().String("target-reserved-node-offering-id", "", "The unique identifier of the DC2 Reserved Node offering to be used for the exchange.")
	redshift_acceptReservedNodeExchangeCmd.MarkFlagRequired("reserved-node-id")
	redshift_acceptReservedNodeExchangeCmd.MarkFlagRequired("target-reserved-node-offering-id")
	redshiftCmd.AddCommand(redshift_acceptReservedNodeExchangeCmd)
}
