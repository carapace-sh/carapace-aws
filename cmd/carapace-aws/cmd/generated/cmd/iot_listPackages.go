package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listPackagesCmd = &cobra.Command{
	Use:   "list-packages",
	Short: "Lists the software packages associated to the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listPackagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listPackagesCmd).Standalone()

		iot_listPackagesCmd.Flags().String("max-results", "", "The maximum number of results returned at one time.")
		iot_listPackagesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	iotCmd.AddCommand(iot_listPackagesCmd)
}
