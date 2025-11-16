package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_disassociateFleetCmd = &cobra.Command{
	Use:   "disassociate-fleet",
	Short: "Disassociates the specified fleet from the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_disassociateFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_disassociateFleetCmd).Standalone()

		appstream_disassociateFleetCmd.Flags().String("fleet-name", "", "The name of the fleet.")
		appstream_disassociateFleetCmd.Flags().String("stack-name", "", "The name of the stack.")
		appstream_disassociateFleetCmd.MarkFlagRequired("fleet-name")
		appstream_disassociateFleetCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_disassociateFleetCmd)
}
