/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2021 IBM, Inc.
 *
 */

package main

import (
	"kubevirt.io/client-go/log"
	"kubevirt.io/kubevirt/tools/perfscale-load-generator/burst"
	"kubevirt.io/kubevirt/tools/perfscale-load-generator/config"
	"kubevirt.io/kubevirt/tools/perfscale-load-generator/flags"
)

func main() {
	log.Log.SetVerbosityLevel(flags.Verbosity)

	log.Log.V(1).Infof("Running Load Generator Banchmark")

	workload := config.NewWorkload(flags.WorkloadConfigFile)
	client := config.NewKubevirtClient()

	lg := burst.NewBurstLoadGenerator(client, workload)

	if flags.Delete {
		lg.DeleteWorkload()
	} else {
		lg.CreateWorkload()
	}
}
