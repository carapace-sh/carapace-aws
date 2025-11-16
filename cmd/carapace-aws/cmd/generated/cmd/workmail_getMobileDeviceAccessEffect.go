package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getMobileDeviceAccessEffectCmd = &cobra.Command{
	Use:   "get-mobile-device-access-effect",
	Short: "Simulates the effect of the mobile device access rules for the given attributes of a sample access event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getMobileDeviceAccessEffectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_getMobileDeviceAccessEffectCmd).Standalone()

		workmail_getMobileDeviceAccessEffectCmd.Flags().String("device-model", "", "Device model the simulated user will report.")
		workmail_getMobileDeviceAccessEffectCmd.Flags().String("device-operating-system", "", "Device operating system the simulated user will report.")
		workmail_getMobileDeviceAccessEffectCmd.Flags().String("device-type", "", "Device type the simulated user will report.")
		workmail_getMobileDeviceAccessEffectCmd.Flags().String("device-user-agent", "", "Device user agent the simulated user will report.")
		workmail_getMobileDeviceAccessEffectCmd.Flags().String("organization-id", "", "The WorkMail organization to simulate the access effect for.")
		workmail_getMobileDeviceAccessEffectCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_getMobileDeviceAccessEffectCmd)
}
