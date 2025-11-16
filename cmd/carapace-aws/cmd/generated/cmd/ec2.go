package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Amazon Elastic Compute Cloud\n\nYou can access the features of Amazon Elastic Compute Cloud (Amazon EC2) programmatically.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2Cmd).Standalone()

	rootCmd.AddCommand(ec2Cmd)
}
