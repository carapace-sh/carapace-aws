package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createCustomDomainAssociationCmd = &cobra.Command{
	Use:   "create-custom-domain-association",
	Short: "Used to create a custom domain name for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createCustomDomainAssociationCmd).Standalone()

	redshift_createCustomDomainAssociationCmd.Flags().String("cluster-identifier", "", "The cluster identifier that the custom domain is associated with.")
	redshift_createCustomDomainAssociationCmd.Flags().String("custom-domain-certificate-arn", "", "The certificate Amazon Resource Name (ARN) for the custom domain name association.")
	redshift_createCustomDomainAssociationCmd.Flags().String("custom-domain-name", "", "The custom domain name for a custom domain association.")
	redshift_createCustomDomainAssociationCmd.MarkFlagRequired("cluster-identifier")
	redshift_createCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-certificate-arn")
	redshift_createCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-name")
	redshiftCmd.AddCommand(redshift_createCustomDomainAssociationCmd)
}
