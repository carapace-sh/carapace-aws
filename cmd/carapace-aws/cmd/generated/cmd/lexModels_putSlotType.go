package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_putSlotTypeCmd = &cobra.Command{
	Use:   "put-slot-type",
	Short: "Creates a custom slot type or replaces an existing custom slot type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_putSlotTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexModels_putSlotTypeCmd).Standalone()

		lexModels_putSlotTypeCmd.Flags().String("checksum", "", "Identifies a specific revision of the `$LATEST` version.")
		lexModels_putSlotTypeCmd.Flags().Bool("create-version", false, "When set to `true` a new numbered version of the slot type is created.")
		lexModels_putSlotTypeCmd.Flags().String("description", "", "A description of the slot type.")
		lexModels_putSlotTypeCmd.Flags().String("enumeration-values", "", "A list of `EnumerationValue` objects that defines the values that the slot type can take.")
		lexModels_putSlotTypeCmd.Flags().String("name", "", "The name of the slot type.")
		lexModels_putSlotTypeCmd.Flags().Bool("no-create-version", false, "When set to `true` a new numbered version of the slot type is created.")
		lexModels_putSlotTypeCmd.Flags().String("parent-slot-type-signature", "", "The built-in slot type used as the parent of the slot type.")
		lexModels_putSlotTypeCmd.Flags().String("slot-type-configurations", "", "Configuration information that extends the parent built-in slot type.")
		lexModels_putSlotTypeCmd.Flags().String("value-selection-strategy", "", "Determines the slot resolution strategy that Amazon Lex uses to return slot type values.")
		lexModels_putSlotTypeCmd.MarkFlagRequired("name")
		lexModels_putSlotTypeCmd.Flag("no-create-version").Hidden = true
	})
	lexModelsCmd.AddCommand(lexModels_putSlotTypeCmd)
}
