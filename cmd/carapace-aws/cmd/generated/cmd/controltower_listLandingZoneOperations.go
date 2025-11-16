package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_listLandingZoneOperationsCmd = &cobra.Command{
	Use:   "list-landing-zone-operations",
	Short: "Lists all landing zone operations from the past 90 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_listLandingZoneOperationsCmd).Standalone()

	controltower_listLandingZoneOperationsCmd.Flags().String("filter", "", "An input filter for the `ListLandingZoneOperations` API that lets you select the types of landing zone operations to view.")
	controltower_listLandingZoneOperationsCmd.Flags().String("max-results", "", "How many results to return per API call.")
	controltower_listLandingZoneOperationsCmd.Flags().String("next-token", "", "The token to continue the list from a previous API call with the same parameters.")
	controltowerCmd.AddCommand(controltower_listLandingZoneOperationsCmd)
}
