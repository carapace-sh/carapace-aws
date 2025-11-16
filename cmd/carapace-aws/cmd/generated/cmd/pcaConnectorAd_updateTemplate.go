package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_updateTemplateCmd = &cobra.Command{
	Use:   "update-template",
	Short: "Update template configuration to define the information included in certificates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_updateTemplateCmd).Standalone()

	pcaConnectorAd_updateTemplateCmd.Flags().String("definition", "", "Template configuration to define the information included in certificates.")
	pcaConnectorAd_updateTemplateCmd.Flags().Bool("no-reenroll-all-certificate-holders", false, "This setting allows the major version of a template to be increased automatically.")
	pcaConnectorAd_updateTemplateCmd.Flags().Bool("reenroll-all-certificate-holders", false, "This setting allows the major version of a template to be increased automatically.")
	pcaConnectorAd_updateTemplateCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
	pcaConnectorAd_updateTemplateCmd.Flag("no-reenroll-all-certificate-holders").Hidden = true
	pcaConnectorAd_updateTemplateCmd.MarkFlagRequired("template-arn")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_updateTemplateCmd)
}
