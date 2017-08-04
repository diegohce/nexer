package httpcontent

import (
	"log"
	"io/ioutil"
	"encoding/json"
)


type hostByRule struct {
	hostname string
	rewrite  string
}


func (t *HttpContentTunnel) readRules() error {

	file_content, err := ioutil.ReadFile(t.RulesFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file_content, &t.Rules); err != nil {
		return err
	}

	for _, rule := range t.Rules {
		log.Println("RULES", rule)
	}

	return nil
}



func (t *HttpContentTunnel) getHostByRules(function_name, target, terminalid string) (*hostByRule, error) {

	var hostport string
	var rewrite string

	for _, rule := range t.Rules {

		//fmt.Println("RULES", rule)
		hostport = rule[4]
		rewrite = rule[3]

		if rule[0] == "*" && rule[1] == "*" && rule[2] == "*" {
			break
		} else if rule[0] == function_name && rule[1] == "*" && rule[2] == "*" {
			break
		} else if rule[0] == "*" && rule[1] == target && rule[2] == "*" {
			break
		} else if rule[0] == "*" && rule[1] == "*" && rule[2] == terminalid {
			break
		} else if rule[0] == function_name && rule[1] == target && rule[2] == "*" {
			break
		} else if rule[0] == function_name && rule[1] == target && rule[2] == terminalid {
			break
		} else if rule[0] == "*" && rule[1] == target && rule[2] == terminalid {
			break
		}
	}

	return &hostByRule{hostname: hostport, rewrite: rewrite}, nil
}



