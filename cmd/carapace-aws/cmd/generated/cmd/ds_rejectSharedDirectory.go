package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_rejectSharedDirectoryCmd = &cobra.Command{
	Use:   "reject-shared-directory",
	Short: "Rejects a directory sharing request that was sent from the directory owner account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_rejectSharedDirectoryCmd).Standalone()

	ds_rejectSharedDirectoryCmd.Flags().String("shared-directory-id", "", "Identifier of the shared directory in the directory consumer account.")
	ds_rejectSharedDirectoryCmd.MarkFlagRequired("shared-directory-id")
	dsCmd.AddCommand(ds_rejectSharedDirectoryCmd)
}
