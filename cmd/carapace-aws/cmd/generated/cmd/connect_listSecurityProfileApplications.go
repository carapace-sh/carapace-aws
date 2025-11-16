package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listSecurityProfileApplicationsCmd = &cobra.Command{
	Use:   "list-security-profile-applications",
	Short: "Returns a list of third-party applications in a specific security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listSecurityProfileApplicationsCmd).Standalone()

	connect_listSecurityProfileApplicationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listSecurityProfileApplicationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listSecurityProfileApplicationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listSecurityProfileApplicationsCmd.Flags().String("security-profile-id", "", "The identifier for the security profle.")
	connect_listSecurityProfileApplicationsCmd.MarkFlagRequired("instance-id")
	connect_listSecurityProfileApplicationsCmd.MarkFlagRequired("security-profile-id")
	connectCmd.AddCommand(connect_listSecurityProfileApplicationsCmd)
}
