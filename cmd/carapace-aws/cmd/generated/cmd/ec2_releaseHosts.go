package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_releaseHostsCmd = &cobra.Command{
	Use:   "release-hosts",
	Short: "When you no longer want to use an On-Demand Dedicated Host it can be released.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_releaseHostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_releaseHostsCmd).Standalone()

		ec2_releaseHostsCmd.Flags().String("host-ids", "", "The IDs of the Dedicated Hosts to release.")
		ec2_releaseHostsCmd.MarkFlagRequired("host-ids")
	})
	ec2Cmd.AddCommand(ec2_releaseHostsCmd)
}
