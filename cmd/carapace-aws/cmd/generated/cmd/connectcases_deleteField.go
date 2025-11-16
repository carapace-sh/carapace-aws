package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_deleteFieldCmd = &cobra.Command{
	Use:   "delete-field",
	Short: "Deletes a field from a cases template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_deleteFieldCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_deleteFieldCmd).Standalone()

		connectcases_deleteFieldCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_deleteFieldCmd.Flags().String("field-id", "", "Unique identifier of the field.")
		connectcases_deleteFieldCmd.MarkFlagRequired("domain-id")
		connectcases_deleteFieldCmd.MarkFlagRequired("field-id")
	})
	connectcasesCmd.AddCommand(connectcases_deleteFieldCmd)
}
