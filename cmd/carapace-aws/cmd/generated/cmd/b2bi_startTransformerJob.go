package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_startTransformerJobCmd = &cobra.Command{
	Use:   "start-transformer-job",
	Short: "Runs a job, using a transformer, to parse input EDI (electronic data interchange) file into the output structures used by Amazon Web Services B2B Data Interchange.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_startTransformerJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_startTransformerJobCmd).Standalone()

		b2bi_startTransformerJobCmd.Flags().String("client-token", "", "Reserved for future use.")
		b2bi_startTransformerJobCmd.Flags().String("input-file", "", "Specifies the location of the input file for the transformation.")
		b2bi_startTransformerJobCmd.Flags().String("output-location", "", "Specifies the location of the output file for the transformation.")
		b2bi_startTransformerJobCmd.Flags().String("transformer-id", "", "Specifies the system-assigned unique identifier for the transformer.")
		b2bi_startTransformerJobCmd.MarkFlagRequired("input-file")
		b2bi_startTransformerJobCmd.MarkFlagRequired("output-location")
		b2bi_startTransformerJobCmd.MarkFlagRequired("transformer-id")
	})
	b2biCmd.AddCommand(b2bi_startTransformerJobCmd)
}
