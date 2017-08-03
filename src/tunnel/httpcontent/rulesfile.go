package httpcontent

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)


func (t *HttpContentTunnel) readRules() error {

	file_content, err := ioutil.ReadFile(t.RulesFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file_content, &t.Rules); err != nil {
		return err
	}

	return nil
}



func (t *HttpContentTunnel) getHostByRules(function_name, target, terminalid string) (string, error) {

	for _, rule := range t.Rules {
		fmt.Println(rule)
	}

	return "defaulthostport", nil
}



