package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_deleteRulesetCmd = &cobra.Command{
	Use:   "delete-ruleset",
	Short: "Deletes a ruleset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_deleteRulesetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_deleteRulesetCmd).Standalone()

		databrew_deleteRulesetCmd.Flags().String("name", "", "The name of the ruleset to be deleted.")
		databrew_deleteRulesetCmd.MarkFlagRequired("name")
	})
	databrewCmd.AddCommand(databrew_deleteRulesetCmd)
}
