package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var translate_importTerminologyCmd = &cobra.Command{
	Use:   "import-terminology",
	Short: "Creates or updates a custom terminology, depending on whether one already exists for the given terminology name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translate_importTerminologyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(translate_importTerminologyCmd).Standalone()

		translate_importTerminologyCmd.Flags().String("description", "", "The description of the custom terminology being imported.")
		translate_importTerminologyCmd.Flags().String("encryption-key", "", "The encryption key for the custom terminology being imported.")
		translate_importTerminologyCmd.Flags().String("merge-strategy", "", "The merge strategy of the custom terminology being imported.")
		translate_importTerminologyCmd.Flags().String("name", "", "The name of the custom terminology being imported.")
		translate_importTerminologyCmd.Flags().String("tags", "", "Tags to be associated with this resource.")
		translate_importTerminologyCmd.Flags().String("terminology-data", "", "The terminology data for the custom terminology being imported.")
		translate_importTerminologyCmd.MarkFlagRequired("merge-strategy")
		translate_importTerminologyCmd.MarkFlagRequired("name")
		translate_importTerminologyCmd.MarkFlagRequired("terminology-data")
	})
	translateCmd.AddCommand(translate_importTerminologyCmd)
}
