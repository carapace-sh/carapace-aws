package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteClusterSecurityGroupCmd = &cobra.Command{
	Use:   "delete-cluster-security-group",
	Short: "Deletes an Amazon Redshift security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteClusterSecurityGroupCmd).Standalone()

	redshift_deleteClusterSecurityGroupCmd.Flags().String("cluster-security-group-name", "", "The name of the cluster security group to be deleted.")
	redshift_deleteClusterSecurityGroupCmd.MarkFlagRequired("cluster-security-group-name")
	redshiftCmd.AddCommand(redshift_deleteClusterSecurityGroupCmd)
}
