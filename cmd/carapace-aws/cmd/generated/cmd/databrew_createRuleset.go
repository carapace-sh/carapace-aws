package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_createRulesetCmd = &cobra.Command{
	Use:   "create-ruleset",
	Short: "Creates a new ruleset that can be used in a profile job to validate the data quality of a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_createRulesetCmd).Standalone()

	databrew_createRulesetCmd.Flags().String("description", "", "The description of the ruleset.")
	databrew_createRulesetCmd.Flags().String("name", "", "The name of the ruleset to be created.")
	databrew_createRulesetCmd.Flags().String("rules", "", "A list of rules that are defined with the ruleset.")
	databrew_createRulesetCmd.Flags().String("tags", "", "Metadata tags to apply to the ruleset.")
	databrew_createRulesetCmd.Flags().String("target-arn", "", "The Amazon Resource Name (ARN) of a resource (dataset) that the ruleset is associated with.")
	databrew_createRulesetCmd.MarkFlagRequired("name")
	databrew_createRulesetCmd.MarkFlagRequired("rules")
	databrew_createRulesetCmd.MarkFlagRequired("target-arn")
	databrewCmd.AddCommand(databrew_createRulesetCmd)
}
