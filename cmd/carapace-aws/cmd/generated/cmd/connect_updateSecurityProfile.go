package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateSecurityProfileCmd = &cobra.Command{
	Use:   "update-security-profile",
	Short: "Updates a security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateSecurityProfileCmd).Standalone()

	connect_updateSecurityProfileCmd.Flags().String("allowed-access-control-hierarchy-group-id", "", "The identifier of the hierarchy group that a security profile uses to restrict access to resources in Amazon Connect.")
	connect_updateSecurityProfileCmd.Flags().String("allowed-access-control-tags", "", "The list of tags that a security profile uses to restrict access to resources in Amazon Connect.")
	connect_updateSecurityProfileCmd.Flags().String("applications", "", "A list of the third-party application's metadata.")
	connect_updateSecurityProfileCmd.Flags().String("description", "", "The description of the security profile.")
	connect_updateSecurityProfileCmd.Flags().String("hierarchy-restricted-resources", "", "The list of resources that a security profile applies hierarchy restrictions to in Amazon Connect.")
	connect_updateSecurityProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateSecurityProfileCmd.Flags().String("permissions", "", "The permissions granted to a security profile.")
	connect_updateSecurityProfileCmd.Flags().String("security-profile-id", "", "The identifier for the security profle.")
	connect_updateSecurityProfileCmd.Flags().String("tag-restricted-resources", "", "The list of resources that a security profile applies tag restrictions to in Amazon Connect.")
	connect_updateSecurityProfileCmd.MarkFlagRequired("instance-id")
	connect_updateSecurityProfileCmd.MarkFlagRequired("security-profile-id")
	connectCmd.AddCommand(connect_updateSecurityProfileCmd)
}
