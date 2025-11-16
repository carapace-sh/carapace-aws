package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listFuotaTasksCmd = &cobra.Command{
	Use:   "list-fuota-tasks",
	Short: "Lists the FUOTA tasks registered to your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listFuotaTasksCmd).Standalone()

	iotwireless_listFuotaTasksCmd.Flags().String("max-results", "", "")
	iotwireless_listFuotaTasksCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
	iotwirelessCmd.AddCommand(iotwireless_listFuotaTasksCmd)
}
