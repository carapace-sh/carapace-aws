package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listSecurityProfilePermissionsCmd = &cobra.Command{
	Use:   "list-security-profile-permissions",
	Short: "Lists the permissions granted to a security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listSecurityProfilePermissionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listSecurityProfilePermissionsCmd).Standalone()

		connect_listSecurityProfilePermissionsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listSecurityProfilePermissionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listSecurityProfilePermissionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listSecurityProfilePermissionsCmd.Flags().String("security-profile-id", "", "The identifier for the security profle.")
		connect_listSecurityProfilePermissionsCmd.MarkFlagRequired("instance-id")
		connect_listSecurityProfilePermissionsCmd.MarkFlagRequired("security-profile-id")
	})
	connectCmd.AddCommand(connect_listSecurityProfilePermissionsCmd)
}
