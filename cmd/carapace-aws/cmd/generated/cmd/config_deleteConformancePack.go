package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteConformancePackCmd = &cobra.Command{
	Use:   "delete-conformance-pack",
	Short: "Deletes the specified conformance pack and all the Config rules, remediation actions, and all evaluation results within that conformance pack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteConformancePackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteConformancePackCmd).Standalone()

		config_deleteConformancePackCmd.Flags().String("conformance-pack-name", "", "Name of the conformance pack you want to delete.")
		config_deleteConformancePackCmd.MarkFlagRequired("conformance-pack-name")
	})
	configCmd.AddCommand(config_deleteConformancePackCmd)
}
