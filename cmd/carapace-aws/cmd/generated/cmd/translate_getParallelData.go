package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_getParallelDataCmd = &cobra.Command{
	Use:   "get-parallel-data",
	Short: "Provides information about a parallel data resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_getParallelDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_getParallelDataCmd).Standalone()

		translate_getParallelDataCmd.Flags().String("name", "", "The name of the parallel data resource that is being retrieved.")
		translate_getParallelDataCmd.MarkFlagRequired("name")
	})
	translateCmd.AddCommand(translate_getParallelDataCmd)
}
