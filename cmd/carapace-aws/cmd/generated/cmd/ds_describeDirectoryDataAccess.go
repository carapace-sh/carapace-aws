package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeDirectoryDataAccessCmd = &cobra.Command{
	Use:   "describe-directory-data-access",
	Short: "Obtains status of directory data access enablement through the Directory Service Data API for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeDirectoryDataAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeDirectoryDataAccessCmd).Standalone()

		ds_describeDirectoryDataAccessCmd.Flags().String("directory-id", "", "The directory identifier.")
		ds_describeDirectoryDataAccessCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_describeDirectoryDataAccessCmd)
}
