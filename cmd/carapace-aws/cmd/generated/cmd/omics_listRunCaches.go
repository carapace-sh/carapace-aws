package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listRunCachesCmd = &cobra.Command{
	Use:   "list-run-caches",
	Short: "Retrieves a list of your run caches and the metadata for each cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listRunCachesCmd).Standalone()

	omics_listRunCachesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	omics_listRunCachesCmd.Flags().String("starting-token", "", "Optional pagination token returned from a prior call to the `ListRunCaches` API operation.")
	omicsCmd.AddCommand(omics_listRunCachesCmd)
}
