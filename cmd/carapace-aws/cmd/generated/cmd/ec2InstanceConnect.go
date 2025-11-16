package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2InstanceConnectCmd = &cobra.Command{
	Use:   "ec2-instance-connect",
	Short: "This is the *Amazon EC2 Instance Connect API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2InstanceConnectCmd).Standalone()

	rootCmd.AddCommand(ec2InstanceConnectCmd)
}
