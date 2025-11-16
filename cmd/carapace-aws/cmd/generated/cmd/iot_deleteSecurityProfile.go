package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteSecurityProfileCmd = &cobra.Command{
	Use:   "delete-security-profile",
	Short: "Deletes a Device Defender security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteSecurityProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteSecurityProfileCmd).Standalone()

		iot_deleteSecurityProfileCmd.Flags().String("expected-version", "", "The expected version of the security profile.")
		iot_deleteSecurityProfileCmd.Flags().String("security-profile-name", "", "The name of the security profile to be deleted.")
		iot_deleteSecurityProfileCmd.MarkFlagRequired("security-profile-name")
	})
	iotCmd.AddCommand(iot_deleteSecurityProfileCmd)
}
