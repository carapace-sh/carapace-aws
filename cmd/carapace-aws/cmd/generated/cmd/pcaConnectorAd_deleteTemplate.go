package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_deleteTemplateCmd = &cobra.Command{
	Use:   "delete-template",
	Short: "Deletes a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_deleteTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pcaConnectorAd_deleteTemplateCmd).Standalone()

		pcaConnectorAd_deleteTemplateCmd.Flags().String("template-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateTemplate](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateTemplate.html).")
		pcaConnectorAd_deleteTemplateCmd.MarkFlagRequired("template-arn")
	})
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_deleteTemplateCmd)
}
