package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_startFleetCmd = &cobra.Command{
	Use:   "start-fleet",
	Short: "Starts the specified fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_startFleetCmd).Standalone()

	appstream_startFleetCmd.Flags().String("name", "", "The name of the fleet.")
	appstream_startFleetCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_startFleetCmd)
}
