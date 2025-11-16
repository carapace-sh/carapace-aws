package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_describeHapgCmd = &cobra.Command{
	Use:   "describe-hapg",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_describeHapgCmd).Standalone()

	cloudhsm_describeHapgCmd.Flags().String("hapg-arn", "", "The ARN of the high-availability partition group to describe.")
	cloudhsm_describeHapgCmd.MarkFlagRequired("hapg-arn")
	cloudhsmCmd.AddCommand(cloudhsm_describeHapgCmd)
}
