package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeComputeCmd = &cobra.Command{
	Use:   "describe-compute",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves properties for a specific compute resource in an Amazon GameLift Servers fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeComputeCmd).Standalone()

	gamelift_describeComputeCmd.Flags().String("compute-name", "", "The unique identifier of the compute resource to retrieve properties for.")
	gamelift_describeComputeCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet that the compute belongs to.")
	gamelift_describeComputeCmd.MarkFlagRequired("compute-name")
	gamelift_describeComputeCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_describeComputeCmd)
}
