package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeSourceNetworksCmd = &cobra.Command{
	Use:   "describe-source-networks",
	Short: "Lists all Source Networks or multiple Source Networks filtered by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeSourceNetworksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_describeSourceNetworksCmd).Standalone()

		drs_describeSourceNetworksCmd.Flags().String("filters", "", "A set of filters by which to return Source Networks.")
		drs_describeSourceNetworksCmd.Flags().String("max-results", "", "Maximum number of Source Networks to retrieve.")
		drs_describeSourceNetworksCmd.Flags().String("next-token", "", "The token of the next Source Networks to retrieve.")
	})
	drsCmd.AddCommand(drs_describeSourceNetworksCmd)
}
