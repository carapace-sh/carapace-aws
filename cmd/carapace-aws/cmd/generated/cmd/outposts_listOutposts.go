package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listOutpostsCmd = &cobra.Command{
	Use:   "list-outposts",
	Short: "Lists the Outposts for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listOutpostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_listOutpostsCmd).Standalone()

		outposts_listOutpostsCmd.Flags().String("availability-zone-filter", "", "Filters the results by Availability Zone (for example, `us-east-1a`).")
		outposts_listOutpostsCmd.Flags().String("availability-zone-id-filter", "", "Filters the results by AZ ID (for example, `use1-az1`).")
		outposts_listOutpostsCmd.Flags().String("life-cycle-status-filter", "", "Filters the results by the lifecycle status.")
		outposts_listOutpostsCmd.Flags().String("max-results", "", "")
		outposts_listOutpostsCmd.Flags().String("next-token", "", "")
	})
	outpostsCmd.AddCommand(outposts_listOutpostsCmd)
}
