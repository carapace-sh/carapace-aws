package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeAutoMljobV2Cmd = &cobra.Command{
	Use:   "describe-auto-mljob-v2",
	Short: "Returns information about an AutoML job created by calling [CreateAutoMLJobV2](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html) or [CreateAutoMLJob](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeAutoMljobV2Cmd).Standalone()

	sagemaker_describeAutoMljobV2Cmd.Flags().String("auto-mljob-name", "", "Requests information about an AutoML job V2 using its unique name.")
	sagemaker_describeAutoMljobV2Cmd.MarkFlagRequired("auto-mljob-name")
	sagemakerCmd.AddCommand(sagemaker_describeAutoMljobV2Cmd)
}
