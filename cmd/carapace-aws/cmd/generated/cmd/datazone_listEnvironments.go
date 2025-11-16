package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Lists Amazon DataZone environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listEnvironmentsCmd).Standalone()

		datazone_listEnvironmentsCmd.Flags().String("aws-account-id", "", "The identifier of the Amazon Web Services account where you want to list environments.")
		datazone_listEnvironmentsCmd.Flags().String("aws-account-region", "", "The Amazon Web Services region where you want to list environments.")
		datazone_listEnvironmentsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_listEnvironmentsCmd.Flags().String("environment-blueprint-identifier", "", "The identifier of the Amazon DataZone blueprint.")
		datazone_listEnvironmentsCmd.Flags().String("environment-profile-identifier", "", "The identifier of the environment profile.")
		datazone_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of environments to return in a single call to `ListEnvironments`.")
		datazone_listEnvironmentsCmd.Flags().String("name", "", "The name of the environment.")
		datazone_listEnvironmentsCmd.Flags().String("next-token", "", "When the number of environments is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of environments, the response includes a pagination token named `NextToken`.")
		datazone_listEnvironmentsCmd.Flags().String("project-identifier", "", "The identifier of the Amazon DataZone project.")
		datazone_listEnvironmentsCmd.Flags().String("provider", "", "The provider of the environment.")
		datazone_listEnvironmentsCmd.Flags().String("status", "", "The status of the environments that you want to list.")
		datazone_listEnvironmentsCmd.MarkFlagRequired("domain-identifier")
		datazone_listEnvironmentsCmd.MarkFlagRequired("project-identifier")
	})
	datazoneCmd.AddCommand(datazone_listEnvironmentsCmd)
}
