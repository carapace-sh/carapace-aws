package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getOfferingStatusCmd = &cobra.Command{
	Use:   "get-offering-status",
	Short: "Gets the current status and future status of all offerings purchased by an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getOfferingStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getOfferingStatusCmd).Standalone()

		devicefarm_getOfferingStatusCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	})
	devicefarmCmd.AddCommand(devicefarm_getOfferingStatusCmd)
}
