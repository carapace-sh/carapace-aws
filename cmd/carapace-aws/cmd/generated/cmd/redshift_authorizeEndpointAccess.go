package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_authorizeEndpointAccessCmd = &cobra.Command{
	Use:   "authorize-endpoint-access",
	Short: "Grants access to a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_authorizeEndpointAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_authorizeEndpointAccessCmd).Standalone()

		redshift_authorizeEndpointAccessCmd.Flags().String("account", "", "The Amazon Web Services account ID to grant access to.")
		redshift_authorizeEndpointAccessCmd.Flags().String("cluster-identifier", "", "The cluster identifier of the cluster to grant access to.")
		redshift_authorizeEndpointAccessCmd.Flags().String("vpc-ids", "", "The virtual private cloud (VPC) identifiers to grant access to.")
		redshift_authorizeEndpointAccessCmd.MarkFlagRequired("account")
	})
	redshiftCmd.AddCommand(redshift_authorizeEndpointAccessCmd)
}
