package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_acceptSharedDirectoryCmd = &cobra.Command{
	Use:   "accept-shared-directory",
	Short: "Accepts a directory sharing request that was sent from the directory owner account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_acceptSharedDirectoryCmd).Standalone()

	ds_acceptSharedDirectoryCmd.Flags().String("shared-directory-id", "", "Identifier of the shared directory in the directory consumer account.")
	ds_acceptSharedDirectoryCmd.MarkFlagRequired("shared-directory-id")
	dsCmd.AddCommand(ds_acceptSharedDirectoryCmd)
}
