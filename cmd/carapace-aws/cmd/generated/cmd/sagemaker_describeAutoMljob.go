package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeAutoMljobCmd = &cobra.Command{
	Use:   "describe-auto-mljob",
	Short: "Returns information about an AutoML job created by calling [CreateAutoMLJob](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJob.html).\n\nAutoML jobs created by calling [CreateAutoMLJobV2](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateAutoMLJobV2.html) cannot be described by `DescribeAutoMLJob`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeAutoMljobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeAutoMljobCmd).Standalone()

		sagemaker_describeAutoMljobCmd.Flags().String("auto-mljob-name", "", "Requests information about an AutoML job using its unique name.")
		sagemaker_describeAutoMljobCmd.MarkFlagRequired("auto-mljob-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeAutoMljobCmd)
}
