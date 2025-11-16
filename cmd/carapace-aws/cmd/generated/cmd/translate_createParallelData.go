package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_createParallelDataCmd = &cobra.Command{
	Use:   "create-parallel-data",
	Short: "Creates a parallel data resource in Amazon Translate by importing an input file from Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_createParallelDataCmd).Standalone()

	translate_createParallelDataCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	translate_createParallelDataCmd.Flags().String("description", "", "A custom description for the parallel data resource in Amazon Translate.")
	translate_createParallelDataCmd.Flags().String("encryption-key", "", "")
	translate_createParallelDataCmd.Flags().String("name", "", "A custom name for the parallel data resource in Amazon Translate.")
	translate_createParallelDataCmd.Flags().String("parallel-data-config", "", "Specifies the format and S3 location of the parallel data input file.")
	translate_createParallelDataCmd.Flags().String("tags", "", "Tags to be associated with this resource.")
	translate_createParallelDataCmd.MarkFlagRequired("client-token")
	translate_createParallelDataCmd.MarkFlagRequired("name")
	translate_createParallelDataCmd.MarkFlagRequired("parallel-data-config")
	translateCmd.AddCommand(translate_createParallelDataCmd)
}
