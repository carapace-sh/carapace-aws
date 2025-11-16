package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listEnvironmentBlueprintConfigurationsCmd = &cobra.Command{
	Use:   "list-environment-blueprint-configurations",
	Short: "Lists blueprint configurations for a Amazon DataZone environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listEnvironmentBlueprintConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listEnvironmentBlueprintConfigurationsCmd).Standalone()

		datazone_listEnvironmentBlueprintConfigurationsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_listEnvironmentBlueprintConfigurationsCmd.Flags().String("max-results", "", "The maximum number of blueprint configurations to return in a single call to `ListEnvironmentBlueprintConfigurations`.")
		datazone_listEnvironmentBlueprintConfigurationsCmd.Flags().String("next-token", "", "When the number of blueprint configurations is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of configurations, the response includes a pagination token named `NextToken`.")
		datazone_listEnvironmentBlueprintConfigurationsCmd.MarkFlagRequired("domain-identifier")
	})
	datazoneCmd.AddCommand(datazone_listEnvironmentBlueprintConfigurationsCmd)
}
