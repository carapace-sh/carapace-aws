package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteFleetCmd = &cobra.Command{
	Use:   "delete-fleet",
	Short: "Deletes the specified fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteFleetCmd).Standalone()

	appstream_deleteFleetCmd.Flags().String("name", "", "The name of the fleet.")
	appstream_deleteFleetCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_deleteFleetCmd)
}
