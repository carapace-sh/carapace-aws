package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateTrialCmd = &cobra.Command{
	Use:   "update-trial",
	Short: "Updates the display name of a trial.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateTrialCmd).Standalone()

	sagemaker_updateTrialCmd.Flags().String("display-name", "", "The name of the trial as displayed.")
	sagemaker_updateTrialCmd.Flags().String("trial-name", "", "The name of the trial to update.")
	sagemaker_updateTrialCmd.MarkFlagRequired("trial-name")
	sagemakerCmd.AddCommand(sagemaker_updateTrialCmd)
}
