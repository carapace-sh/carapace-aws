package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowballCmd = &cobra.Command{
	Use:   "snowball",
	Short: "The Amazon Web Services Snow Family provides a petabyte-scale data transport solution that uses secure devices to transfer large amounts of data between your on-premises data centers and Amazon Simple Storage Service (Amazon S3).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowballCmd).Standalone()

	rootCmd.AddCommand(snowballCmd)
}
