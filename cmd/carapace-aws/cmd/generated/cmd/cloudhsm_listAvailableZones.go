package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_listAvailableZonesCmd = &cobra.Command{
	Use:   "list-available-zones",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_listAvailableZonesCmd).Standalone()

	cloudhsmCmd.AddCommand(cloudhsm_listAvailableZonesCmd)
}
