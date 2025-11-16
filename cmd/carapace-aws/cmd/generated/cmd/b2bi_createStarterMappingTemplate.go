package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_createStarterMappingTemplateCmd = &cobra.Command{
	Use:   "create-starter-mapping-template",
	Short: "Amazon Web Services B2B Data Interchange uses a mapping template in JSONata or XSLT format to transform a customer input file into a JSON or XML file that can be converted to EDI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_createStarterMappingTemplateCmd).Standalone()

	b2bi_createStarterMappingTemplateCmd.Flags().String("mapping-type", "", "Specify the format for the mapping template: either JSONATA or XSLT.")
	b2bi_createStarterMappingTemplateCmd.Flags().String("output-sample-location", "", "Specify the location of the sample EDI file that is used to generate the mapping template.")
	b2bi_createStarterMappingTemplateCmd.Flags().String("template-details", "", "Describes the details needed for generating the template.")
	b2bi_createStarterMappingTemplateCmd.MarkFlagRequired("mapping-type")
	b2bi_createStarterMappingTemplateCmd.MarkFlagRequired("template-details")
	b2biCmd.AddCommand(b2bi_createStarterMappingTemplateCmd)
}
