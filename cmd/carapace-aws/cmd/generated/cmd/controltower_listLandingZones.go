package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_listLandingZonesCmd = &cobra.Command{
	Use:   "list-landing-zones",
	Short: "Returns the landing zone ARN for the landing zone deployed in your managed account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_listLandingZonesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_listLandingZonesCmd).Standalone()

		controltower_listLandingZonesCmd.Flags().String("max-results", "", "The maximum number of returned landing zone ARNs, which is one.")
		controltower_listLandingZonesCmd.Flags().String("next-token", "", "The token to continue the list from a previous API call with the same parameters.")
	})
	controltowerCmd.AddCommand(controltower_listLandingZonesCmd)
}
