package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_updateFieldCmd = &cobra.Command{
	Use:   "update-field",
	Short: "Updates the properties of an existing field.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_updateFieldCmd).Standalone()

	connectcases_updateFieldCmd.Flags().String("description", "", "The description of a field.")
	connectcases_updateFieldCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_updateFieldCmd.Flags().String("field-id", "", "The unique identifier of a field.")
	connectcases_updateFieldCmd.Flags().String("name", "", "The name of the field.")
	connectcases_updateFieldCmd.MarkFlagRequired("domain-id")
	connectcases_updateFieldCmd.MarkFlagRequired("field-id")
	connectcasesCmd.AddCommand(connectcases_updateFieldCmd)
}
