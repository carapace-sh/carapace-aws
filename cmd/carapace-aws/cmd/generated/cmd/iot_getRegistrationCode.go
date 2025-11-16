package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getRegistrationCodeCmd = &cobra.Command{
	Use:   "get-registration-code",
	Short: "Gets a registration code used to register a CA certificate with IoT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getRegistrationCodeCmd).Standalone()

	iotCmd.AddCommand(iot_getRegistrationCodeCmd)
}
