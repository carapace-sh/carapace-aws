package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_startSessionCmd = &cobra.Command{
	Use:   "start-session",
	Short: "Initiates a remote connection session between a local integrated development environments (IDEs) and a remote SageMaker space.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_startSessionCmd).Standalone()

	sagemaker_startSessionCmd.Flags().String("resource-identifier", "", "The Amazon Resource Name (ARN) of the resource to which the remote connection will be established.")
	sagemaker_startSessionCmd.MarkFlagRequired("resource-identifier")
	sagemakerCmd.AddCommand(sagemaker_startSessionCmd)
}
