package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeSchedulingPoliciesCmd = &cobra.Command{
	Use:   "describe-scheduling-policies",
	Short: "Describes one or more of your scheduling policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeSchedulingPoliciesCmd).Standalone()

	batch_describeSchedulingPoliciesCmd.Flags().String("arns", "", "A list of up to 100 scheduling policy Amazon Resource Name (ARN) entries.")
	batch_describeSchedulingPoliciesCmd.MarkFlagRequired("arns")
	batchCmd.AddCommand(batch_describeSchedulingPoliciesCmd)
}
