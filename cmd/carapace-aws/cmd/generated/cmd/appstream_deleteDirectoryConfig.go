package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteDirectoryConfigCmd = &cobra.Command{
	Use:   "delete-directory-config",
	Short: "Deletes the specified Directory Config object from AppStream 2.0.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteDirectoryConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_deleteDirectoryConfigCmd).Standalone()

		appstream_deleteDirectoryConfigCmd.Flags().String("directory-name", "", "The name of the directory configuration.")
		appstream_deleteDirectoryConfigCmd.MarkFlagRequired("directory-name")
	})
	appstreamCmd.AddCommand(appstream_deleteDirectoryConfigCmd)
}
