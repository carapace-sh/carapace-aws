package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ebsCmd = &cobra.Command{
	Use:   "ebs",
	Short: "You can use the Amazon Elastic Block Store (Amazon EBS) direct APIs to create Amazon EBS snapshots, write data directly to your snapshots, read data on your snapshots, and identify the differences or changes between two snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ebsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ebsCmd).Standalone()

	})
	rootCmd.AddCommand(ebsCmd)
}
