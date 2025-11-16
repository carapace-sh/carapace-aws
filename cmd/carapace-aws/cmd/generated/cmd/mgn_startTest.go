package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_startTestCmd = &cobra.Command{
	Use:   "start-test",
	Short: "Launches a Test Instance for specific Source Servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_startTestCmd).Standalone()

	mgn_startTestCmd.Flags().String("account-id", "", "Start Test for Account ID.")
	mgn_startTestCmd.Flags().String("source-server-ids", "", "Start Test for Source Server IDs.")
	mgn_startTestCmd.Flags().String("tags", "", "Start Test by Tags.")
	mgn_startTestCmd.MarkFlagRequired("source-server-ids")
	mgnCmd.AddCommand(mgn_startTestCmd)
}
