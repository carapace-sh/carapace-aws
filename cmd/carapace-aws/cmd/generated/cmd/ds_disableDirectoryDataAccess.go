package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_disableDirectoryDataAccessCmd = &cobra.Command{
	Use:   "disable-directory-data-access",
	Short: "Deactivates access to directory data via the Directory Service Data API for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_disableDirectoryDataAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_disableDirectoryDataAccessCmd).Standalone()

		ds_disableDirectoryDataAccessCmd.Flags().String("directory-id", "", "The directory identifier.")
		ds_disableDirectoryDataAccessCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_disableDirectoryDataAccessCmd)
}
