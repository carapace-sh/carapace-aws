package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listActiveViolationsCmd = &cobra.Command{
	Use:   "list-active-violations",
	Short: "Lists the active violations for a given Device Defender security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listActiveViolationsCmd).Standalone()

	iot_listActiveViolationsCmd.Flags().String("behavior-criteria-type", "", "The criteria for a behavior.")
	iot_listActiveViolationsCmd.Flags().String("list-suppressed-alerts", "", "A list of all suppressed alerts.")
	iot_listActiveViolationsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listActiveViolationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iot_listActiveViolationsCmd.Flags().String("security-profile-name", "", "The name of the Device Defender security profile for which violations are listed.")
	iot_listActiveViolationsCmd.Flags().String("thing-name", "", "The name of the thing whose active violations are listed.")
	iot_listActiveViolationsCmd.Flags().String("verification-state", "", "The verification state of the violation (detect alarm).")
	iotCmd.AddCommand(iot_listActiveViolationsCmd)
}
