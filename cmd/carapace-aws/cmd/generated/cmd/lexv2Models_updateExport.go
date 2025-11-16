package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_updateExportCmd = &cobra.Command{
	Use:   "update-export",
	Short: "Updates the password used to protect an export zip archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_updateExportCmd).Standalone()

	lexv2Models_updateExportCmd.Flags().String("export-id", "", "The unique identifier Amazon Lex assigned to the export.")
	lexv2Models_updateExportCmd.Flags().String("file-password", "", "The new password to use to encrypt the export zip archive.")
	lexv2Models_updateExportCmd.MarkFlagRequired("export-id")
	lexv2ModelsCmd.AddCommand(lexv2Models_updateExportCmd)
}
