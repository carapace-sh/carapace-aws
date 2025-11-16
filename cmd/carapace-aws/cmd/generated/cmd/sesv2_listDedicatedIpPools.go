package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listDedicatedIpPoolsCmd = &cobra.Command{
	Use:   "list-dedicated-ip-pools",
	Short: "List all of the dedicated IP pools that exist in your Amazon Web Services account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listDedicatedIpPoolsCmd).Standalone()

	sesv2_listDedicatedIpPoolsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListDedicatedIpPools` to indicate the position in the list of dedicated IP pools.")
	sesv2_listDedicatedIpPoolsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListDedicatedIpPools`.")
	sesv2Cmd.AddCommand(sesv2_listDedicatedIpPoolsCmd)
}
