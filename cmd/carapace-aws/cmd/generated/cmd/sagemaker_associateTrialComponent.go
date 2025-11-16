package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_associateTrialComponentCmd = &cobra.Command{
	Use:   "associate-trial-component",
	Short: "Associates a trial component with a trial.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_associateTrialComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_associateTrialComponentCmd).Standalone()

		sagemaker_associateTrialComponentCmd.Flags().String("trial-component-name", "", "The name of the component to associated with the trial.")
		sagemaker_associateTrialComponentCmd.Flags().String("trial-name", "", "The name of the trial to associate with.")
		sagemaker_associateTrialComponentCmd.MarkFlagRequired("trial-component-name")
		sagemaker_associateTrialComponentCmd.MarkFlagRequired("trial-name")
	})
	sagemakerCmd.AddCommand(sagemaker_associateTrialComponentCmd)
}
