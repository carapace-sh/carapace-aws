package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeSecurityProfileCmd = &cobra.Command{
	Use:   "describe-security-profile",
	Short: "Gets information about a Device Defender security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeSecurityProfileCmd).Standalone()

	iot_describeSecurityProfileCmd.Flags().String("security-profile-name", "", "The name of the security profile whose information you want to get.")
	iot_describeSecurityProfileCmd.MarkFlagRequired("security-profile-name")
	iotCmd.AddCommand(iot_describeSecurityProfileCmd)
}
