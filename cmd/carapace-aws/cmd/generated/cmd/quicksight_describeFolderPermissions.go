package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeFolderPermissionsCmd = &cobra.Command{
	Use:   "describe-folder-permissions",
	Short: "Describes permissions for a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeFolderPermissionsCmd).Standalone()

	quicksight_describeFolderPermissionsCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
	quicksight_describeFolderPermissionsCmd.Flags().String("folder-id", "", "The ID of the folder.")
	quicksight_describeFolderPermissionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	quicksight_describeFolderPermissionsCmd.Flags().String("namespace", "", "The namespace of the folder whose permissions you want described.")
	quicksight_describeFolderPermissionsCmd.Flags().String("next-token", "", "A pagination token for the next set of results.")
	quicksight_describeFolderPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeFolderPermissionsCmd.MarkFlagRequired("folder-id")
	quicksightCmd.AddCommand(quicksight_describeFolderPermissionsCmd)
}
