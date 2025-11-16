package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeTrialComponentCmd = &cobra.Command{
	Use:   "describe-trial-component",
	Short: "Provides a list of a trials component's properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeTrialComponentCmd).Standalone()

	sagemaker_describeTrialComponentCmd.Flags().String("trial-component-name", "", "The name of the trial component to describe.")
	sagemaker_describeTrialComponentCmd.MarkFlagRequired("trial-component-name")
	sagemakerCmd.AddCommand(sagemaker_describeTrialComponentCmd)
}
