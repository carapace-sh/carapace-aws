package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_revokeCacheSecurityGroupIngressCmd = &cobra.Command{
	Use:   "revoke-cache-security-group-ingress",
	Short: "Revokes ingress from a cache security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_revokeCacheSecurityGroupIngressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_revokeCacheSecurityGroupIngressCmd).Standalone()

		elasticache_revokeCacheSecurityGroupIngressCmd.Flags().String("cache-security-group-name", "", "The name of the cache security group to revoke ingress from.")
		elasticache_revokeCacheSecurityGroupIngressCmd.Flags().String("ec2-security-group-name", "", "The name of the Amazon EC2 security group to revoke access from.")
		elasticache_revokeCacheSecurityGroupIngressCmd.Flags().String("ec2-security-group-owner-id", "", "The Amazon account number of the Amazon EC2 security group owner.")
		elasticache_revokeCacheSecurityGroupIngressCmd.MarkFlagRequired("cache-security-group-name")
		elasticache_revokeCacheSecurityGroupIngressCmd.MarkFlagRequired("ec2-security-group-name")
		elasticache_revokeCacheSecurityGroupIngressCmd.MarkFlagRequired("ec2-security-group-owner-id")
	})
	elasticacheCmd.AddCommand(elasticache_revokeCacheSecurityGroupIngressCmd)
}
