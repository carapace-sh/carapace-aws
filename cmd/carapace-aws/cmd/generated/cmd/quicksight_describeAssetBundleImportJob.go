package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAssetBundleImportJobCmd = &cobra.Command{
	Use:   "describe-asset-bundle-import-job",
	Short: "Describes an existing import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAssetBundleImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeAssetBundleImportJobCmd).Standalone()

		quicksight_describeAssetBundleImportJobCmd.Flags().String("asset-bundle-import-job-id", "", "The ID of the job.")
		quicksight_describeAssetBundleImportJobCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account the import job was executed in.")
		quicksight_describeAssetBundleImportJobCmd.MarkFlagRequired("asset-bundle-import-job-id")
		quicksight_describeAssetBundleImportJobCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_describeAssetBundleImportJobCmd)
}
