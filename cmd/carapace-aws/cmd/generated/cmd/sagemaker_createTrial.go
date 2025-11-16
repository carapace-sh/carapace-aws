package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createTrialCmd = &cobra.Command{
	Use:   "create-trial",
	Short: "Creates an SageMaker *trial*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createTrialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createTrialCmd).Standalone()

		sagemaker_createTrialCmd.Flags().String("display-name", "", "The name of the trial as displayed.")
		sagemaker_createTrialCmd.Flags().String("experiment-name", "", "The name of the experiment to associate the trial with.")
		sagemaker_createTrialCmd.Flags().String("metadata-properties", "", "")
		sagemaker_createTrialCmd.Flags().String("tags", "", "A list of tags to associate with the trial.")
		sagemaker_createTrialCmd.Flags().String("trial-name", "", "The name of the trial.")
		sagemaker_createTrialCmd.MarkFlagRequired("experiment-name")
		sagemaker_createTrialCmd.MarkFlagRequired("trial-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createTrialCmd)
}
