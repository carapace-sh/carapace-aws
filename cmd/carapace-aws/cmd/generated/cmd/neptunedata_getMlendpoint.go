package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getMlendpointCmd = &cobra.Command{
	Use:   "get-mlendpoint",
	Short: "Retrieves details about an inference endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getMlendpointCmd).Standalone()

	neptunedata_getMlendpointCmd.Flags().String("id", "", "The unique identifier of the inference endpoint.")
	neptunedata_getMlendpointCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_getMlendpointCmd.MarkFlagRequired("id")
	neptunedataCmd.AddCommand(neptunedata_getMlendpointCmd)
}
