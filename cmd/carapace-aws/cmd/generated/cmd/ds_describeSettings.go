package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeSettingsCmd = &cobra.Command{
	Use:   "describe-settings",
	Short: "Retrieves information about the configurable settings for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeSettingsCmd).Standalone()

	ds_describeSettingsCmd.Flags().String("directory-id", "", "The identifier of the directory for which to retrieve information.")
	ds_describeSettingsCmd.Flags().String("next-token", "", "The `DescribeSettingsResult.NextToken` value from a previous call to [DescribeSettings]().")
	ds_describeSettingsCmd.Flags().String("status", "", "The status of the directory settings for which to retrieve information.")
	ds_describeSettingsCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_describeSettingsCmd)
}
