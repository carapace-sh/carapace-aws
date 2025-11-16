package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_testParsingCmd = &cobra.Command{
	Use:   "test-parsing",
	Short: "Parses the input EDI (electronic data interchange) file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_testParsingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_testParsingCmd).Standalone()

		b2bi_testParsingCmd.Flags().String("advanced-options", "", "Specifies advanced options for parsing the input EDI file.")
		b2bi_testParsingCmd.Flags().String("edi-type", "", "Specifies the details for the EDI standard that is being used for the transformer.")
		b2bi_testParsingCmd.Flags().String("file-format", "", "Specifies that the currently supported file formats for EDI transformations are `JSON` and `XML`.")
		b2bi_testParsingCmd.Flags().String("input-file", "", "Specifies an `S3Location` object, which contains the Amazon S3 bucket and prefix for the location of the input file.")
		b2bi_testParsingCmd.MarkFlagRequired("edi-type")
		b2bi_testParsingCmd.MarkFlagRequired("file-format")
		b2bi_testParsingCmd.MarkFlagRequired("input-file")
	})
	b2biCmd.AddCommand(b2bi_testParsingCmd)
}
