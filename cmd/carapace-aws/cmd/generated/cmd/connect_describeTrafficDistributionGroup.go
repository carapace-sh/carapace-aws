package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeTrafficDistributionGroupCmd = &cobra.Command{
	Use:   "describe-traffic-distribution-group",
	Short: "Gets details and status of a traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeTrafficDistributionGroupCmd).Standalone()

	connect_describeTrafficDistributionGroupCmd.Flags().String("traffic-distribution-group-id", "", "The identifier of the traffic distribution group.")
	connect_describeTrafficDistributionGroupCmd.MarkFlagRequired("traffic-distribution-group-id")
	connectCmd.AddCommand(connect_describeTrafficDistributionGroupCmd)
}
