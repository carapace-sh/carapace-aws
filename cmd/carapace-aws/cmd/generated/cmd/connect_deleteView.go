package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteViewCmd = &cobra.Command{
	Use:   "delete-view",
	Short: "Deletes the view entirely.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteViewCmd).Standalone()

	connect_deleteViewCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteViewCmd.Flags().String("view-id", "", "The identifier of the view.")
	connect_deleteViewCmd.MarkFlagRequired("instance-id")
	connect_deleteViewCmd.MarkFlagRequired("view-id")
	connectCmd.AddCommand(connect_deleteViewCmd)
}
