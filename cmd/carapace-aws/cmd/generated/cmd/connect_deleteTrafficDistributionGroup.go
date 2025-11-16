package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteTrafficDistributionGroupCmd = &cobra.Command{
	Use:   "delete-traffic-distribution-group",
	Short: "Deletes a traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteTrafficDistributionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_deleteTrafficDistributionGroupCmd).Standalone()

		connect_deleteTrafficDistributionGroupCmd.Flags().String("traffic-distribution-group-id", "", "The identifier of the traffic distribution group.")
		connect_deleteTrafficDistributionGroupCmd.MarkFlagRequired("traffic-distribution-group-id")
	})
	connectCmd.AddCommand(connect_deleteTrafficDistributionGroupCmd)
}
