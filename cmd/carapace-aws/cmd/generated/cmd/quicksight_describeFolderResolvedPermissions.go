package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeFolderResolvedPermissionsCmd = &cobra.Command{
	Use:   "describe-folder-resolved-permissions",
	Short: "Describes the folder resolved permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeFolderResolvedPermissionsCmd).Standalone()

	quicksight_describeFolderResolvedPermissionsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
	quicksight_describeFolderResolvedPermissionsCmd.Flags().String("folder-id", "", "The ID of the folder.")
	quicksight_describeFolderResolvedPermissionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_describeFolderResolvedPermissionsCmd.Flags().String("namespace", "", "The namespace of the folder whose permissions you want described.")
	quicksight_describeFolderResolvedPermissionsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	quicksight_describeFolderResolvedPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeFolderResolvedPermissionsCmd.MarkFlagRequired("folder-id")
	quicksightCmd.AddCommand(quicksight_describeFolderResolvedPermissionsCmd)
}
