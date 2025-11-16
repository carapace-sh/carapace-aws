package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_listPickupLocationsCmd = &cobra.Command{
	Use:   "list-pickup-locations",
	Short: "A list of locations from which the customer can choose to pickup a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_listPickupLocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_listPickupLocationsCmd).Standalone()

		snowball_listPickupLocationsCmd.Flags().String("max-results", "", "The maximum number of locations to list per page.")
		snowball_listPickupLocationsCmd.Flags().String("next-token", "", "HTTP requests are stateless.")
	})
	snowballCmd.AddCommand(snowball_listPickupLocationsCmd)
}
