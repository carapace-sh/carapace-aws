package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createSecurityProfileCmd = &cobra.Command{
	Use:   "create-security-profile",
	Short: "Creates a Device Defender security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createSecurityProfileCmd).Standalone()

	iot_createSecurityProfileCmd.Flags().String("additional-metrics-to-retain", "", "*Please use [CreateSecurityProfileRequest$additionalMetricsToRetainV2]() instead.*")
	iot_createSecurityProfileCmd.Flags().String("additional-metrics-to-retain-v2", "", "A list of metrics whose data is retained (stored).")
	iot_createSecurityProfileCmd.Flags().String("alert-targets", "", "Specifies the destinations to which alerts are sent.")
	iot_createSecurityProfileCmd.Flags().String("behaviors", "", "Specifies the behaviors that, when violated by a device (thing), cause an alert.")
	iot_createSecurityProfileCmd.Flags().String("metrics-export-config", "", "Specifies the MQTT topic and role ARN required for metric export.")
	iot_createSecurityProfileCmd.Flags().String("security-profile-description", "", "A description of the security profile.")
	iot_createSecurityProfileCmd.Flags().String("security-profile-name", "", "The name you are giving to the security profile.")
	iot_createSecurityProfileCmd.Flags().String("tags", "", "Metadata that can be used to manage the security profile.")
	iot_createSecurityProfileCmd.MarkFlagRequired("security-profile-name")
	iotCmd.AddCommand(iot_createSecurityProfileCmd)
}
