package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_detachSecurityProfileCmd = &cobra.Command{
	Use:   "detach-security-profile",
	Short: "Disassociates a Device Defender security profile from a thing group or from this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_detachSecurityProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_detachSecurityProfileCmd).Standalone()

		iot_detachSecurityProfileCmd.Flags().String("security-profile-name", "", "The security profile that is detached.")
		iot_detachSecurityProfileCmd.Flags().String("security-profile-target-arn", "", "The ARN of the thing group from which the security profile is detached.")
		iot_detachSecurityProfileCmd.MarkFlagRequired("security-profile-name")
		iot_detachSecurityProfileCmd.MarkFlagRequired("security-profile-target-arn")
	})
	iotCmd.AddCommand(iot_detachSecurityProfileCmd)
}
