package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listSecurityProfilesCmd = &cobra.Command{
	Use:   "list-security-profiles",
	Short: "Lists the Device Defender security profiles you've created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listSecurityProfilesCmd).Standalone()

	iot_listSecurityProfilesCmd.Flags().String("dimension-name", "", "A filter to limit results to the security profiles that use the defined dimension.")
	iot_listSecurityProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listSecurityProfilesCmd.Flags().String("metric-name", "", "The name of the custom metric.")
	iot_listSecurityProfilesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	iotCmd.AddCommand(iot_listSecurityProfilesCmd)
}
