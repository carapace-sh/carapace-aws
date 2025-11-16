package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateTrafficDistributionGroupUserCmd = &cobra.Command{
	Use:   "disassociate-traffic-distribution-group-user",
	Short: "Disassociates an agent from a traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateTrafficDistributionGroupUserCmd).Standalone()

	connect_disassociateTrafficDistributionGroupUserCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_disassociateTrafficDistributionGroupUserCmd.Flags().String("traffic-distribution-group-id", "", "The identifier of the traffic distribution group.")
	connect_disassociateTrafficDistributionGroupUserCmd.Flags().String("user-id", "", "The identifier for the user.")
	connect_disassociateTrafficDistributionGroupUserCmd.MarkFlagRequired("instance-id")
	connect_disassociateTrafficDistributionGroupUserCmd.MarkFlagRequired("traffic-distribution-group-id")
	connect_disassociateTrafficDistributionGroupUserCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_disassociateTrafficDistributionGroupUserCmd)
}
