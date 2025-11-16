package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateCustomDomainAssociationCmd = &cobra.Command{
	Use:   "update-custom-domain-association",
	Short: "Updates an Amazon Redshift Serverless certificate associated with a custom domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateCustomDomainAssociationCmd).Standalone()

	redshiftServerless_updateCustomDomainAssociationCmd.Flags().String("custom-domain-certificate-arn", "", "The custom domain name’s certificate Amazon resource name (ARN).")
	redshiftServerless_updateCustomDomainAssociationCmd.Flags().String("custom-domain-name", "", "The custom domain name associated with the workgroup.")
	redshiftServerless_updateCustomDomainAssociationCmd.Flags().String("workgroup-name", "", "The name of the workgroup associated with the database.")
	redshiftServerless_updateCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-certificate-arn")
	redshiftServerless_updateCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-name")
	redshiftServerless_updateCustomDomainAssociationCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateCustomDomainAssociationCmd)
}
