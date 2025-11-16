package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_associateApplicationFleetCmd = &cobra.Command{
	Use:   "associate-application-fleet",
	Short: "Associates the specified application with the specified fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_associateApplicationFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_associateApplicationFleetCmd).Standalone()

		appstream_associateApplicationFleetCmd.Flags().String("application-arn", "", "The ARN of the application.")
		appstream_associateApplicationFleetCmd.Flags().String("fleet-name", "", "The name of the fleet.")
		appstream_associateApplicationFleetCmd.MarkFlagRequired("application-arn")
		appstream_associateApplicationFleetCmd.MarkFlagRequired("fleet-name")
	})
	appstreamCmd.AddCommand(appstream_associateApplicationFleetCmd)
}
