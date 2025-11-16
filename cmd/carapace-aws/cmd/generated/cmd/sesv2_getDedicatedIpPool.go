package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getDedicatedIpPoolCmd = &cobra.Command{
	Use:   "get-dedicated-ip-pool",
	Short: "Retrieve information about the dedicated pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getDedicatedIpPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getDedicatedIpPoolCmd).Standalone()

		sesv2_getDedicatedIpPoolCmd.Flags().String("pool-name", "", "The name of the dedicated IP pool to retrieve.")
		sesv2_getDedicatedIpPoolCmd.MarkFlagRequired("pool-name")
	})
	sesv2Cmd.AddCommand(sesv2_getDedicatedIpPoolCmd)
}
