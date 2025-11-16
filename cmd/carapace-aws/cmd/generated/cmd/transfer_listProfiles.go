package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_listProfilesCmd = &cobra.Command{
	Use:   "list-profiles",
	Short: "Returns a list of the profiles for your system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_listProfilesCmd).Standalone()

	transfer_listProfilesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	transfer_listProfilesCmd.Flags().String("next-token", "", "When there are additional results that were not returned, a `NextToken` parameter is returned.")
	transfer_listProfilesCmd.Flags().String("profile-type", "", "Indicates whether to list only `LOCAL` type profiles or only `PARTNER` type profiles.")
	transferCmd.AddCommand(transfer_listProfilesCmd)
}
