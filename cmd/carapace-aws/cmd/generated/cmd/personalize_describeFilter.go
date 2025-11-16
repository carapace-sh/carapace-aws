package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeFilterCmd = &cobra.Command{
	Use:   "describe-filter",
	Short: "Describes a filter's properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeFilterCmd).Standalone()

	personalize_describeFilterCmd.Flags().String("filter-arn", "", "The ARN of the filter to describe.")
	personalize_describeFilterCmd.MarkFlagRequired("filter-arn")
	personalizeCmd.AddCommand(personalize_describeFilterCmd)
}
