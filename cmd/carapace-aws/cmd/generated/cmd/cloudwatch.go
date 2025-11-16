package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatchCmd = &cobra.Command{
	Use:   "cloudwatch",
	Short: "Amazon CloudWatch monitors your Amazon Web Services (Amazon Web Services) resources and the applications you run on Amazon Web Services in real time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatchCmd).Standalone()

	rootCmd.AddCommand(cloudwatchCmd)
}
