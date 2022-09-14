package job

import (
	"fmt"
	"github.com/spf13/cobra"
	"gosky/infra/app"
)

var JobCmd = &cobra.Command{
	Use: "job",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.TimenowInTimezone())
	},
}
