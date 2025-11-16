package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_listLogSubscriptionsCmd = &cobra.Command{
	Use:   "list-log-subscriptions",
	Short: "Lists the active log subscriptions for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_listLogSubscriptionsCmd).Standalone()

	ds_listLogSubscriptionsCmd.Flags().String("directory-id", "", "If a *DirectoryID* is provided, lists only the log subscription associated with that directory.")
	ds_listLogSubscriptionsCmd.Flags().String("limit", "", "The maximum number of items returned.")
	ds_listLogSubscriptionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	dsCmd.AddCommand(ds_listLogSubscriptionsCmd)
}
