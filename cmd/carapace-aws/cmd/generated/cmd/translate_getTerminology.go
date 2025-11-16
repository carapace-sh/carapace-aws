package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_getTerminologyCmd = &cobra.Command{
	Use:   "get-terminology",
	Short: "Retrieves a custom terminology.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_getTerminologyCmd).Standalone()

	translate_getTerminologyCmd.Flags().String("name", "", "The name of the custom terminology being retrieved.")
	translate_getTerminologyCmd.Flags().String("terminology-data-format", "", "The data format of the custom terminology being retrieved.")
	translate_getTerminologyCmd.MarkFlagRequired("name")
	translateCmd.AddCommand(translate_getTerminologyCmd)
}
