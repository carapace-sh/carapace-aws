package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listEnvironmentProfilesCmd = &cobra.Command{
	Use:   "list-environment-profiles",
	Short: "Lists Amazon DataZone environment profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listEnvironmentProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listEnvironmentProfilesCmd).Standalone()

		datazone_listEnvironmentProfilesCmd.Flags().String("aws-account-id", "", "The identifier of the Amazon Web Services account where you want to list environment profiles.")
		datazone_listEnvironmentProfilesCmd.Flags().String("aws-account-region", "", "The Amazon Web Services region where you want to list environment profiles.")
		datazone_listEnvironmentProfilesCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_listEnvironmentProfilesCmd.Flags().String("environment-blueprint-identifier", "", "The identifier of the blueprint that was used to create the environment profiles that you want to list.")
		datazone_listEnvironmentProfilesCmd.Flags().String("max-results", "", "The maximum number of environment profiles to return in a single call to `ListEnvironmentProfiles`.")
		datazone_listEnvironmentProfilesCmd.Flags().String("name", "", "")
		datazone_listEnvironmentProfilesCmd.Flags().String("next-token", "", "When the number of environment profiles is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of environment profiles, the response includes a pagination token named `NextToken`.")
		datazone_listEnvironmentProfilesCmd.Flags().String("project-identifier", "", "The identifier of the Amazon DataZone project.")
		datazone_listEnvironmentProfilesCmd.MarkFlagRequired("domain-identifier")
	})
	datazoneCmd.AddCommand(datazone_listEnvironmentProfilesCmd)
}
