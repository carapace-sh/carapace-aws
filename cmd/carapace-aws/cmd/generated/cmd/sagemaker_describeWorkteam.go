package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeWorkteamCmd = &cobra.Command{
	Use:   "describe-workteam",
	Short: "Gets information about a specific work team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeWorkteamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeWorkteamCmd).Standalone()

		sagemaker_describeWorkteamCmd.Flags().String("workteam-name", "", "The name of the work team to return a description of.")
		sagemaker_describeWorkteamCmd.MarkFlagRequired("workteam-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeWorkteamCmd)
}
