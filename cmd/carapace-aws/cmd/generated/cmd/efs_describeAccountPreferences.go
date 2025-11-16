package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_describeAccountPreferencesCmd = &cobra.Command{
	Use:   "describe-account-preferences",
	Short: "Returns the account preferences settings for the Amazon Web Services account associated with the user making the request, in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_describeAccountPreferencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_describeAccountPreferencesCmd).Standalone()

		efs_describeAccountPreferencesCmd.Flags().String("max-results", "", "(Optional) When retrieving account preferences, you can optionally specify the `MaxItems` parameter to limit the number of objects returned in a response.")
		efs_describeAccountPreferencesCmd.Flags().String("next-token", "", "(Optional) You can use `NextToken` in a subsequent request to fetch the next page of Amazon Web Services account preferences if the response payload was paginated.")
	})
	efsCmd.AddCommand(efs_describeAccountPreferencesCmd)
}
