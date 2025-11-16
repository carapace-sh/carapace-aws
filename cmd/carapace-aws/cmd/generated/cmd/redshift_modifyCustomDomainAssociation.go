package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyCustomDomainAssociationCmd = &cobra.Command{
	Use:   "modify-custom-domain-association",
	Short: "Contains information for changing a custom domain association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyCustomDomainAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_modifyCustomDomainAssociationCmd).Standalone()

		redshift_modifyCustomDomainAssociationCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster to change a custom domain association for.")
		redshift_modifyCustomDomainAssociationCmd.Flags().String("custom-domain-certificate-arn", "", "The certificate Amazon Resource Name (ARN) for the changed custom domain association.")
		redshift_modifyCustomDomainAssociationCmd.Flags().String("custom-domain-name", "", "The custom domain name for a changed custom domain association.")
		redshift_modifyCustomDomainAssociationCmd.MarkFlagRequired("cluster-identifier")
		redshift_modifyCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-certificate-arn")
		redshift_modifyCustomDomainAssociationCmd.MarkFlagRequired("custom-domain-name")
	})
	redshiftCmd.AddCommand(redshift_modifyCustomDomainAssociationCmd)
}
