package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeNamespaceCmd = &cobra.Command{
	Use:   "describe-namespace",
	Short: "Describes the current namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeNamespaceCmd).Standalone()

	quicksight_describeNamespaceCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that contains the Quick Sight namespace that you want to describe.")
	quicksight_describeNamespaceCmd.Flags().String("namespace", "", "The namespace that you want to describe.")
	quicksight_describeNamespaceCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeNamespaceCmd.MarkFlagRequired("namespace")
	quicksightCmd.AddCommand(quicksight_describeNamespaceCmd)
}
