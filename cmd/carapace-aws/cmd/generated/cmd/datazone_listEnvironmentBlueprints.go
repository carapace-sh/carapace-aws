package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listEnvironmentBlueprintsCmd = &cobra.Command{
	Use:   "list-environment-blueprints",
	Short: "Lists blueprints in an Amazon DataZone environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listEnvironmentBlueprintsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listEnvironmentBlueprintsCmd).Standalone()

		datazone_listEnvironmentBlueprintsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_listEnvironmentBlueprintsCmd.Flags().Bool("managed", false, "Specifies whether the environment blueprint is managed by Amazon DataZone.")
		datazone_listEnvironmentBlueprintsCmd.Flags().String("max-results", "", "The maximum number of blueprints to return in a single call to `ListEnvironmentBlueprints`.")
		datazone_listEnvironmentBlueprintsCmd.Flags().String("name", "", "The name of the Amazon DataZone environment.")
		datazone_listEnvironmentBlueprintsCmd.Flags().String("next-token", "", "When the number of blueprints in the environment is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of blueprints in the environment, the response includes a pagination token named `NextToken`.")
		datazone_listEnvironmentBlueprintsCmd.Flags().Bool("no-managed", false, "Specifies whether the environment blueprint is managed by Amazon DataZone.")
		datazone_listEnvironmentBlueprintsCmd.MarkFlagRequired("domain-identifier")
		datazone_listEnvironmentBlueprintsCmd.Flag("no-managed").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_listEnvironmentBlueprintsCmd)
}
