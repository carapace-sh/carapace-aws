package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteCustomDomainAssociationCmd = &cobra.Command{
	Use:   "delete-custom-domain-association",
	Short: "Deletes a custom domain association for Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteCustomDomainAssociationCmd).Standalone()

	redshiftServerless_deleteCustomDomainAssociationCmd.Flags().String("custom-domain-name", "", "The custom domain name associated with the workgroup.")
	redshiftServerless_deleteCustomDomainAssociationCmd.Flags().String("workgroup-name", "", "The name of the workgroup associated with the database.")
	redshiftServerless_deleteCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-name")
	redshiftServerless_deleteCustomDomainAssociationCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteCustomDomainAssociationCmd)
}
