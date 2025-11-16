package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_createFieldCmd = &cobra.Command{
	Use:   "create-field",
	Short: "Creates a field in the Cases domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_createFieldCmd).Standalone()

	connectcases_createFieldCmd.Flags().String("description", "", "The description of the field.")
	connectcases_createFieldCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
	connectcases_createFieldCmd.Flags().String("name", "", "The name of the field.")
	connectcases_createFieldCmd.Flags().String("type", "", "Defines the data type, some system constraints, and default display of the field.")
	connectcases_createFieldCmd.MarkFlagRequired("domain-id")
	connectcases_createFieldCmd.MarkFlagRequired("name")
	connectcases_createFieldCmd.MarkFlagRequired("type")
	connectcasesCmd.AddCommand(connectcases_createFieldCmd)
}
