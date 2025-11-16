package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deleteDirectoryCmd = &cobra.Command{
	Use:   "delete-directory",
	Short: "Deletes an Directory Service directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deleteDirectoryCmd).Standalone()

	ds_deleteDirectoryCmd.Flags().String("directory-id", "", "The identifier of the directory to delete.")
	ds_deleteDirectoryCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_deleteDirectoryCmd)
}
