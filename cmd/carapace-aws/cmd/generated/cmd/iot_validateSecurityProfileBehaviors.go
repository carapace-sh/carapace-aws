package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_validateSecurityProfileBehaviorsCmd = &cobra.Command{
	Use:   "validate-security-profile-behaviors",
	Short: "Validates a Device Defender security profile behaviors specification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_validateSecurityProfileBehaviorsCmd).Standalone()

	iot_validateSecurityProfileBehaviorsCmd.Flags().String("behaviors", "", "Specifies the behaviors that, when violated by a device (thing), cause an alert.")
	iot_validateSecurityProfileBehaviorsCmd.MarkFlagRequired("behaviors")
	iotCmd.AddCommand(iot_validateSecurityProfileBehaviorsCmd)
}
