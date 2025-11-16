package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_enableDirectoryDataAccessCmd = &cobra.Command{
	Use:   "enable-directory-data-access",
	Short: "Enables access to directory data via the Directory Service Data API for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_enableDirectoryDataAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_enableDirectoryDataAccessCmd).Standalone()

		ds_enableDirectoryDataAccessCmd.Flags().String("directory-id", "", "The directory identifier.")
		ds_enableDirectoryDataAccessCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_enableDirectoryDataAccessCmd)
}
