package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putConformancePackCmd = &cobra.Command{
	Use:   "put-conformance-pack",
	Short: "Creates or updates a conformance pack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putConformancePackCmd).Standalone()

	config_putConformancePackCmd.Flags().String("conformance-pack-input-parameters", "", "A list of `ConformancePackInputParameter` objects.")
	config_putConformancePackCmd.Flags().String("conformance-pack-name", "", "The unique name of the conformance pack you want to deploy.")
	config_putConformancePackCmd.Flags().String("delivery-s3-bucket", "", "The name of the Amazon S3 bucket where Config stores conformance pack templates.")
	config_putConformancePackCmd.Flags().String("delivery-s3-key-prefix", "", "The prefix for the Amazon S3 bucket.")
	config_putConformancePackCmd.Flags().String("template-body", "", "A string that contains the full conformance pack template body.")
	config_putConformancePackCmd.Flags().String("template-s3-uri", "", "The location of the file containing the template body (`s3://bucketname/prefix`).")
	config_putConformancePackCmd.Flags().String("template-ssmdocument-details", "", "An object of type `TemplateSSMDocumentDetails`, which contains the name or the Amazon Resource Name (ARN) of the Amazon Web Services Systems Manager document (SSM document) and the version of the SSM document that is used to create a conformance pack.")
	config_putConformancePackCmd.MarkFlagRequired("conformance-pack-name")
	configCmd.AddCommand(config_putConformancePackCmd)
}
