package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeFileCachesCmd = &cobra.Command{
	Use:   "describe-file-caches",
	Short: "Returns the description of a specific Amazon File Cache resource, if a `FileCacheIds` value is provided for that cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeFileCachesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_describeFileCachesCmd).Standalone()

		fsx_describeFileCachesCmd.Flags().String("file-cache-ids", "", "IDs of the caches whose descriptions you want to retrieve (String).")
		fsx_describeFileCachesCmd.Flags().String("max-results", "", "")
		fsx_describeFileCachesCmd.Flags().String("next-token", "", "")
	})
	fsxCmd.AddCommand(fsx_describeFileCachesCmd)
}
