package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteCustomDomainAssociationCmd = &cobra.Command{
	Use:   "delete-custom-domain-association",
	Short: "Contains information about deleting a custom domain association for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteCustomDomainAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_deleteCustomDomainAssociationCmd).Standalone()

		redshift_deleteCustomDomainAssociationCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster to delete a custom domain association for.")
		redshift_deleteCustomDomainAssociationCmd.Flags().String("custom-domain-name", "", "The custom domain name for the custom domain association.")
		redshift_deleteCustomDomainAssociationCmd.MarkFlagRequired("cluster-identifier")
		redshift_deleteCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-name")
	})
	redshiftCmd.AddCommand(redshift_deleteCustomDomainAssociationCmd)
}
