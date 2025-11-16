package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efsCmd = &cobra.Command{
	Use:   "efs",
	Short: "Amazon Elastic File System\n\nAmazon Elastic File System (Amazon EFS) provides simple, scalable file storage for use with Amazon EC2 Linux and Mac instances in the Amazon Web Services Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efsCmd).Standalone()

	rootCmd.AddCommand(efsCmd)
}
