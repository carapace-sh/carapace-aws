package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_updateRunCacheCmd = &cobra.Command{
	Use:   "update-run-cache",
	Short: "Updates a run cache using its ID and returns a response with no body if the operation is successful.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_updateRunCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(omics_updateRunCacheCmd).Standalone()

		omics_updateRunCacheCmd.Flags().String("cache-behavior", "", "Update the default run cache behavior.")
		omics_updateRunCacheCmd.Flags().String("description", "", "Update the run cache description.")
		omics_updateRunCacheCmd.Flags().String("id", "", "The identifier of the run cache you want to update.")
		omics_updateRunCacheCmd.Flags().String("name", "", "Update the name of the run cache.")
		omics_updateRunCacheCmd.MarkFlagRequired("id")
	})
	omicsCmd.AddCommand(omics_updateRunCacheCmd)
}
