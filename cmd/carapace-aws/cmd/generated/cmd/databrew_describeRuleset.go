package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_describeRulesetCmd = &cobra.Command{
	Use:   "describe-ruleset",
	Short: "Retrieves detailed information about the ruleset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_describeRulesetCmd).Standalone()

	databrew_describeRulesetCmd.Flags().String("name", "", "The name of the ruleset to be described.")
	databrew_describeRulesetCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_describeRulesetCmd)
}
