package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_stopFleetCmd = &cobra.Command{
	Use:   "stop-fleet",
	Short: "Stops the specified fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_stopFleetCmd).Standalone()

	appstream_stopFleetCmd.Flags().String("name", "", "The name of the fleet.")
	appstream_stopFleetCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_stopFleetCmd)
}
