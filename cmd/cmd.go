package cmd

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	verbose bool
)

var persistentEnvs = make(map[string]string)
var localEnvs = make(map[string]string)

// RootCmd is a global variable which will handle all subcommands
var RootCmd = &cobra.Command{
	Use: "bivac",
}

func initConfig() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	viper.AutomaticEnv()
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	localEnvs["BIVAC_VERBOSE"] = "verbose"

	SetValuesFromEnv(localEnvs, RootCmd.PersistentFlags())
	SetValuesFromEnv(persistentEnvs, RootCmd.PersistentFlags())
}

// Execute is the main thread, required by Cobra
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func SetValuesFromEnv(envs map[string]string, flags *pflag.FlagSet) {
	for env, flag := range envs {
		flag := flags.Lookup(flag)
		flag.Usage = fmt.Sprintf("%v [%v]", flag.Usage, env)
		if value := os.Getenv(env); value != "" {
			flag.Value.Set(value)
		}
	}
	return
}