package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_rejectDomainNameAccessAssociationCmd = &cobra.Command{
	Use:   "reject-domain-name-access-association",
	Short: "Rejects a domain name access association with a private custom domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_rejectDomainNameAccessAssociationCmd).Standalone()

	apigateway_rejectDomainNameAccessAssociationCmd.Flags().String("domain-name-access-association-arn", "", "The ARN of the domain name access association resource.")
	apigateway_rejectDomainNameAccessAssociationCmd.Flags().String("domain-name-arn", "", "The ARN of the domain name.")
	apigateway_rejectDomainNameAccessAssociationCmd.MarkFlagRequired("domain-name-access-association-arn")
	apigateway_rejectDomainNameAccessAssociationCmd.MarkFlagRequired("domain-name-arn")
	apigatewayCmd.AddCommand(apigateway_rejectDomainNameAccessAssociationCmd)
}
