package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_getComputeAccessCmd = &cobra.Command{
	Use:   "get-compute-access",
	Short: "**This API works with the following fleet types:** EC2, Container\n\nRequests authorization to remotely connect to a hosting resource in a Amazon GameLift Servers managed fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_getComputeAccessCmd).Standalone()

	gamelift_getComputeAccessCmd.Flags().String("compute-name", "", "A unique identifier for the compute resource that you want to connect to.")
	gamelift_getComputeAccessCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet that holds the compute resource that you want to connect to.")
	gamelift_getComputeAccessCmd.MarkFlagRequired("compute-name")
	gamelift_getComputeAccessCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_getComputeAccessCmd)
}
