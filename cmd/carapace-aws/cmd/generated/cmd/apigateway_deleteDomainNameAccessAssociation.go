package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteDomainNameAccessAssociationCmd = &cobra.Command{
	Use:   "delete-domain-name-access-association",
	Short: "Deletes the DomainNameAccessAssociation resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteDomainNameAccessAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteDomainNameAccessAssociationCmd).Standalone()

		apigateway_deleteDomainNameAccessAssociationCmd.Flags().String("domain-name-access-association-arn", "", "The ARN of the domain name access association resource.")
		apigateway_deleteDomainNameAccessAssociationCmd.MarkFlagRequired("domain-name-access-association-arn")
	})
	apigatewayCmd.AddCommand(apigateway_deleteDomainNameAccessAssociationCmd)
}
