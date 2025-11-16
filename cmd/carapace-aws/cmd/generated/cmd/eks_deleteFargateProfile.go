package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deleteFargateProfileCmd = &cobra.Command{
	Use:   "delete-fargate-profile",
	Short: "Deletes an Fargate profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deleteFargateProfileCmd).Standalone()

	eks_deleteFargateProfileCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_deleteFargateProfileCmd.Flags().String("fargate-profile-name", "", "The name of the Fargate profile to delete.")
	eks_deleteFargateProfileCmd.MarkFlagRequired("cluster-name")
	eks_deleteFargateProfileCmd.MarkFlagRequired("fargate-profile-name")
	eksCmd.AddCommand(eks_deleteFargateProfileCmd)
}
