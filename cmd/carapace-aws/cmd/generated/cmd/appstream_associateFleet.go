package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_associateFleetCmd = &cobra.Command{
	Use:   "associate-fleet",
	Short: "Associates the specified fleet with the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_associateFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_associateFleetCmd).Standalone()

		appstream_associateFleetCmd.Flags().String("fleet-name", "", "The name of the fleet.")
		appstream_associateFleetCmd.Flags().String("stack-name", "", "The name of the stack.")
		appstream_associateFleetCmd.MarkFlagRequired("fleet-name")
		appstream_associateFleetCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_associateFleetCmd)
}
