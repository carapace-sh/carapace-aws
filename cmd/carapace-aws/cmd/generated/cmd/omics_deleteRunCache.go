package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_deleteRunCacheCmd = &cobra.Command{
	Use:   "delete-run-cache",
	Short: "Deletes a run cache and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_deleteRunCacheCmd).Standalone()

	omics_deleteRunCacheCmd.Flags().String("id", "", "Run cache identifier for the cache you want to delete.")
	omics_deleteRunCacheCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_deleteRunCacheCmd)
}
