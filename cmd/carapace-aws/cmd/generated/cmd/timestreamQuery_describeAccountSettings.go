package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_describeAccountSettingsCmd = &cobra.Command{
	Use:   "describe-account-settings",
	Short: "Describes the settings for your account that include the query pricing model and the configured maximum TCUs the service can use for your query workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_describeAccountSettingsCmd).Standalone()

	timestreamQueryCmd.AddCommand(timestreamQuery_describeAccountSettingsCmd)
}
