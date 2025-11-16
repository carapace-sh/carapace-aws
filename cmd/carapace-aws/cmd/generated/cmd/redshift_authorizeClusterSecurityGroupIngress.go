package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_authorizeClusterSecurityGroupIngressCmd = &cobra.Command{
	Use:   "authorize-cluster-security-group-ingress",
	Short: "Adds an inbound (ingress) rule to an Amazon Redshift security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_authorizeClusterSecurityGroupIngressCmd).Standalone()

	redshift_authorizeClusterSecurityGroupIngressCmd.Flags().String("cidrip", "", "The IP range to be added the Amazon Redshift security group.")
	redshift_authorizeClusterSecurityGroupIngressCmd.Flags().String("cluster-security-group-name", "", "The name of the security group to which the ingress rule is added.")
	redshift_authorizeClusterSecurityGroupIngressCmd.Flags().String("ec2-security-group-name", "", "The EC2 security group to be added the Amazon Redshift security group.")
	redshift_authorizeClusterSecurityGroupIngressCmd.Flags().String("ec2-security-group-owner-id", "", "The Amazon Web Services account number of the owner of the security group specified by the *EC2SecurityGroupName* parameter.")
	redshift_authorizeClusterSecurityGroupIngressCmd.MarkFlagRequired("cluster-security-group-name")
	redshiftCmd.AddCommand(redshift_authorizeClusterSecurityGroupIngressCmd)
}
