// Copyright 2026 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e_parallel

package ports

import (
	"io"

	"github.com/haproxytech/kubernetes-ingress/deploy/tests/e2e"
)

func (suite *PortsSuite) Test_Ports() {
	suite.Run("Different Backend Ports OK", func() {
		suite.Eventually(func() bool {
			counter := map[string]int{}
			for i := 0; i < 4; i++ {
				func() {
					res, cls, err := suite.client.Do()
					if err != nil {
						suite.T().Log(err.Error())
					}
					defer cls()
					if res.StatusCode == 200 {
						body, err := io.ReadAll(res.Body)
						if err != nil {
							suite.T().Log(err.Error())
							return
						}
						counter[string(body)]++
					}
				}()
			}
			for _, v := range counter {
				if v != 2 {
					return false
				}
			}
			return true
		}, e2e.WaitDuration, e2e.TickDuration)
	})
}
