package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createCustomDomainAssociationCmd = &cobra.Command{
	Use:   "create-custom-domain-association",
	Short: "Creates a custom domain association for Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createCustomDomainAssociationCmd).Standalone()

	redshiftServerless_createCustomDomainAssociationCmd.Flags().String("custom-domain-certificate-arn", "", "The custom domain name’s certificate Amazon resource name (ARN).")
	redshiftServerless_createCustomDomainAssociationCmd.Flags().String("custom-domain-name", "", "The custom domain name to associate with the workgroup.")
	redshiftServerless_createCustomDomainAssociationCmd.Flags().String("workgroup-name", "", "The name of the workgroup associated with the database.")
	redshiftServerless_createCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-certificate-arn")
	redshiftServerless_createCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-name")
	redshiftServerless_createCustomDomainAssociationCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_createCustomDomainAssociationCmd)
}
