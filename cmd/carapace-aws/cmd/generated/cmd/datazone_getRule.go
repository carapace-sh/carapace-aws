package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getRuleCmd = &cobra.Command{
	Use:   "get-rule",
	Short: "Gets the details of a rule in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getRuleCmd).Standalone()

		datazone_getRuleCmd.Flags().String("domain-identifier", "", "The ID of the domain where the `GetRule` action is to be invoked.")
		datazone_getRuleCmd.Flags().String("identifier", "", "The ID of the rule.")
		datazone_getRuleCmd.Flags().String("revision", "", "The revision of the rule.")
		datazone_getRuleCmd.MarkFlagRequired("domain-identifier")
		datazone_getRuleCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getRuleCmd)
}
