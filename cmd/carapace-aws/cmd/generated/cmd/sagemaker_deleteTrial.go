package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteTrialCmd = &cobra.Command{
	Use:   "delete-trial",
	Short: "Deletes the specified trial.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteTrialCmd).Standalone()

	sagemaker_deleteTrialCmd.Flags().String("trial-name", "", "The name of the trial to delete.")
	sagemaker_deleteTrialCmd.MarkFlagRequired("trial-name")
	sagemakerCmd.AddCommand(sagemaker_deleteTrialCmd)
}
