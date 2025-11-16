package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateSecurityProfileCmd = &cobra.Command{
	Use:   "update-security-profile",
	Short: "Updates a Device Defender security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateSecurityProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateSecurityProfileCmd).Standalone()

		iot_updateSecurityProfileCmd.Flags().String("additional-metrics-to-retain", "", "*Please use [UpdateSecurityProfileRequest$additionalMetricsToRetainV2]() instead.*")
		iot_updateSecurityProfileCmd.Flags().String("additional-metrics-to-retain-v2", "", "A list of metrics whose data is retained (stored).")
		iot_updateSecurityProfileCmd.Flags().String("alert-targets", "", "Where the alerts are sent.")
		iot_updateSecurityProfileCmd.Flags().String("behaviors", "", "Specifies the behaviors that, when violated by a device (thing), cause an alert.")
		iot_updateSecurityProfileCmd.Flags().String("delete-additional-metrics-to-retain", "", "If true, delete all `additionalMetricsToRetain` defined for this security profile.")
		iot_updateSecurityProfileCmd.Flags().String("delete-alert-targets", "", "If true, delete all `alertTargets` defined for this security profile.")
		iot_updateSecurityProfileCmd.Flags().String("delete-behaviors", "", "If true, delete all `behaviors` defined for this security profile.")
		iot_updateSecurityProfileCmd.Flags().String("delete-metrics-export-config", "", "Set the value as true to delete metrics export related configurations.")
		iot_updateSecurityProfileCmd.Flags().String("expected-version", "", "The expected version of the security profile.")
		iot_updateSecurityProfileCmd.Flags().String("metrics-export-config", "", "Specifies the MQTT topic and role ARN required for metric export.")
		iot_updateSecurityProfileCmd.Flags().String("security-profile-description", "", "A description of the security profile.")
		iot_updateSecurityProfileCmd.Flags().String("security-profile-name", "", "The name of the security profile you want to update.")
		iot_updateSecurityProfileCmd.MarkFlagRequired("security-profile-name")
	})
	iotCmd.AddCommand(iot_updateSecurityProfileCmd)
}
