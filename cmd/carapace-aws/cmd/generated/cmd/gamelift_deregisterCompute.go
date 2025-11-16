package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deregisterComputeCmd = &cobra.Command{
	Use:   "deregister-compute",
	Short: "**This API works with the following fleet types:** Anywhere\n\nRemoves a compute resource from an Anywhere fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deregisterComputeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_deregisterComputeCmd).Standalone()

		gamelift_deregisterComputeCmd.Flags().String("compute-name", "", "The unique identifier of the compute resource to deregister.")
		gamelift_deregisterComputeCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet the compute resource is currently registered to.")
		gamelift_deregisterComputeCmd.MarkFlagRequired("compute-name")
		gamelift_deregisterComputeCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_deregisterComputeCmd)
}
