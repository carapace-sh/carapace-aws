package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_startCutoverCmd = &cobra.Command{
	Use:   "start-cutover",
	Short: "Launches a Cutover Instance for specific Source Servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_startCutoverCmd).Standalone()

	mgn_startCutoverCmd.Flags().String("account-id", "", "Start Cutover by Account IDs")
	mgn_startCutoverCmd.Flags().String("source-server-ids", "", "Start Cutover by Source Server IDs.")
	mgn_startCutoverCmd.Flags().String("tags", "", "Start Cutover by Tags.")
	mgn_startCutoverCmd.MarkFlagRequired("source-server-ids")
	mgnCmd.AddCommand(mgn_startCutoverCmd)
}
