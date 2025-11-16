package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_generateMappingCmd = &cobra.Command{
	Use:   "generate-mapping",
	Short: "Takes sample input and output documents and uses Amazon Bedrock to generate a mapping automatically.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_generateMappingCmd).Standalone()

	b2bi_generateMappingCmd.Flags().String("input-file-content", "", "Provide the contents of a sample X12 EDI file, either in JSON or XML format, to use as a starting point for the mapping.")
	b2bi_generateMappingCmd.Flags().String("mapping-type", "", "Specify the mapping type: either `JSONATA` or `XSLT.`")
	b2bi_generateMappingCmd.Flags().String("output-file-content", "", "Provide the contents of a sample X12 EDI file, either in JSON or XML format, to use as a target for the mapping.")
	b2bi_generateMappingCmd.MarkFlagRequired("input-file-content")
	b2bi_generateMappingCmd.MarkFlagRequired("mapping-type")
	b2bi_generateMappingCmd.MarkFlagRequired("output-file-content")
	b2biCmd.AddCommand(b2bi_generateMappingCmd)
}
