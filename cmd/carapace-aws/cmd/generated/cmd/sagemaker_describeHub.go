package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeHubCmd = &cobra.Command{
	Use:   "describe-hub",
	Short: "Describes a hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeHubCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeHubCmd).Standalone()

		sagemaker_describeHubCmd.Flags().String("hub-name", "", "The name of the hub to describe.")
		sagemaker_describeHubCmd.MarkFlagRequired("hub-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeHubCmd)
}
