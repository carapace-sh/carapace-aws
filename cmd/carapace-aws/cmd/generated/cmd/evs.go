package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evsCmd = &cobra.Command{
	Use:   "evs",
	Short: "Amazon Elastic VMware Service (Amazon EVS) is a service that you can use to deploy a VMware Cloud Foundation (VCF) software environment directly on EC2 bare metal instances within an Amazon Virtual Private Cloud (VPC).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evsCmd).Standalone()

	})
	rootCmd.AddCommand(evsCmd)
}
