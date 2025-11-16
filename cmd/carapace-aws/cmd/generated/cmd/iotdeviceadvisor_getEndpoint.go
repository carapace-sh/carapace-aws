package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotdeviceadvisor_getEndpointCmd = &cobra.Command{
	Use:   "get-endpoint",
	Short: "Gets information about an Device Advisor endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotdeviceadvisor_getEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotdeviceadvisor_getEndpointCmd).Standalone()

		iotdeviceadvisor_getEndpointCmd.Flags().String("authentication-method", "", "The authentication method used during the device connection.")
		iotdeviceadvisor_getEndpointCmd.Flags().String("certificate-arn", "", "The certificate ARN of the device.")
		iotdeviceadvisor_getEndpointCmd.Flags().String("device-role-arn", "", "The device role ARN of the device.")
		iotdeviceadvisor_getEndpointCmd.Flags().String("thing-arn", "", "The thing ARN of the device.")
	})
	iotdeviceadvisorCmd.AddCommand(iotdeviceadvisor_getEndpointCmd)
}
