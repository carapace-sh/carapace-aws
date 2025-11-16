package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteTrialComponentCmd = &cobra.Command{
	Use:   "delete-trial-component",
	Short: "Deletes the specified trial component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteTrialComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteTrialComponentCmd).Standalone()

		sagemaker_deleteTrialComponentCmd.Flags().String("trial-component-name", "", "The name of the component to delete.")
		sagemaker_deleteTrialComponentCmd.MarkFlagRequired("trial-component-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteTrialComponentCmd)
}
