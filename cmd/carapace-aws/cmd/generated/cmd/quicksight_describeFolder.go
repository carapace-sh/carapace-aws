package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeFolderCmd = &cobra.Command{
	Use:   "describe-folder",
	Short: "Describes a folder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeFolderCmd).Standalone()

	quicksight_describeFolderCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the folder.")
	quicksight_describeFolderCmd.Flags().String("folder-id", "", "The ID of the folder.")
	quicksight_describeFolderCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeFolderCmd.MarkFlagRequired("folder-id")
	quicksightCmd.AddCommand(quicksight_describeFolderCmd)
}
