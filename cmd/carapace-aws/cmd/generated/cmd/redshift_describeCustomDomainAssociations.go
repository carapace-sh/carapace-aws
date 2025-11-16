package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeCustomDomainAssociationsCmd = &cobra.Command{
	Use:   "describe-custom-domain-associations",
	Short: "Contains information about custom domain associations for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeCustomDomainAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_describeCustomDomainAssociationsCmd).Standalone()

		redshift_describeCustomDomainAssociationsCmd.Flags().String("custom-domain-certificate-arn", "", "The certificate Amazon Resource Name (ARN) for the custom domain association.")
		redshift_describeCustomDomainAssociationsCmd.Flags().String("custom-domain-name", "", "The custom domain name for the custom domain association.")
		redshift_describeCustomDomainAssociationsCmd.Flags().String("marker", "", "The marker for the custom domain association.")
		redshift_describeCustomDomainAssociationsCmd.Flags().String("max-records", "", "The maximum records setting for the associated custom domain.")
	})
	redshiftCmd.AddCommand(redshift_describeCustomDomainAssociationsCmd)
}
