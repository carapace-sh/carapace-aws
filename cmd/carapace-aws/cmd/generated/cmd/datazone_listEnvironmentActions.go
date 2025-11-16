package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listEnvironmentActionsCmd = &cobra.Command{
	Use:   "list-environment-actions",
	Short: "Lists existing environment actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listEnvironmentActionsCmd).Standalone()

	datazone_listEnvironmentActionsCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the environment actions are listed.")
	datazone_listEnvironmentActionsCmd.Flags().String("environment-identifier", "", "The ID of the envrironment whose environment actions are listed.")
	datazone_listEnvironmentActionsCmd.Flags().String("max-results", "", "The maximum number of environment actions to return in a single call to `ListEnvironmentActions`.")
	datazone_listEnvironmentActionsCmd.Flags().String("next-token", "", "When the number of environment actions is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of environment actions, the response includes a pagination token named `NextToken`.")
	datazone_listEnvironmentActionsCmd.MarkFlagRequired("domain-identifier")
	datazone_listEnvironmentActionsCmd.MarkFlagRequired("environment-identifier")
	datazoneCmd.AddCommand(datazone_listEnvironmentActionsCmd)
}
