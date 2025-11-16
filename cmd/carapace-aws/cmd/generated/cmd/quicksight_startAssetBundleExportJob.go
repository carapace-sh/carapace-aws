package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_startAssetBundleExportJobCmd = &cobra.Command{
	Use:   "start-asset-bundle-export-job",
	Short: "Starts an Asset Bundle export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_startAssetBundleExportJobCmd).Standalone()

	quicksight_startAssetBundleExportJobCmd.Flags().String("asset-bundle-export-job-id", "", "The ID of the job.")
	quicksight_startAssetBundleExportJobCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account to export assets from.")
	quicksight_startAssetBundleExportJobCmd.Flags().String("cloud-formation-override-property-configuration", "", "An optional collection of structures that generate CloudFormation parameters to override the existing resource property values when the resource is exported to a new CloudFormation template.")
	quicksight_startAssetBundleExportJobCmd.Flags().String("export-format", "", "The export data format.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("include-all-dependencies", false, "A Boolean that determines whether all dependencies of each resource ARN are recursively exported with the job.")
	quicksight_startAssetBundleExportJobCmd.Flags().String("include-folder-members", "", "A setting that indicates whether you want to include folder assets.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("include-folder-memberships", false, "A Boolean that determines if the exported asset carries over information about the folders that the asset is a member of.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("include-permissions", false, "A Boolean that determines whether all permissions for each resource ARN are exported with the job.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("include-tags", false, "A Boolean that determines whether all tags for each resource ARN are exported with the job.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("no-include-all-dependencies", false, "A Boolean that determines whether all dependencies of each resource ARN are recursively exported with the job.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("no-include-folder-memberships", false, "A Boolean that determines if the exported asset carries over information about the folders that the asset is a member of.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("no-include-permissions", false, "A Boolean that determines whether all permissions for each resource ARN are exported with the job.")
	quicksight_startAssetBundleExportJobCmd.Flags().Bool("no-include-tags", false, "A Boolean that determines whether all tags for each resource ARN are exported with the job.")
	quicksight_startAssetBundleExportJobCmd.Flags().String("resource-arns", "", "An array of resource ARNs to export.")
	quicksight_startAssetBundleExportJobCmd.Flags().String("validation-strategy", "", "An optional parameter that determines which validation strategy to use for the export job.")
	quicksight_startAssetBundleExportJobCmd.MarkFlagRequired("asset-bundle-export-job-id")
	quicksight_startAssetBundleExportJobCmd.MarkFlagRequired("aws-account-id")
	quicksight_startAssetBundleExportJobCmd.MarkFlagRequired("export-format")
	quicksight_startAssetBundleExportJobCmd.Flag("no-include-all-dependencies").Hidden = true
	quicksight_startAssetBundleExportJobCmd.Flag("no-include-folder-memberships").Hidden = true
	quicksight_startAssetBundleExportJobCmd.Flag("no-include-permissions").Hidden = true
	quicksight_startAssetBundleExportJobCmd.Flag("no-include-tags").Hidden = true
	quicksight_startAssetBundleExportJobCmd.MarkFlagRequired("resource-arns")
	quicksightCmd.AddCommand(quicksight_startAssetBundleExportJobCmd)
}
