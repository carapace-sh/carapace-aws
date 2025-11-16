package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteRuleCmd = &cobra.Command{
	Use:   "delete-rule",
	Short: "Deletes a rule in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteRuleCmd).Standalone()

	datazone_deleteRuleCmd.Flags().String("domain-identifier", "", "The ID of the domain that where the rule is to be deleted.")
	datazone_deleteRuleCmd.Flags().String("identifier", "", "The ID of the rule that is to be deleted.")
	datazone_deleteRuleCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteRuleCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteRuleCmd)
}
