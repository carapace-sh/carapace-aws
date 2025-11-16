package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeDirectoryConfigsCmd = &cobra.Command{
	Use:   "describe-directory-configs",
	Short: "Retrieves a list that describes one or more specified Directory Config objects for AppStream 2.0, if the names for these objects are provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeDirectoryConfigsCmd).Standalone()

	appstream_describeDirectoryConfigsCmd.Flags().String("directory-names", "", "The directory names.")
	appstream_describeDirectoryConfigsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_describeDirectoryConfigsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	appstreamCmd.AddCommand(appstream_describeDirectoryConfigsCmd)
}
