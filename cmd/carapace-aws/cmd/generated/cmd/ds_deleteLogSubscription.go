package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deleteLogSubscriptionCmd = &cobra.Command{
	Use:   "delete-log-subscription",
	Short: "Deletes the specified log subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deleteLogSubscriptionCmd).Standalone()

	ds_deleteLogSubscriptionCmd.Flags().String("directory-id", "", "Identifier of the directory whose log subscription you want to delete.")
	ds_deleteLogSubscriptionCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_deleteLogSubscriptionCmd)
}
