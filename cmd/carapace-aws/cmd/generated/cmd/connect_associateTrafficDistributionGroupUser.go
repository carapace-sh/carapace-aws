package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateTrafficDistributionGroupUserCmd = &cobra.Command{
	Use:   "associate-traffic-distribution-group-user",
	Short: "Associates an agent with a traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateTrafficDistributionGroupUserCmd).Standalone()

	connect_associateTrafficDistributionGroupUserCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateTrafficDistributionGroupUserCmd.Flags().String("traffic-distribution-group-id", "", "The identifier of the traffic distribution group.")
	connect_associateTrafficDistributionGroupUserCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_associateTrafficDistributionGroupUserCmd.MarkFlagRequired("instance-id")
	connect_associateTrafficDistributionGroupUserCmd.MarkFlagRequired("traffic-distribution-group-id")
	connect_associateTrafficDistributionGroupUserCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_associateTrafficDistributionGroupUserCmd)
}
