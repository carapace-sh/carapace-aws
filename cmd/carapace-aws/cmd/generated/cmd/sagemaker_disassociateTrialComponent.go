package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_disassociateTrialComponentCmd = &cobra.Command{
	Use:   "disassociate-trial-component",
	Short: "Disassociates a trial component from a trial.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_disassociateTrialComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_disassociateTrialComponentCmd).Standalone()

		sagemaker_disassociateTrialComponentCmd.Flags().String("trial-component-name", "", "The name of the component to disassociate from the trial.")
		sagemaker_disassociateTrialComponentCmd.Flags().String("trial-name", "", "The name of the trial to disassociate from.")
		sagemaker_disassociateTrialComponentCmd.MarkFlagRequired("trial-component-name")
		sagemaker_disassociateTrialComponentCmd.MarkFlagRequired("trial-name")
	})
	sagemakerCmd.AddCommand(sagemaker_disassociateTrialComponentCmd)
}
