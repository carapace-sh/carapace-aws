package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_createTransformerCmd = &cobra.Command{
	Use:   "create-transformer",
	Short: "Creates a transformer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_createTransformerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_createTransformerCmd).Standalone()

		b2bi_createTransformerCmd.Flags().String("client-token", "", "Reserved for future use.")
		b2bi_createTransformerCmd.Flags().String("edi-type", "", "Specifies the details for the EDI standard that is being used for the transformer.")
		b2bi_createTransformerCmd.Flags().String("file-format", "", "Specifies that the currently supported file formats for EDI transformations are `JSON` and `XML`.")
		b2bi_createTransformerCmd.Flags().String("input-conversion", "", "Specify the `InputConversion` object, which contains the format options for the inbound transformation.")
		b2bi_createTransformerCmd.Flags().String("mapping", "", "Specify the structure that contains the mapping template and its language (either XSLT or JSONATA).")
		b2bi_createTransformerCmd.Flags().String("mapping-template", "", "Specifies the mapping template for the transformer.")
		b2bi_createTransformerCmd.Flags().String("name", "", "Specifies the name of the transformer, used to identify it.")
		b2bi_createTransformerCmd.Flags().String("output-conversion", "", "A structure that contains the `OutputConversion` object, which contains the format options for the outbound transformation.")
		b2bi_createTransformerCmd.Flags().String("sample-document", "", "Specifies a sample EDI document that is used by a transformer as a guide for processing the EDI data.")
		b2bi_createTransformerCmd.Flags().String("sample-documents", "", "Specify a structure that contains the Amazon S3 bucket and an array of the corresponding keys used to identify the location for your sample documents.")
		b2bi_createTransformerCmd.Flags().String("tags", "", "Specifies the key-value pairs assigned to ARNs that you can use to group and search for resources by type.")
		b2bi_createTransformerCmd.MarkFlagRequired("name")
	})
	b2biCmd.AddCommand(b2bi_createTransformerCmd)
}
