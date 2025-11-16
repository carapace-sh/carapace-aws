package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listTargetsForSecurityProfileCmd = &cobra.Command{
	Use:   "list-targets-for-security-profile",
	Short: "Lists the targets (thing groups) associated with a given Device Defender security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listTargetsForSecurityProfileCmd).Standalone()

	iot_listTargetsForSecurityProfileCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listTargetsForSecurityProfileCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listTargetsForSecurityProfileCmd.Flags().String("security-profile-name", "", "The security profile.")
	iot_listTargetsForSecurityProfileCmd.MarkFlagRequired("security-profile-name")
	iotCmd.AddCommand(iot_listTargetsForSecurityProfileCmd)
}
