package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_authorizeCacheSecurityGroupIngressCmd = &cobra.Command{
	Use:   "authorize-cache-security-group-ingress",
	Short: "Allows network ingress to a cache security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_authorizeCacheSecurityGroupIngressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_authorizeCacheSecurityGroupIngressCmd).Standalone()

		elasticache_authorizeCacheSecurityGroupIngressCmd.Flags().String("cache-security-group-name", "", "The cache security group that allows network ingress.")
		elasticache_authorizeCacheSecurityGroupIngressCmd.Flags().String("ec2-security-group-name", "", "The Amazon EC2 security group to be authorized for ingress to the cache security group.")
		elasticache_authorizeCacheSecurityGroupIngressCmd.Flags().String("ec2-security-group-owner-id", "", "The Amazon account number of the Amazon EC2 security group owner.")
		elasticache_authorizeCacheSecurityGroupIngressCmd.MarkFlagRequired("cache-security-group-name")
		elasticache_authorizeCacheSecurityGroupIngressCmd.MarkFlagRequired("ec2-security-group-name")
		elasticache_authorizeCacheSecurityGroupIngressCmd.MarkFlagRequired("ec2-security-group-owner-id")
	})
	elasticacheCmd.AddCommand(elasticache_authorizeCacheSecurityGroupIngressCmd)
}
