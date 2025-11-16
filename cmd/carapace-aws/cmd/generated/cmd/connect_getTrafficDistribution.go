package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getTrafficDistributionCmd = &cobra.Command{
	Use:   "get-traffic-distribution",
	Short: "Retrieves the current traffic distribution for a given traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getTrafficDistributionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getTrafficDistributionCmd).Standalone()

		connect_getTrafficDistributionCmd.Flags().String("id", "", "The identifier of the traffic distribution group.")
		connect_getTrafficDistributionCmd.MarkFlagRequired("id")
	})
	connectCmd.AddCommand(connect_getTrafficDistributionCmd)
}
