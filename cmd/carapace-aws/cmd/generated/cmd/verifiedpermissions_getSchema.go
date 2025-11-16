package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_getSchemaCmd = &cobra.Command{
	Use:   "get-schema",
	Short: "Retrieve the details for the specified schema in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_getSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_getSchemaCmd).Standalone()

		verifiedpermissions_getSchemaCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the schema.")
		verifiedpermissions_getSchemaCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_getSchemaCmd)
}
