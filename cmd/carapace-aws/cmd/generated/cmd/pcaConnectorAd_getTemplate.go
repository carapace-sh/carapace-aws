package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_getTemplateCmd = &cobra.Command{
	Use:   "get-template",
	Short: "Retrieves a certificate template that the connector uses to issue certificates from a private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_getTemplateCmd).Standalone()

	pcaConnectorAd_getTemplateCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
	pcaConnectorAd_getTemplateCmd.MarkFlagRequired("template-arn")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_getTemplateCmd)
}
