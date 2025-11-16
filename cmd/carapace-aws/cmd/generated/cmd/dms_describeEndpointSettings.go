package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeEndpointSettingsCmd = &cobra.Command{
	Use:   "describe-endpoint-settings",
	Short: "Returns information about the possible endpoint settings available when you create an endpoint for a specific database engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeEndpointSettingsCmd).Standalone()

	dms_describeEndpointSettingsCmd.Flags().String("engine-name", "", "The database engine used for your source or target endpoint.")
	dms_describeEndpointSettingsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeEndpointSettingsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dms_describeEndpointSettingsCmd.MarkFlagRequired("engine-name")
	dmsCmd.AddCommand(dms_describeEndpointSettingsCmd)
}
