package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getBlueprintsCmd = &cobra.Command{
	Use:   "get-blueprints",
	Short: "Returns the list of available instance images, or *blueprints*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getBlueprintsCmd).Standalone()

	lightsail_getBlueprintsCmd.Flags().String("app-category", "", "Returns a list of blueprints that are specific to Lightsail for Research.")
	lightsail_getBlueprintsCmd.Flags().String("include-inactive", "", "A Boolean value that indicates whether to include inactive (unavailable) blueprints in the response of your request.")
	lightsail_getBlueprintsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getBlueprintsCmd)
}
