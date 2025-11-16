package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_disassociateApplicationFleetCmd = &cobra.Command{
	Use:   "disassociate-application-fleet",
	Short: "Disassociates the specified application from the fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_disassociateApplicationFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_disassociateApplicationFleetCmd).Standalone()

		appstream_disassociateApplicationFleetCmd.Flags().String("application-arn", "", "The ARN of the application.")
		appstream_disassociateApplicationFleetCmd.Flags().String("fleet-name", "", "The name of the fleet.")
		appstream_disassociateApplicationFleetCmd.MarkFlagRequired("application-arn")
		appstream_disassociateApplicationFleetCmd.MarkFlagRequired("fleet-name")
	})
	appstreamCmd.AddCommand(appstream_disassociateApplicationFleetCmd)
}
