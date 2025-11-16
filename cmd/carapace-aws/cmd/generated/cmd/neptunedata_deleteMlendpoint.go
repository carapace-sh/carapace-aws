package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_deleteMlendpointCmd = &cobra.Command{
	Use:   "delete-mlendpoint",
	Short: "Cancels the creation of a Neptune ML inference endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_deleteMlendpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_deleteMlendpointCmd).Standalone()

		neptunedata_deleteMlendpointCmd.Flags().Bool("clean", false, "If this flag is set to `TRUE`, all Neptune ML S3 artifacts should be deleted when the job is stopped.")
		neptunedata_deleteMlendpointCmd.Flags().String("id", "", "The unique identifier of the inference endpoint.")
		neptunedata_deleteMlendpointCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role providing Neptune access to SageMaker and Amazon S3 resources.")
		neptunedata_deleteMlendpointCmd.Flags().Bool("no-clean", false, "If this flag is set to `TRUE`, all Neptune ML S3 artifacts should be deleted when the job is stopped.")
		neptunedata_deleteMlendpointCmd.MarkFlagRequired("id")
		neptunedata_deleteMlendpointCmd.Flag("no-clean").Hidden = true
	})
	neptunedataCmd.AddCommand(neptunedata_deleteMlendpointCmd)
}
