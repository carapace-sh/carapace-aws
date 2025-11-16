package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeInsightCmd = &cobra.Command{
	Use:   "describe-insight",
	Short: "Returns details about an insight that you specify using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeInsightCmd).Standalone()

	eks_describeInsightCmd.Flags().String("cluster-name", "", "The name of the cluster to describe the insight for.")
	eks_describeInsightCmd.Flags().String("id", "", "The identity of the insight to describe.")
	eks_describeInsightCmd.MarkFlagRequired("cluster-name")
	eks_describeInsightCmd.MarkFlagRequired("id")
	eksCmd.AddCommand(eks_describeInsightCmd)
}
