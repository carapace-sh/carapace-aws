package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_attachSecurityProfileCmd = &cobra.Command{
	Use:   "attach-security-profile",
	Short: "Associates a Device Defender security profile with a thing group or this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_attachSecurityProfileCmd).Standalone()

	iot_attachSecurityProfileCmd.Flags().String("security-profile-name", "", "The security profile that is attached.")
	iot_attachSecurityProfileCmd.Flags().String("security-profile-target-arn", "", "The ARN of the target (thing group) to which the security profile is attached.")
	iot_attachSecurityProfileCmd.MarkFlagRequired("security-profile-name")
	iot_attachSecurityProfileCmd.MarkFlagRequired("security-profile-target-arn")
	iotCmd.AddCommand(iot_attachSecurityProfileCmd)
}
