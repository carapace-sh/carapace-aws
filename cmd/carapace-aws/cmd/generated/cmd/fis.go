package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fisCmd = &cobra.Command{
	Use:   "fis",
	Short: "Amazon Web Services Fault Injection Service is a managed service that enables you to perform fault injection experiments on your Amazon Web Services workloads.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fisCmd).Standalone()

	})
	rootCmd.AddCommand(fisCmd)
}
