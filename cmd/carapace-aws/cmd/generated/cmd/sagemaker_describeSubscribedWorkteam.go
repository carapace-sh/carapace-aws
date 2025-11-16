package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeSubscribedWorkteamCmd = &cobra.Command{
	Use:   "describe-subscribed-workteam",
	Short: "Gets information about a work team provided by a vendor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeSubscribedWorkteamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeSubscribedWorkteamCmd).Standalone()

		sagemaker_describeSubscribedWorkteamCmd.Flags().String("workteam-arn", "", "The Amazon Resource Name (ARN) of the subscribed work team to describe.")
		sagemaker_describeSubscribedWorkteamCmd.MarkFlagRequired("workteam-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_describeSubscribedWorkteamCmd)
}
