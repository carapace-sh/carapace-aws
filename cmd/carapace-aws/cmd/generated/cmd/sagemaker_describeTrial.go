package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeTrialCmd = &cobra.Command{
	Use:   "describe-trial",
	Short: "Provides a list of a trial's properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeTrialCmd).Standalone()

	sagemaker_describeTrialCmd.Flags().String("trial-name", "", "The name of the trial to describe.")
	sagemaker_describeTrialCmd.MarkFlagRequired("trial-name")
	sagemakerCmd.AddCommand(sagemaker_describeTrialCmd)
}
