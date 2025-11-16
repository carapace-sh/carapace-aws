package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeClientAuthenticationSettingsCmd = &cobra.Command{
	Use:   "describe-client-authentication-settings",
	Short: "Retrieves information about the type of client authentication for the specified directory, if the type is specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeClientAuthenticationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeClientAuthenticationSettingsCmd).Standalone()

		ds_describeClientAuthenticationSettingsCmd.Flags().String("directory-id", "", "The identifier of the directory for which to retrieve information.")
		ds_describeClientAuthenticationSettingsCmd.Flags().String("limit", "", "The maximum number of items to return.")
		ds_describeClientAuthenticationSettingsCmd.Flags().String("next-token", "", "The *DescribeClientAuthenticationSettingsResult.NextToken* value from a previous call to [DescribeClientAuthenticationSettings]().")
		ds_describeClientAuthenticationSettingsCmd.Flags().String("type", "", "The type of client authentication for which to retrieve information.")
		ds_describeClientAuthenticationSettingsCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_describeClientAuthenticationSettingsCmd)
}
