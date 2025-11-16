package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_revokeDbsecurityGroupIngressCmd = &cobra.Command{
	Use:   "revoke-dbsecurity-group-ingress",
	Short: "Revokes ingress from a DBSecurityGroup for previously authorized IP ranges or EC2 or VPC security groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_revokeDbsecurityGroupIngressCmd).Standalone()

	rds_revokeDbsecurityGroupIngressCmd.Flags().String("cidrip", "", "The IP range to revoke access from.")
	rds_revokeDbsecurityGroupIngressCmd.Flags().String("dbsecurity-group-name", "", "The name of the DB security group to revoke ingress from.")
	rds_revokeDbsecurityGroupIngressCmd.Flags().String("ec2-security-group-id", "", "The id of the EC2 security group to revoke access from.")
	rds_revokeDbsecurityGroupIngressCmd.Flags().String("ec2-security-group-name", "", "The name of the EC2 security group to revoke access from.")
	rds_revokeDbsecurityGroupIngressCmd.Flags().String("ec2-security-group-owner-id", "", "The Amazon Web Services account number of the owner of the EC2 security group specified in the `EC2SecurityGroupName` parameter.")
	rds_revokeDbsecurityGroupIngressCmd.MarkFlagRequired("dbsecurity-group-name")
	rdsCmd.AddCommand(rds_revokeDbsecurityGroupIngressCmd)
}
