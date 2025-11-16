package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "Creates a rule in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createRuleCmd).Standalone()

	datazone_createRuleCmd.Flags().String("action", "", "The action of the rule.")
	datazone_createRuleCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_createRuleCmd.Flags().String("description", "", "The description of the rule.")
	datazone_createRuleCmd.Flags().String("detail", "", "The detail of the rule.")
	datazone_createRuleCmd.Flags().String("domain-identifier", "", "The ID of the domain where the rule is created.")
	datazone_createRuleCmd.Flags().String("name", "", "The name of the rule.")
	datazone_createRuleCmd.Flags().String("scope", "", "The scope of the rule.")
	datazone_createRuleCmd.Flags().String("target", "", "The target of the rule.")
	datazone_createRuleCmd.MarkFlagRequired("action")
	datazone_createRuleCmd.MarkFlagRequired("detail")
	datazone_createRuleCmd.MarkFlagRequired("domain-identifier")
	datazone_createRuleCmd.MarkFlagRequired("name")
	datazone_createRuleCmd.MarkFlagRequired("scope")
	datazone_createRuleCmd.MarkFlagRequired("target")
	datazoneCmd.AddCommand(datazone_createRuleCmd)
}
