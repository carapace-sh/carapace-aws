package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteRegistrationCodeCmd = &cobra.Command{
	Use:   "delete-registration-code",
	Short: "Deletes a CA certificate registration code.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteRegistrationCodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteRegistrationCodeCmd).Standalone()

	})
	iotCmd.AddCommand(iot_deleteRegistrationCodeCmd)
}
