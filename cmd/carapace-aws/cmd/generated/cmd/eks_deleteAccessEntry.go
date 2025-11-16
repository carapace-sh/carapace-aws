package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deleteAccessEntryCmd = &cobra.Command{
	Use:   "delete-access-entry",
	Short: "Deletes an access entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deleteAccessEntryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_deleteAccessEntryCmd).Standalone()

		eks_deleteAccessEntryCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_deleteAccessEntryCmd.Flags().String("principal-arn", "", "The ARN of the IAM principal for the `AccessEntry`.")
		eks_deleteAccessEntryCmd.MarkFlagRequired("cluster-name")
		eks_deleteAccessEntryCmd.MarkFlagRequired("principal-arn")
	})
	eksCmd.AddCommand(eks_deleteAccessEntryCmd)
}
