package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_getLandingZoneOperationCmd = &cobra.Command{
	Use:   "get-landing-zone-operation",
	Short: "Returns the status of the specified landing zone operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_getLandingZoneOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_getLandingZoneOperationCmd).Standalone()

		controltower_getLandingZoneOperationCmd.Flags().String("operation-identifier", "", "A unique identifier assigned to a landing zone operation.")
		controltower_getLandingZoneOperationCmd.MarkFlagRequired("operation-identifier")
	})
	controltowerCmd.AddCommand(controltower_getLandingZoneOperationCmd)
}
