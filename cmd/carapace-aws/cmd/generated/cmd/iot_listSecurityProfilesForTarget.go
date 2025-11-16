package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listSecurityProfilesForTargetCmd = &cobra.Command{
	Use:   "list-security-profiles-for-target",
	Short: "Lists the Device Defender security profiles attached to a target (thing group).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listSecurityProfilesForTargetCmd).Standalone()

	iot_listSecurityProfilesForTargetCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listSecurityProfilesForTargetCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listSecurityProfilesForTargetCmd.Flags().String("recursive", "", "If true, return child groups too.")
	iot_listSecurityProfilesForTargetCmd.Flags().String("security-profile-target-arn", "", "The ARN of the target (thing group) whose attached security profiles you want to get.")
	iot_listSecurityProfilesForTargetCmd.MarkFlagRequired("security-profile-target-arn")
	iotCmd.AddCommand(iot_listSecurityProfilesForTargetCmd)
}
