package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createWebAclmigrationStackCmd = &cobra.Command{
	Use:   "create-web-aclmigration-stack",
	Short: "Creates an AWS CloudFormation WAFV2 template for the specified web ACL in the specified Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createWebAclmigrationStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createWebAclmigrationStackCmd).Standalone()

		wafRegional_createWebAclmigrationStackCmd.Flags().String("ignore-unsupported-type", "", "Indicates whether to exclude entities that can't be migrated or to stop the migration.")
		wafRegional_createWebAclmigrationStackCmd.Flags().String("s3-bucket-name", "", "The name of the Amazon S3 bucket to store the CloudFormation template in.")
		wafRegional_createWebAclmigrationStackCmd.Flags().String("web-aclid", "", "The UUID of the WAF Classic web ACL that you want to migrate to WAF v2.")
		wafRegional_createWebAclmigrationStackCmd.MarkFlagRequired("ignore-unsupported-type")
		wafRegional_createWebAclmigrationStackCmd.MarkFlagRequired("s3-bucket-name")
		wafRegional_createWebAclmigrationStackCmd.MarkFlagRequired("web-aclid")
	})
	wafRegionalCmd.AddCommand(wafRegional_createWebAclmigrationStackCmd)
}
