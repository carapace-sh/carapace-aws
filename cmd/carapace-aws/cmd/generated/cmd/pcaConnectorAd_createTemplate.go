package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorAd_createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Creates an Active Directory compatible certificate template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorAd_createTemplateCmd).Standalone()

	pcaConnectorAd_createTemplateCmd.Flags().String("client-token", "", "Idempotency token.")
	pcaConnectorAd_createTemplateCmd.Flags().String("connector-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateConnector](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html).")
	pcaConnectorAd_createTemplateCmd.Flags().String("definition", "", "Template configuration to define the information included in certificates.")
	pcaConnectorAd_createTemplateCmd.Flags().String("name", "", "Name of the template.")
	pcaConnectorAd_createTemplateCmd.Flags().String("tags", "", "Metadata assigned to a template consisting of a key-value pair.")
	pcaConnectorAd_createTemplateCmd.MarkFlagRequired("connector-arn")
	pcaConnectorAd_createTemplateCmd.MarkFlagRequired("definition")
	pcaConnectorAd_createTemplateCmd.MarkFlagRequired("name")
	pcaConnectorAdCmd.AddCommand(pcaConnectorAd_createTemplateCmd)
}
