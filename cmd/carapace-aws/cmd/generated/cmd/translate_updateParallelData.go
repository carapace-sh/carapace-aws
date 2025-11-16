package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_updateParallelDataCmd = &cobra.Command{
	Use:   "update-parallel-data",
	Short: "Updates a previously created parallel data resource by importing a new input file from Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_updateParallelDataCmd).Standalone()

	translate_updateParallelDataCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	translate_updateParallelDataCmd.Flags().String("description", "", "A custom description for the parallel data resource in Amazon Translate.")
	translate_updateParallelDataCmd.Flags().String("name", "", "The name of the parallel data resource being updated.")
	translate_updateParallelDataCmd.Flags().String("parallel-data-config", "", "Specifies the format and S3 location of the parallel data input file.")
	translate_updateParallelDataCmd.MarkFlagRequired("client-token")
	translate_updateParallelDataCmd.MarkFlagRequired("name")
	translate_updateParallelDataCmd.MarkFlagRequired("parallel-data-config")
	translateCmd.AddCommand(translate_updateParallelDataCmd)
}
