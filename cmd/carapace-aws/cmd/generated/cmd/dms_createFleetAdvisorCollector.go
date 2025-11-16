package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createFleetAdvisorCollectorCmd = &cobra.Command{
	Use:   "create-fleet-advisor-collector",
	Short: "End of support notice: On May 20, 2026, Amazon Web Services will end support for Amazon Web Services DMS Fleet Advisor;.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createFleetAdvisorCollectorCmd).Standalone()

	dms_createFleetAdvisorCollectorCmd.Flags().String("collector-name", "", "The name of your Fleet Advisor collector (for example, `sample-collector`).")
	dms_createFleetAdvisorCollectorCmd.Flags().String("description", "", "A summary description of your Fleet Advisor collector.")
	dms_createFleetAdvisorCollectorCmd.Flags().String("s3-bucket-name", "", "The Amazon S3 bucket that the Fleet Advisor collector uses to store inventory metadata.")
	dms_createFleetAdvisorCollectorCmd.Flags().String("service-access-role-arn", "", "The IAM role that grants permissions to access the specified Amazon S3 bucket.")
	dms_createFleetAdvisorCollectorCmd.MarkFlagRequired("collector-name")
	dms_createFleetAdvisorCollectorCmd.MarkFlagRequired("s3-bucket-name")
	dms_createFleetAdvisorCollectorCmd.MarkFlagRequired("service-access-role-arn")
	dmsCmd.AddCommand(dms_createFleetAdvisorCollectorCmd)
}
