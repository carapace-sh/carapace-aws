package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_testConversionCmd = &cobra.Command{
	Use:   "test-conversion",
	Short: "This operation mimics the latter half of a typical Outbound EDI request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_testConversionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_testConversionCmd).Standalone()

		b2bi_testConversionCmd.Flags().String("source", "", "Specify the source file for an outbound EDI request.")
		b2bi_testConversionCmd.Flags().String("target", "", "Specify the format (X12 is the only currently supported format), and other details for the conversion target.")
		b2bi_testConversionCmd.MarkFlagRequired("source")
		b2bi_testConversionCmd.MarkFlagRequired("target")
	})
	b2biCmd.AddCommand(b2bi_testConversionCmd)
}
