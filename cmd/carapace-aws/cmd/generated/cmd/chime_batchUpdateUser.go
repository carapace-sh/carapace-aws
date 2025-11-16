package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_batchUpdateUserCmd = &cobra.Command{
	Use:   "batch-update-user",
	Short: "Updates user details within the [UpdateUserRequestItem]() object for up to 20 users for the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_batchUpdateUserCmd).Standalone()

	chime_batchUpdateUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_batchUpdateUserCmd.Flags().String("update-user-request-items", "", "The request containing the user IDs and details to update.")
	chime_batchUpdateUserCmd.MarkFlagRequired("account-id")
	chime_batchUpdateUserCmd.MarkFlagRequired("update-user-request-items")
	chimeCmd.AddCommand(chime_batchUpdateUserCmd)
}
