package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eksCmd = &cobra.Command{
	Use:   "eks",
	Short: "Amazon Elastic Kubernetes Service (Amazon EKS) is a managed service that makes it easy for you to run Kubernetes on Amazon Web Services without needing to setup or maintain your own Kubernetes control plane.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eksCmd).Standalone()

	})
	rootCmd.AddCommand(eksCmd)
}
