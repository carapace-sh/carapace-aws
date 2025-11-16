package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConfigRulesCmd = &cobra.Command{
	Use:   "describe-config-rules",
	Short: "Returns details about your Config rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConfigRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeConfigRulesCmd).Standalone()

		config_describeConfigRulesCmd.Flags().String("config-rule-names", "", "The names of the Config rules for which you want details.")
		config_describeConfigRulesCmd.Flags().String("filters", "", "Returns a list of Detective or Proactive Config rules.")
		config_describeConfigRulesCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	})
	configCmd.AddCommand(config_describeConfigRulesCmd)
}
