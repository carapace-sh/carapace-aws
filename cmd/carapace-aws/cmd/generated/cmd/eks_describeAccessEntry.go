package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeAccessEntryCmd = &cobra.Command{
	Use:   "describe-access-entry",
	Short: "Describes an access entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeAccessEntryCmd).Standalone()

	eks_describeAccessEntryCmd.Flags().String("cluster-name", "", "The name of your cluster.")
	eks_describeAccessEntryCmd.Flags().String("principal-arn", "", "The ARN of the IAM principal for the `AccessEntry`.")
	eks_describeAccessEntryCmd.MarkFlagRequired("cluster-name")
	eks_describeAccessEntryCmd.MarkFlagRequired("principal-arn")
	eksCmd.AddCommand(eks_describeAccessEntryCmd)
}
