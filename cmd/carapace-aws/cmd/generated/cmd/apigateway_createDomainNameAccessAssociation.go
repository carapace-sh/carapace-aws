package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createDomainNameAccessAssociationCmd = &cobra.Command{
	Use:   "create-domain-name-access-association",
	Short: "Creates a domain name access association resource between an access association source and a private custom domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createDomainNameAccessAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_createDomainNameAccessAssociationCmd).Standalone()

		apigateway_createDomainNameAccessAssociationCmd.Flags().String("access-association-source", "", "The identifier of the domain name access association source.")
		apigateway_createDomainNameAccessAssociationCmd.Flags().String("access-association-source-type", "", "The type of the domain name access association source.")
		apigateway_createDomainNameAccessAssociationCmd.Flags().String("domain-name-arn", "", "The ARN of the domain name.")
		apigateway_createDomainNameAccessAssociationCmd.Flags().String("tags", "", "The key-value map of strings.")
		apigateway_createDomainNameAccessAssociationCmd.MarkFlagRequired("access-association-source")
		apigateway_createDomainNameAccessAssociationCmd.MarkFlagRequired("access-association-source-type")
		apigateway_createDomainNameAccessAssociationCmd.MarkFlagRequired("domain-name-arn")
	})
	apigatewayCmd.AddCommand(apigateway_createDomainNameAccessAssociationCmd)
}
