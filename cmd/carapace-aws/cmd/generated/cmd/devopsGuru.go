package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuruCmd = &cobra.Command{
	Use:   "devops-guru",
	Short: "Amazon DevOps Guru is a fully managed service that helps you identify anomalous behavior in business critical operational applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuruCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuruCmd).Standalone()

	})
	rootCmd.AddCommand(devopsGuruCmd)
}
