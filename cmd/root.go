package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "./BCAuth",
	Short: "BCAuth application",
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the application webserver",
	RunE: func(cmd *cobra.Command, args []string) error {
		path := os.Getenv("BCAUTH_CONFIG")
		if path == "" {
			return fmt.Errorf("BCAUTH_CONFIG env variable was not set")
		}
		return nil
	},
}

//var TestCmd = &cobra.Command{
//	Use:   "test",
//	Short: "Runs unit and mock tests",
//	RunE: func(c *cobra.Command, args []string) error {
//		var err error = nil
//		if c.Flag("config").Changed {
//			ConfigPath, err = c.Flags().GetString("config")
//		} else {
//			ConfigPath = c.Flag("config").DefValue
//		}
//		if err != nil {
//			return err
//		}
//		command := exec.Command("gorc", "test")
//		var out bytes.Buffer
//		command.Stdout = &out
//		err = command.Run()
//		if err != nil {
//			return err
//		}
//		fmt.Printf("in all caps: %s\n", out.String())
//
//		return nil
//	},
//}

var ConfigPath = ""

func init() {
	RootCmd.AddCommand(RunCmd)
	//RootCmd.AddCommand(TestCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
