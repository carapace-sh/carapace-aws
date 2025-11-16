package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_putSchemaCmd = &cobra.Command{
	Use:   "put-schema",
	Short: "Creates or updates the policy schema in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_putSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_putSchemaCmd).Standalone()

		verifiedpermissions_putSchemaCmd.Flags().String("definition", "", "Specifies the definition of the schema to be stored.")
		verifiedpermissions_putSchemaCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store in which to place the schema.")
		verifiedpermissions_putSchemaCmd.MarkFlagRequired("definition")
		verifiedpermissions_putSchemaCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_putSchemaCmd)
}
