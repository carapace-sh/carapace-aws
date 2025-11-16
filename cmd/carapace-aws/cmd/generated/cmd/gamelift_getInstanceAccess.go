package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_getInstanceAccessCmd = &cobra.Command{
	Use:   "get-instance-access",
	Short: "**This API works with the following fleet types:** EC2\n\nRequests authorization to remotely connect to an instance in an Amazon GameLift Servers managed fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_getInstanceAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_getInstanceAccessCmd).Standalone()

		gamelift_getInstanceAccessCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet that contains the instance you want to access.")
		gamelift_getInstanceAccessCmd.Flags().String("instance-id", "", "A unique identifier for the instance you want to access.")
		gamelift_getInstanceAccessCmd.MarkFlagRequired("fleet-id")
		gamelift_getInstanceAccessCmd.MarkFlagRequired("instance-id")
	})
	gameliftCmd.AddCommand(gamelift_getInstanceAccessCmd)
}
