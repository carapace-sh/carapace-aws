package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getRunCacheCmd = &cobra.Command{
	Use:   "get-run-cache",
	Short: "Retrieves detailed information about the specified run cache using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getRunCacheCmd).Standalone()

	omics_getRunCacheCmd.Flags().String("id", "", "The identifier of the run cache to retrieve.")
	omics_getRunCacheCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_getRunCacheCmd)
}
