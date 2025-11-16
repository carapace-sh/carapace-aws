package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_updateTransformerCmd = &cobra.Command{
	Use:   "update-transformer",
	Short: "Updates the specified parameters for a transformer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_updateTransformerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_updateTransformerCmd).Standalone()

		b2bi_updateTransformerCmd.Flags().String("edi-type", "", "Specifies the details for the EDI standard that is being used for the transformer.")
		b2bi_updateTransformerCmd.Flags().String("file-format", "", "Specifies that the currently supported file formats for EDI transformations are `JSON` and `XML`.")
		b2bi_updateTransformerCmd.Flags().String("input-conversion", "", "To update, specify the `InputConversion` object, which contains the format options for the inbound transformation.")
		b2bi_updateTransformerCmd.Flags().String("mapping", "", "Specify the structure that contains the mapping template and its language (either XSLT or JSONATA).")
		b2bi_updateTransformerCmd.Flags().String("mapping-template", "", "Specifies the mapping template for the transformer.")
		b2bi_updateTransformerCmd.Flags().String("name", "", "Specify a new name for the transformer, if you want to update it.")
		b2bi_updateTransformerCmd.Flags().String("output-conversion", "", "To update, specify the `OutputConversion` object, which contains the format options for the outbound transformation.")
		b2bi_updateTransformerCmd.Flags().String("sample-document", "", "Specifies a sample EDI document that is used by a transformer as a guide for processing the EDI data.")
		b2bi_updateTransformerCmd.Flags().String("sample-documents", "", "Specify a structure that contains the Amazon S3 bucket and an array of the corresponding keys used to identify the location for your sample documents.")
		b2bi_updateTransformerCmd.Flags().String("status", "", "Specifies the transformer's status.")
		b2bi_updateTransformerCmd.Flags().String("transformer-id", "", "Specifies the system-assigned unique identifier for the transformer.")
		b2bi_updateTransformerCmd.MarkFlagRequired("transformer-id")
	})
	b2biCmd.AddCommand(b2bi_updateTransformerCmd)
}
