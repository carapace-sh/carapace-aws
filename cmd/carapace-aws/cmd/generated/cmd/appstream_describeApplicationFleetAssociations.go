package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_describeApplicationFleetAssociationsCmd = &cobra.Command{
	Use:   "describe-application-fleet-associations",
	Short: "Retrieves a list that describes one or more application fleet associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_describeApplicationFleetAssociationsCmd).Standalone()

	appstream_describeApplicationFleetAssociationsCmd.Flags().String("application-arn", "", "The ARN of the application.")
	appstream_describeApplicationFleetAssociationsCmd.Flags().String("fleet-name", "", "The name of the fleet.")
	appstream_describeApplicationFleetAssociationsCmd.Flags().String("max-results", "", "The maximum size of each page of results.")
	appstream_describeApplicationFleetAssociationsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	appstreamCmd.AddCommand(appstream_describeApplicationFleetAssociationsCmd)
}
