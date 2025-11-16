package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_unshareDirectoryCmd = &cobra.Command{
	Use:   "unshare-directory",
	Short: "Stops the directory sharing between the directory owner and consumer accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_unshareDirectoryCmd).Standalone()

	ds_unshareDirectoryCmd.Flags().String("directory-id", "", "The identifier of the Managed Microsoft AD directory that you want to stop sharing.")
	ds_unshareDirectoryCmd.Flags().String("unshare-target", "", "Identifier for the directory consumer account with whom the directory has to be unshared.")
	ds_unshareDirectoryCmd.MarkFlagRequired("directory-id")
	ds_unshareDirectoryCmd.MarkFlagRequired("unshare-target")
	dsCmd.AddCommand(ds_unshareDirectoryCmd)
}
