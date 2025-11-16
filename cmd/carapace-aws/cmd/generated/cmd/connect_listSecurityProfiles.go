package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listSecurityProfilesCmd = &cobra.Command{
	Use:   "list-security-profiles",
	Short: "Provides summary information about the security profiles for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listSecurityProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listSecurityProfilesCmd).Standalone()

		connect_listSecurityProfilesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listSecurityProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listSecurityProfilesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listSecurityProfilesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listSecurityProfilesCmd)
}
