package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_testMappingCmd = &cobra.Command{
	Use:   "test-mapping",
	Short: "Maps the input file according to the provided template file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_testMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_testMappingCmd).Standalone()

		b2bi_testMappingCmd.Flags().String("file-format", "", "Specifies that the currently supported file formats for EDI transformations are `JSON` and `XML`.")
		b2bi_testMappingCmd.Flags().String("input-file-content", "", "Specify the contents of the EDI (electronic data interchange) XML or JSON file that is used as input for the transform.")
		b2bi_testMappingCmd.Flags().String("mapping-template", "", "Specifies the mapping template for the transformer.")
		b2bi_testMappingCmd.MarkFlagRequired("file-format")
		b2bi_testMappingCmd.MarkFlagRequired("input-file-content")
		b2bi_testMappingCmd.MarkFlagRequired("mapping-template")
	})
	b2biCmd.AddCommand(b2bi_testMappingCmd)
}
