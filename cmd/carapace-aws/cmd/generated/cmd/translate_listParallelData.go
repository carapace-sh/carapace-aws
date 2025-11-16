package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_listParallelDataCmd = &cobra.Command{
	Use:   "list-parallel-data",
	Short: "Provides a list of your parallel data resources in Amazon Translate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_listParallelDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_listParallelDataCmd).Standalone()

		translate_listParallelDataCmd.Flags().String("max-results", "", "The maximum number of parallel data resources returned for each request.")
		translate_listParallelDataCmd.Flags().String("next-token", "", "A string that specifies the next page of results to return in a paginated response.")
	})
	translateCmd.AddCommand(translate_listParallelDataCmd)
}
