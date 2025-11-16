package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createTrafficDistributionGroupCmd = &cobra.Command{
	Use:   "create-traffic-distribution-group",
	Short: "Creates a traffic distribution group given an Amazon Connect instance that has been replicated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createTrafficDistributionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createTrafficDistributionGroupCmd).Standalone()

		connect_createTrafficDistributionGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_createTrafficDistributionGroupCmd.Flags().String("description", "", "A description for the traffic distribution group.")
		connect_createTrafficDistributionGroupCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance that has been replicated.")
		connect_createTrafficDistributionGroupCmd.Flags().String("name", "", "The name for the traffic distribution group.")
		connect_createTrafficDistributionGroupCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createTrafficDistributionGroupCmd.MarkFlagRequired("instance-id")
		connect_createTrafficDistributionGroupCmd.MarkFlagRequired("name")
	})
	connectCmd.AddCommand(connect_createTrafficDistributionGroupCmd)
}
