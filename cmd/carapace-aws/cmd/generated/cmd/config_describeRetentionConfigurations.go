package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeRetentionConfigurationsCmd = &cobra.Command{
	Use:   "describe-retention-configurations",
	Short: "Returns the details of one or more retention configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeRetentionConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeRetentionConfigurationsCmd).Standalone()

		config_describeRetentionConfigurationsCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_describeRetentionConfigurationsCmd.Flags().String("retention-configuration-names", "", "A list of names of retention configurations for which you want details.")
	})
	configCmd.AddCommand(config_describeRetentionConfigurationsCmd)
}
