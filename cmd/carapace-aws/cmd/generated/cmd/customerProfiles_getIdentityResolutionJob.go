package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getIdentityResolutionJobCmd = &cobra.Command{
	Use:   "get-identity-resolution-job",
	Short: "Returns information about an Identity Resolution Job in a specific domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getIdentityResolutionJobCmd).Standalone()

	customerProfiles_getIdentityResolutionJobCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getIdentityResolutionJobCmd.Flags().String("job-id", "", "The unique identifier of the Identity Resolution Job.")
	customerProfiles_getIdentityResolutionJobCmd.MarkFlagRequired("domain-name")
	customerProfiles_getIdentityResolutionJobCmd.MarkFlagRequired("job-id")
	customerProfilesCmd.AddCommand(customerProfiles_getIdentityResolutionJobCmd)
}
