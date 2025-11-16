package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eksAuthCmd = &cobra.Command{
	Use:   "eks-auth",
	Short: "The Amazon EKS Auth API and the `AssumeRoleForPodIdentity` action are only used by the EKS Pod Identity Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eksAuthCmd).Standalone()

	rootCmd.AddCommand(eksAuthCmd)
}
