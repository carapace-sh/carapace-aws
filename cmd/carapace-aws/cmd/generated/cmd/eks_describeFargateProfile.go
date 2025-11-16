package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeFargateProfileCmd = &cobra.Command{
	Use:   "describe-fargate-profile",
	Short: "Describes an Fargate profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeFargateProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_describeFargateProfileCmd).Standalone()

		eks_describeFargateProfileCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_describeFargateProfileCmd.Flags().String("fargate-profile-name", "", "The name of the Fargate profile to describe.")
		eks_describeFargateProfileCmd.MarkFlagRequired("cluster-name")
		eks_describeFargateProfileCmd.MarkFlagRequired("fargate-profile-name")
	})
	eksCmd.AddCommand(eks_describeFargateProfileCmd)
}
