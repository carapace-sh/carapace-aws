package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_updateSecurityCmd = &cobra.Command{
	Use:   "update-security",
	Short: "Updates the security settings for the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_updateSecurityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_updateSecurityCmd).Standalone()

		kafka_updateSecurityCmd.Flags().String("client-authentication", "", "Includes all client authentication related information.")
		kafka_updateSecurityCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_updateSecurityCmd.Flags().String("current-version", "", "The version of the MSK cluster to update.")
		kafka_updateSecurityCmd.Flags().String("encryption-info", "", "Includes all encryption-related information.")
		kafka_updateSecurityCmd.MarkFlagRequired("cluster-arn")
		kafka_updateSecurityCmd.MarkFlagRequired("current-version")
	})
	kafkaCmd.AddCommand(kafka_updateSecurityCmd)
}
