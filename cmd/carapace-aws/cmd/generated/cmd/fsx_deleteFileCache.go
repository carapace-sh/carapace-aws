package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_deleteFileCacheCmd = &cobra.Command{
	Use:   "delete-file-cache",
	Short: "Deletes an Amazon File Cache resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_deleteFileCacheCmd).Standalone()

	fsx_deleteFileCacheCmd.Flags().String("client-request-token", "", "")
	fsx_deleteFileCacheCmd.Flags().String("file-cache-id", "", "The ID of the cache that's being deleted.")
	fsx_deleteFileCacheCmd.MarkFlagRequired("file-cache-id")
	fsxCmd.AddCommand(fsx_deleteFileCacheCmd)
}
