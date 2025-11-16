package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listSuppressedDestinationsCmd = &cobra.Command{
	Use:   "list-suppressed-destinations",
	Short: "Retrieves a list of email addresses that are on the suppression list for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listSuppressedDestinationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listSuppressedDestinationsCmd).Standalone()

		sesv2_listSuppressedDestinationsCmd.Flags().String("end-date", "", "Used to filter the list of suppressed email destinations so that it only includes addresses that were added to the list before a specific date.")
		sesv2_listSuppressedDestinationsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListSuppressedDestinations` to indicate the position in the list of suppressed email addresses.")
		sesv2_listSuppressedDestinationsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListSuppressedDestinations`.")
		sesv2_listSuppressedDestinationsCmd.Flags().String("reasons", "", "The factors that caused the email address to be added to .")
		sesv2_listSuppressedDestinationsCmd.Flags().String("start-date", "", "Used to filter the list of suppressed email destinations so that it only includes addresses that were added to the list after a specific date.")
	})
	sesv2Cmd.AddCommand(sesv2_listSuppressedDestinationsCmd)
}
