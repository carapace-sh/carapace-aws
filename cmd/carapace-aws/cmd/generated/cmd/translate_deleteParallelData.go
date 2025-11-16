package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_deleteParallelDataCmd = &cobra.Command{
	Use:   "delete-parallel-data",
	Short: "Deletes a parallel data resource in Amazon Translate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_deleteParallelDataCmd).Standalone()

	translate_deleteParallelDataCmd.Flags().String("name", "", "The name of the parallel data resource that is being deleted.")
	translate_deleteParallelDataCmd.MarkFlagRequired("name")
	translateCmd.AddCommand(translate_deleteParallelDataCmd)
}
