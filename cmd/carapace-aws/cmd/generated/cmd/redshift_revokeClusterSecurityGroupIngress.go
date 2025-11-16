package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_revokeClusterSecurityGroupIngressCmd = &cobra.Command{
	Use:   "revoke-cluster-security-group-ingress",
	Short: "Revokes an ingress rule in an Amazon Redshift security group for a previously authorized IP range or Amazon EC2 security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_revokeClusterSecurityGroupIngressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_revokeClusterSecurityGroupIngressCmd).Standalone()

		redshift_revokeClusterSecurityGroupIngressCmd.Flags().String("cidrip", "", "The IP range for which to revoke access.")
		redshift_revokeClusterSecurityGroupIngressCmd.Flags().String("cluster-security-group-name", "", "The name of the security Group from which to revoke the ingress rule.")
		redshift_revokeClusterSecurityGroupIngressCmd.Flags().String("ec2-security-group-name", "", "The name of the EC2 Security Group whose access is to be revoked.")
		redshift_revokeClusterSecurityGroupIngressCmd.Flags().String("ec2-security-group-owner-id", "", "The Amazon Web Services account number of the owner of the security group specified in the `EC2SecurityGroupName` parameter.")
		redshift_revokeClusterSecurityGroupIngressCmd.MarkFlagRequired("cluster-security-group-name")
	})
	redshiftCmd.AddCommand(redshift_revokeClusterSecurityGroupIngressCmd)
}
