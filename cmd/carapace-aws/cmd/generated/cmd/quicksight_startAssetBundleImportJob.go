package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_startAssetBundleImportJobCmd = &cobra.Command{
	Use:   "start-asset-bundle-import-job",
	Short: "Starts an Asset Bundle import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_startAssetBundleImportJobCmd).Standalone()

	quicksight_startAssetBundleImportJobCmd.Flags().String("asset-bundle-import-job-id", "", "The ID of the job.")
	quicksight_startAssetBundleImportJobCmd.Flags().String("asset-bundle-import-source", "", "The source of the asset bundle zip file that contains the data that you want to import.")
	quicksight_startAssetBundleImportJobCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account to import assets into.")
	quicksight_startAssetBundleImportJobCmd.Flags().String("failure-action", "", "The failure action for the import job.")
	quicksight_startAssetBundleImportJobCmd.Flags().String("override-parameters", "", "Optional overrides that are applied to the resource configuration before import.")
	quicksight_startAssetBundleImportJobCmd.Flags().String("override-permissions", "", "Optional permission overrides that are applied to the resource configuration before import.")
	quicksight_startAssetBundleImportJobCmd.Flags().String("override-tags", "", "Optional tag overrides that are applied to the resource configuration before import.")
	quicksight_startAssetBundleImportJobCmd.Flags().String("override-validation-strategy", "", "An optional validation strategy override for all analyses and dashboards that is applied to the resource configuration before import.")
	quicksight_startAssetBundleImportJobCmd.MarkFlagRequired("asset-bundle-import-job-id")
	quicksight_startAssetBundleImportJobCmd.MarkFlagRequired("asset-bundle-import-source")
	quicksight_startAssetBundleImportJobCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_startAssetBundleImportJobCmd)
}
