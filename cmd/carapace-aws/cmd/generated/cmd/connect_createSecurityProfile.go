package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createSecurityProfileCmd = &cobra.Command{
	Use:   "create-security-profile",
	Short: "Creates a security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createSecurityProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createSecurityProfileCmd).Standalone()

		connect_createSecurityProfileCmd.Flags().String("allowed-access-control-hierarchy-group-id", "", "The identifier of the hierarchy group that a security profile uses to restrict access to resources in Amazon Connect.")
		connect_createSecurityProfileCmd.Flags().String("allowed-access-control-tags", "", "The list of tags that a security profile uses to restrict access to resources in Amazon Connect.")
		connect_createSecurityProfileCmd.Flags().String("applications", "", "A list of third-party applications that the security profile will give access to.")
		connect_createSecurityProfileCmd.Flags().String("description", "", "The description of the security profile.")
		connect_createSecurityProfileCmd.Flags().String("hierarchy-restricted-resources", "", "The list of resources that a security profile applies hierarchy restrictions to in Amazon Connect.")
		connect_createSecurityProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createSecurityProfileCmd.Flags().String("permissions", "", "Permissions assigned to the security profile.")
		connect_createSecurityProfileCmd.Flags().String("security-profile-name", "", "The name of the security profile.")
		connect_createSecurityProfileCmd.Flags().String("tag-restricted-resources", "", "The list of resources that a security profile applies tag restrictions to in Amazon Connect.")
		connect_createSecurityProfileCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createSecurityProfileCmd.MarkFlagRequired("instance-id")
		connect_createSecurityProfileCmd.MarkFlagRequired("security-profile-name")
	})
	connectCmd.AddCommand(connect_createSecurityProfileCmd)
}
