package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listIdentityResolutionJobsCmd = &cobra.Command{
	Use:   "list-identity-resolution-jobs",
	Short: "Lists all of the Identity Resolution Jobs in your domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listIdentityResolutionJobsCmd).Standalone()

	customerProfiles_listIdentityResolutionJobsCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_listIdentityResolutionJobsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	customerProfiles_listIdentityResolutionJobsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	customerProfiles_listIdentityResolutionJobsCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_listIdentityResolutionJobsCmd)
}
