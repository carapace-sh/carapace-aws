package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteViewVersionCmd = &cobra.Command{
	Use:   "delete-view-version",
	Short: "Deletes the particular version specified in `ViewVersion` identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteViewVersionCmd).Standalone()

	connect_deleteViewVersionCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteViewVersionCmd.Flags().String("view-id", "", "The identifier of the view.")
	connect_deleteViewVersionCmd.Flags().String("view-version", "", "The version number of the view.")
	connect_deleteViewVersionCmd.MarkFlagRequired("instance-id")
	connect_deleteViewVersionCmd.MarkFlagRequired("view-id")
	connect_deleteViewVersionCmd.MarkFlagRequired("view-version")
	connectCmd.AddCommand(connect_deleteViewVersionCmd)
}
