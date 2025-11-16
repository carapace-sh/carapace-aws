package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_authorizeDbsecurityGroupIngressCmd = &cobra.Command{
	Use:   "authorize-dbsecurity-group-ingress",
	Short: "Enables ingress to a DBSecurityGroup using one of two forms of authorization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_authorizeDbsecurityGroupIngressCmd).Standalone()

	rds_authorizeDbsecurityGroupIngressCmd.Flags().String("cidrip", "", "The IP range to authorize.")
	rds_authorizeDbsecurityGroupIngressCmd.Flags().String("dbsecurity-group-name", "", "The name of the DB security group to add authorization to.")
	rds_authorizeDbsecurityGroupIngressCmd.Flags().String("ec2-security-group-id", "", "Id of the EC2 security group to authorize.")
	rds_authorizeDbsecurityGroupIngressCmd.Flags().String("ec2-security-group-name", "", "Name of the EC2 security group to authorize.")
	rds_authorizeDbsecurityGroupIngressCmd.Flags().String("ec2-security-group-owner-id", "", "Amazon Web Services account number of the owner of the EC2 security group specified in the `EC2SecurityGroupName` parameter.")
	rds_authorizeDbsecurityGroupIngressCmd.MarkFlagRequired("dbsecurity-group-name")
	rdsCmd.AddCommand(rds_authorizeDbsecurityGroupIngressCmd)
}
