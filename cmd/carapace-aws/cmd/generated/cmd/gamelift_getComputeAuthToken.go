package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_getComputeAuthTokenCmd = &cobra.Command{
	Use:   "get-compute-auth-token",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRequests an authentication token from Amazon GameLift Servers for a compute resource in an Amazon GameLift Servers fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_getComputeAuthTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_getComputeAuthTokenCmd).Standalone()

		gamelift_getComputeAuthTokenCmd.Flags().String("compute-name", "", "The name of the compute resource you are requesting the authentication token for.")
		gamelift_getComputeAuthTokenCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet that the compute is registered to.")
		gamelift_getComputeAuthTokenCmd.MarkFlagRequired("compute-name")
		gamelift_getComputeAuthTokenCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_getComputeAuthTokenCmd)
}
