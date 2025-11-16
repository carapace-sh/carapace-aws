package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getCustomDomainAssociationCmd = &cobra.Command{
	Use:   "get-custom-domain-association",
	Short: "Gets information about a specific custom domain association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getCustomDomainAssociationCmd).Standalone()

	redshiftServerless_getCustomDomainAssociationCmd.Flags().String("custom-domain-name", "", "The custom domain name associated with the workgroup.")
	redshiftServerless_getCustomDomainAssociationCmd.Flags().String("workgroup-name", "", "The name of the workgroup associated with the database.")
	redshiftServerless_getCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-name")
	redshiftServerless_getCustomDomainAssociationCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getCustomDomainAssociationCmd)
}
