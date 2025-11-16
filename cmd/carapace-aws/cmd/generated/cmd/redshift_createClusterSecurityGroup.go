package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createClusterSecurityGroupCmd = &cobra.Command{
	Use:   "create-cluster-security-group",
	Short: "Creates a new Amazon Redshift security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createClusterSecurityGroupCmd).Standalone()

	redshift_createClusterSecurityGroupCmd.Flags().String("cluster-security-group-name", "", "The name for the security group.")
	redshift_createClusterSecurityGroupCmd.Flags().String("description", "", "A description for the security group.")
	redshift_createClusterSecurityGroupCmd.Flags().String("tags", "", "A list of tag instances.")
	redshift_createClusterSecurityGroupCmd.MarkFlagRequired("cluster-security-group-name")
	redshift_createClusterSecurityGroupCmd.MarkFlagRequired("description")
	redshiftCmd.AddCommand(redshift_createClusterSecurityGroupCmd)
}
