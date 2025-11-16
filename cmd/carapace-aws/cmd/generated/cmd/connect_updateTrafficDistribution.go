package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateTrafficDistributionCmd = &cobra.Command{
	Use:   "update-traffic-distribution",
	Short: "Updates the traffic distribution for a given traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateTrafficDistributionCmd).Standalone()

	connect_updateTrafficDistributionCmd.Flags().String("agent-config", "", "The distribution of agents between the instance and its replica(s).")
	connect_updateTrafficDistributionCmd.Flags().String("id", "", "The identifier of the traffic distribution group.")
	connect_updateTrafficDistributionCmd.Flags().String("sign-in-config", "", "The distribution that determines which Amazon Web Services Regions should be used to sign in agents in to both the instance and its replica(s).")
	connect_updateTrafficDistributionCmd.Flags().String("telephony-config", "", "The distribution of traffic between the instance and its replica(s).")
	connect_updateTrafficDistributionCmd.MarkFlagRequired("id")
	connectCmd.AddCommand(connect_updateTrafficDistributionCmd)
}
