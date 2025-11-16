package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_deleteTerminologyCmd = &cobra.Command{
	Use:   "delete-terminology",
	Short: "A synchronous action that deletes a custom terminology.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_deleteTerminologyCmd).Standalone()

	translate_deleteTerminologyCmd.Flags().String("name", "", "The name of the custom terminology being deleted.")
	translate_deleteTerminologyCmd.MarkFlagRequired("name")
	translateCmd.AddCommand(translate_deleteTerminologyCmd)
}
