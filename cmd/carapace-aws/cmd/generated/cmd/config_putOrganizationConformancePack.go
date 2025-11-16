package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putOrganizationConformancePackCmd = &cobra.Command{
	Use:   "put-organization-conformance-pack",
	Short: "Deploys conformance packs across member accounts in an Amazon Web Services Organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putOrganizationConformancePackCmd).Standalone()

	config_putOrganizationConformancePackCmd.Flags().String("conformance-pack-input-parameters", "", "A list of `ConformancePackInputParameter` objects.")
	config_putOrganizationConformancePackCmd.Flags().String("delivery-s3-bucket", "", "The name of the Amazon S3 bucket where Config stores conformance pack templates.")
	config_putOrganizationConformancePackCmd.Flags().String("delivery-s3-key-prefix", "", "The prefix for the Amazon S3 bucket.")
	config_putOrganizationConformancePackCmd.Flags().String("excluded-accounts", "", "A list of Amazon Web Services accounts to be excluded from an organization conformance pack while deploying a conformance pack.")
	config_putOrganizationConformancePackCmd.Flags().String("organization-conformance-pack-name", "", "Name of the organization conformance pack you want to create.")
	config_putOrganizationConformancePackCmd.Flags().String("template-body", "", "A string that contains the full conformance pack template body.")
	config_putOrganizationConformancePackCmd.Flags().String("template-s3-uri", "", "Location of file containing the template body.")
	config_putOrganizationConformancePackCmd.MarkFlagRequired("organization-conformance-pack-name")
	configCmd.AddCommand(config_putOrganizationConformancePackCmd)
}
