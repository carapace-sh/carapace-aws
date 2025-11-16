package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAssetBundleExportJobCmd = &cobra.Command{
	Use:   "describe-asset-bundle-export-job",
	Short: "Describes an existing export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAssetBundleExportJobCmd).Standalone()

	quicksight_describeAssetBundleExportJobCmd.Flags().String("asset-bundle-export-job-id", "", "The ID of the job that you want described.")
	quicksight_describeAssetBundleExportJobCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account the export job is executed in.")
	quicksight_describeAssetBundleExportJobCmd.MarkFlagRequired("asset-bundle-export-job-id")
	quicksight_describeAssetBundleExportJobCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeAssetBundleExportJobCmd)
}
