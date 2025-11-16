package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_updateFileCacheCmd = &cobra.Command{
	Use:   "update-file-cache",
	Short: "Updates the configuration of an existing Amazon File Cache resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_updateFileCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_updateFileCacheCmd).Standalone()

		fsx_updateFileCacheCmd.Flags().String("client-request-token", "", "")
		fsx_updateFileCacheCmd.Flags().String("file-cache-id", "", "The ID of the cache that you are updating.")
		fsx_updateFileCacheCmd.Flags().String("lustre-configuration", "", "The configuration updates for an Amazon File Cache resource.")
		fsx_updateFileCacheCmd.MarkFlagRequired("file-cache-id")
	})
	fsxCmd.AddCommand(fsx_updateFileCacheCmd)
}
