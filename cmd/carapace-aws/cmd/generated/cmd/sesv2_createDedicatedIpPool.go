package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_createDedicatedIpPoolCmd = &cobra.Command{
	Use:   "create-dedicated-ip-pool",
	Short: "Create a new pool of dedicated IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_createDedicatedIpPoolCmd).Standalone()

	sesv2_createDedicatedIpPoolCmd.Flags().String("pool-name", "", "The name of the dedicated IP pool.")
	sesv2_createDedicatedIpPoolCmd.Flags().String("scaling-mode", "", "The type of scaling mode.")
	sesv2_createDedicatedIpPoolCmd.Flags().String("tags", "", "An object that defines the tags (keys and values) that you want to associate with the pool.")
	sesv2_createDedicatedIpPoolCmd.MarkFlagRequired("pool-name")
	sesv2Cmd.AddCommand(sesv2_createDedicatedIpPoolCmd)
}
