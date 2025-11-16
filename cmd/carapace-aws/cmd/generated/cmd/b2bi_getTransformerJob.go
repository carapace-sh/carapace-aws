package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_getTransformerJobCmd = &cobra.Command{
	Use:   "get-transformer-job",
	Short: "Returns the details of the transformer run, based on the Transformer job ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_getTransformerJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_getTransformerJobCmd).Standalone()

		b2bi_getTransformerJobCmd.Flags().String("transformer-id", "", "Specifies the system-assigned unique identifier for the transformer.")
		b2bi_getTransformerJobCmd.Flags().String("transformer-job-id", "", "Specifies the unique, system-generated identifier for a transformer run.")
		b2bi_getTransformerJobCmd.MarkFlagRequired("transformer-id")
		b2bi_getTransformerJobCmd.MarkFlagRequired("transformer-job-id")
	})
	b2biCmd.AddCommand(b2bi_getTransformerJobCmd)
}
