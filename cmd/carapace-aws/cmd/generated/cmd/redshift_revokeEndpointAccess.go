package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_revokeEndpointAccessCmd = &cobra.Command{
	Use:   "revoke-endpoint-access",
	Short: "Revokes access to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_revokeEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_revokeEndpointAccessCmd).Standalone()

		redshift_revokeEndpointAccessCmd.Flags().String("account", "", "The Amazon Web Services account ID whose access is to be revoked.")
		redshift_revokeEndpointAccessCmd.Flags().String("cluster-identifier", "", "The cluster to revoke access from.")
		redshift_revokeEndpointAccessCmd.Flags().Bool("force", false, "Indicates whether to force the revoke action.")
		redshift_revokeEndpointAccessCmd.Flags().Bool("no-force", false, "Indicates whether to force the revoke action.")
		redshift_revokeEndpointAccessCmd.Flags().String("vpc-ids", "", "The virtual private cloud (VPC) identifiers for which access is to be revoked.")
		redshift_revokeEndpointAccessCmd.Flag("no-force").Hidden = true
	})
	redshiftCmd.AddCommand(redshift_revokeEndpointAccessCmd)
}
