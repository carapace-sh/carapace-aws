package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getDedicatedIpsCmd = &cobra.Command{
	Use:   "get-dedicated-ips",
	Short: "List the dedicated IP addresses that are associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getDedicatedIpsCmd).Standalone()

	sesv2_getDedicatedIpsCmd.Flags().String("next-token", "", "A token returned from a previous call to `GetDedicatedIps` to indicate the position of the dedicated IP pool in the list of IP pools.")
	sesv2_getDedicatedIpsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `GetDedicatedIpsRequest`.")
	sesv2_getDedicatedIpsCmd.Flags().String("pool-name", "", "The name of the IP pool that the dedicated IP address is associated with.")
	sesv2Cmd.AddCommand(sesv2_getDedicatedIpsCmd)
}
